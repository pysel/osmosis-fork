package keeper

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/osmosis-labs/osmosis/osmomath"
	"github.com/osmosis-labs/osmosis/v19/x/txfees/keeper/txfee_filters"
	"github.com/osmosis-labs/osmosis/v19/x/txfees/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// MempoolFeeDecorator will check if the transaction's fee is at least as large
// as the local validator's minimum gasFee (defined in validator config).
// If fee is too low, decorator returns error and tx is rejected from mempool.
// Note this only applies when ctx.CheckTx = true
// If fee is high enough or not CheckTx, then call next AnteHandler
// CONTRACT: Tx must implement FeeTx to use MempoolFeeDecorator.
type MempoolFeeDecorator struct {
	TxFeesKeeper Keeper
	Opts         types.MempoolFeeOptions
}

func NewMempoolFeeDecorator(txFeesKeeper Keeper, opts types.MempoolFeeOptions) MempoolFeeDecorator {
	return MempoolFeeDecorator{
		TxFeesKeeper: txFeesKeeper,
		Opts:         opts,
	}
}

func (mfd MempoolFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	// The SDK currently requires all txs to be FeeTx's in CheckTx, within its mempool fee decorator.
	// See: https://github.com/cosmos/cosmos-sdk/blob/f726a2398a26bdaf71d78dbf56a82621e84fd098/x/auth/middleware/fee.go#L34-L37
	// So this is not a real restriction at the moment.
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	// Ensure that the provided gas is less than the maximum gas per tx,
	// if this is a CheckTx. This is only for local mempool purposes, and thus
	// is only ran on check tx.
	if ctx.IsCheckTx() && !simulate {
		if feeTx.GetGas() > mfd.Opts.MaxGasWantedPerTx {
			msg := "Too much gas wanted: %d, maximum is %d"
			return ctx, errorsmod.Wrapf(sdkerrors.ErrOutOfGas, msg, feeTx.GetGas(), mfd.Opts.MaxGasWantedPerTx)
		}
	}

	feeCoins := feeTx.GetFee()

	if len(feeCoins) > 1 {
		return ctx, types.ErrTooManyFeeCoins
	}

	baseDenom, err := mfd.TxFeesKeeper.GetBaseDenom(ctx)
	if err != nil {
		return ctx, err
	}

	// If there is a fee attached to the tx, make sure the fee denom is a denom accepted by the chain
	if len(feeCoins) == 1 {
		feeDenom := feeCoins.GetDenomByIndex(0)
		if feeDenom != baseDenom {
			_, err := mfd.TxFeesKeeper.GetFeeToken(ctx, feeDenom)
			if err != nil {
				return ctx, err
			}
		}
	}

	// Determine if these fees are sufficient for the tx to pass.
	// Once ABCI++ Process Proposal lands, we can have block validity conditions enforce this.
	minBaseGasPrice := mfd.getMinBaseGasPrice(ctx, baseDenom, simulate, feeTx)

	// If minBaseGasPrice is zero, then we don't need to check the fee. Continue
	if minBaseGasPrice.IsZero() {
		return next(ctx, tx, simulate)
	}
	// You should only be able to pay with one fee token in a single tx
	if len(feeCoins) != 1 {
		return ctx, errorsmod.Wrapf(sdkerrors.ErrInsufficientFee,
			"Expected 1 fee denom attached, got %d", len(feeCoins))
	}
	// The minimum base gas price is in uosmo, convert the fee denom's worth to uosmo terms.
	// Then compare if its sufficient for paying the tx fee.
	err = mfd.TxFeesKeeper.IsSufficientFee(ctx, minBaseGasPrice, feeTx.GetGas(), feeCoins[0])
	if err != nil {
		return ctx, err
	}

	return next(ctx, tx, simulate)
}

func (mfd MempoolFeeDecorator) getMinBaseGasPrice(ctx sdk.Context, baseDenom string, simulate bool, feeTx sdk.FeeTx) osmomath.Dec {
	// In block execution (DeliverTx), its set to the governance decided upon consensus min fee.
	minBaseGasPrice := types.ConsensusMinFee
	// If we are in CheckTx, a separate function is ran locally to ensure sufficient fees for entering our mempool.
	// So we ensure that the provided fees meet a minimum threshold for the validator
	if (ctx.IsCheckTx() || ctx.IsReCheckTx()) && !simulate {
		minBaseGasPrice = sdk.MaxDec(minBaseGasPrice, mfd.GetMinBaseGasPriceForTx(ctx, baseDenom, feeTx))
	}
	// If we are in genesis or are simulating a tx, then we actually override all of the above, to set it to 0.
	if ctx.IsGenesis() || simulate {
		minBaseGasPrice = osmomath.ZeroDec()
	}
	return minBaseGasPrice
}

// IsSufficientFee checks if the feeCoin provided (in any asset), is worth enough osmo at current spot prices
// to pay the gas cost of this tx.
func (k Keeper) IsSufficientFee(ctx sdk.Context, minBaseGasPrice osmomath.Dec, gasRequested uint64, feeCoin sdk.Coin) error {
	baseDenom, err := k.GetBaseDenom(ctx)
	if err != nil {
		return err
	}

	// Determine the required fees by multiplying the required minimum gas
	// price by the gas limit, where fee = ceil(minGasPrice * gasLimit).
	glDec := osmomath.NewDec(int64(gasRequested))
	requiredBaseFee := sdk.NewCoin(baseDenom, minBaseGasPrice.Mul(glDec).Ceil().RoundInt())

	convertedFee, err := k.ConvertToBaseToken(ctx, feeCoin)
	if err != nil {
		return err
	}
	// check to ensure that the convertedFee should always be greater than or equal to the requireBaseFee
	if !(convertedFee.IsGTE(requiredBaseFee)) {
		return errorsmod.Wrapf(sdkerrors.ErrInsufficientFee, "insufficient fees; got: %s which converts to %s. required: %s", feeCoin, convertedFee, requiredBaseFee)
	}

	return nil
}

func (mfd MempoolFeeDecorator) GetMinBaseGasPriceForTx(ctx sdk.Context, baseDenom string, tx sdk.FeeTx) osmomath.Dec {
	cfgMinGasPrice := ctx.MinGasPrices().AmountOf(baseDenom)
	// the check below prevents tx gas from getting over HighGasTxThreshold which is default to 1_000_000
	if tx.GetGas() >= mfd.Opts.HighGasTxThreshold {
		cfgMinGasPrice = sdk.MaxDec(cfgMinGasPrice, mfd.Opts.MinGasPriceForHighGasTx)
	}
	if txfee_filters.IsArbTxLoose(tx) {
		cfgMinGasPrice = sdk.MaxDec(cfgMinGasPrice, mfd.Opts.MinGasPriceForArbitrageTx)
	}
	return cfgMinGasPrice
}

// DeductFeeDecorator deducts fees from the first signer of the tx.
// If the first signer does not have the funds to pay for the fees, we return an InsufficientFunds error.
// We call next AnteHandler if fees successfully deducted.
//
// CONTRACT: Tx must implement FeeTx interface to use DeductFeeDecorator
type DeductFeeDecorator struct {
	ak             types.AccountKeeper
	bankKeeper     types.BankKeeper
	feegrantKeeper types.FeegrantKeeper
	txFeesKeeper   Keeper
}

func NewDeductFeeDecorator(tk Keeper, ak types.AccountKeeper, bk types.BankKeeper, fk types.FeegrantKeeper) DeductFeeDecorator {
	return DeductFeeDecorator{
		ak:             ak,
		bankKeeper:     bk,
		feegrantKeeper: fk,
		txFeesKeeper:   tk,
	}
}

func (dfd DeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	// checks to make sure the module account has been set to collect fees in base token
	if addr := dfd.ak.GetModuleAddress(types.FeeCollectorName); addr == nil {
		return ctx, fmt.Errorf("fee collector module account (%s) has not been set", types.FeeCollectorName)
	}

	// checks to make sure a separate module account has been set to collect fees not in base token
	if addrNonNativeFee := dfd.ak.GetModuleAddress(types.FeeCollectorForStakingRewardsName); addrNonNativeFee == nil {
		return ctx, fmt.Errorf("fee collector for staking module account (%s) has not been set", types.FeeCollectorForStakingRewardsName)
	}

	// fee can be in any denom (checked for validity later)
	fee := feeTx.GetFee()
	feePayer := feeTx.FeePayer()
	feeGranter := feeTx.FeeGranter()

	// set the fee payer as the default address to deduct fees from
	deductFeesFrom := feePayer

	// If a fee granter was set, deduct fee from the fee granter's account.
	if feeGranter != nil {
		if dfd.feegrantKeeper == nil {
			return ctx, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "fee grants is not enabled")
		} else if !feeGranter.Equals(feePayer) {
			err := dfd.feegrantKeeper.UseGrantedFees(ctx, feeGranter, feePayer, fee, tx.GetMsgs())
			if err != nil {
				return ctx, errorsmod.Wrapf(err, "%s not allowed to pay fees from %s", feeGranter, feePayer)
			}
		}

		// if no errors, change the account that is charged for fees to the fee granter
		deductFeesFrom = feeGranter
	}

	deductFeesFromAcc := dfd.ak.GetAccount(ctx, deductFeesFrom)
	if deductFeesFromAcc == nil {
		return ctx, errorsmod.Wrapf(sdkerrors.ErrUnknownAddress, "fee payer address: %s does not exist", deductFeesFrom)
	}

	// deducts the fees and transfer them to the module account
	if !feeTx.GetFee().IsZero() {
		err = DeductFees(dfd.txFeesKeeper, dfd.bankKeeper, ctx, deductFeesFromAcc, feeTx.GetFee())
		if err != nil {
			return ctx, err
		}
	}

	ctx.EventManager().EmitEvents(sdk.Events{sdk.NewEvent(sdk.EventTypeTx,
		sdk.NewAttribute(sdk.AttributeKeyFee, feeTx.GetFee().String()),
	)})

	return next(ctx, tx, simulate)
}

// DeductFees deducts fees from the given account and transfers them to the set module account.
func DeductFees(txFeesKeeper types.TxFeesKeeper, bankKeeper types.BankKeeper, ctx sdk.Context, acc authtypes.AccountI, fees sdk.Coins) error {
	// Checks the validity of the fee tokens (sorted, have positive amount, valid and unique denomination)
	if !fees.IsValid() {
		return errorsmod.Wrapf(sdkerrors.ErrInsufficientFee, "invalid fee amount: %s", fees)
	}

	// pulls base denom from TxFeesKeeper (should be uOSMO)
	baseDenom, err := txFeesKeeper.GetBaseDenom(ctx)
	if err != nil {
		return err
	}

	// checks if input fee is uOSMO (assumes only one fee token exists in the fees array (as per the check in mempoolFeeDecorator))
	if fees[0].Denom == baseDenom {
		// sends to FeeCollectorName module account, which sends to staking rewards
		err := bankKeeper.SendCoinsFromAccountToModule(ctx, acc.GetAddress(), types.FeeCollectorName, fees)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, err.Error())
		}
	} else {
		// sends to FeeCollectorForStakingRewardsName module account
		err := bankKeeper.SendCoinsFromAccountToModule(ctx, acc.GetAddress(), types.FeeCollectorForStakingRewardsName, fees)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, err.Error())
		}
	}

	return nil
}
