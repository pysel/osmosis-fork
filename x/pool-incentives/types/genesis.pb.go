// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/pool-incentives/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the pool incentives module's genesis state.
type GenesisState struct {
	// params defines all the paramaters of the module.
	Params            Params          `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	LockableDurations []time.Duration `protobuf:"bytes,2,rep,name=lockable_durations,json=lockableDurations,proto3,stdduration" json:"lockable_durations" yaml:"lockable_durations"`
	DistrInfo         *DistrInfo      `protobuf:"bytes,3,opt,name=distr_info,json=distrInfo,proto3" json:"distr_info,omitempty" yaml:"distr_info"`
	PoolToGauges      *PoolToGauges   `protobuf:"bytes,4,opt,name=pool_to_gauges,json=poolToGauges,proto3" json:"pool_to_gauges,omitempty" yaml:"pool_to_gauges"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc1f078212600632, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetLockableDurations() []time.Duration {
	if m != nil {
		return m.LockableDurations
	}
	return nil
}

func (m *GenesisState) GetDistrInfo() *DistrInfo {
	if m != nil {
		return m.DistrInfo
	}
	return nil
}

func (m *GenesisState) GetPoolToGauges() *PoolToGauges {
	if m != nil {
		return m.PoolToGauges
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "osmosis.poolincentives.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("osmosis/pool-incentives/v1beta1/genesis.proto", fileDescriptor_cc1f078212600632)
}

var fileDescriptor_cc1f078212600632 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0xaa, 0xd3, 0x40,
	0x14, 0xc7, 0x13, 0x7b, 0xb9, 0x60, 0xee, 0x45, 0xb8, 0x41, 0x21, 0x2d, 0x38, 0x29, 0x01, 0xa5,
	0x82, 0x9d, 0xb1, 0x75, 0xa5, 0xee, 0x42, 0xa1, 0xb8, 0x93, 0xa8, 0x1b, 0x37, 0x61, 0x92, 0x4e,
	0xc7, 0xc1, 0x24, 0x27, 0x66, 0x26, 0xc5, 0xbe, 0x85, 0x4b, 0x37, 0xbe, 0x4f, 0x97, 0x5d, 0xba,
	0x8a, 0xd2, 0xbe, 0x41, 0x9f, 0x40, 0x32, 0x49, 0x68, 0xa5, 0x60, 0x77, 0x19, 0xce, 0xef, 0xff,
	0x71, 0x0e, 0xb1, 0xc6, 0x20, 0x53, 0x90, 0x42, 0x92, 0x1c, 0x20, 0x19, 0x8b, 0x2c, 0x66, 0x99,
	0x12, 0x2b, 0x26, 0xc9, 0x6a, 0x12, 0x31, 0x45, 0x27, 0x84, 0xb3, 0x8c, 0x49, 0x21, 0x71, 0x5e,
	0x80, 0x02, 0x1b, 0xb5, 0x38, 0xae, 0xf1, 0x23, 0x8d, 0x5b, 0x7a, 0xf0, 0x90, 0x03, 0x07, 0x8d,
	0x92, 0xfa, 0xab, 0x51, 0x0d, 0x10, 0x07, 0xe0, 0x09, 0x23, 0xfa, 0x15, 0x95, 0x4b, 0xb2, 0x28,
	0x0b, 0xaa, 0x04, 0x64, 0xed, 0xfc, 0xc5, 0xa5, 0x12, 0x27, 0x49, 0x5a, 0xe1, 0xfd, 0xec, 0x59,
	0xb7, 0xf3, 0xa6, 0xd9, 0x7b, 0x45, 0x15, 0xb3, 0x67, 0xd6, 0x75, 0x4e, 0x0b, 0x9a, 0x4a, 0xc7,
	0x1c, 0x9a, 0xa3, 0x9b, 0xe9, 0x53, 0xfc, 0xff, 0xa6, 0xf8, 0x9d, 0xa6, 0xfd, 0xab, 0x4d, 0xe5,
	0x1a, 0x41, 0xab, 0xb5, 0xc1, 0xb2, 0x13, 0x88, 0xbf, 0xd0, 0x28, 0x61, 0x61, 0xd7, 0x51, 0x3a,
	0xf7, 0x86, 0xbd, 0xd1, 0xcd, 0xb4, 0x8f, 0x9b, 0x2d, 0x70, 0xb7, 0x05, 0x9e, 0xb5, 0x84, 0xff,
	0xa4, 0x36, 0x39, 0x54, 0x6e, 0x7f, 0x4d, 0xd3, 0xe4, 0xb5, 0x77, 0x6e, 0xe1, 0xfd, 0xf8, 0xed,
	0x9a, 0xc1, 0x5d, 0x37, 0xe8, 0x84, 0xd2, 0x8e, 0x2d, 0x6b, 0x21, 0xa4, 0x2a, 0x42, 0x91, 0x2d,
	0xc1, 0xe9, 0xe9, 0xea, 0xcf, 0x2e, 0x55, 0x9f, 0xd5, 0x8a, 0xb7, 0xd9, 0x12, 0xfc, 0xfe, 0xa6,
	0x72, 0xcd, 0x43, 0xe5, 0xde, 0x35, 0xc1, 0x47, 0x2b, 0x2f, 0xb8, 0xbf, 0xe8, 0x28, 0xfb, 0xab,
	0xf5, 0xa0, 0x76, 0x0a, 0x15, 0x84, 0x9c, 0x96, 0x9c, 0x49, 0xe7, 0x4a, 0x07, 0x3d, 0xbf, 0x78,
	0x23, 0x80, 0xe4, 0x03, 0xcc, 0xb5, 0xc6, 0x7f, 0xdc, 0x66, 0x3d, 0x6a, 0xb2, 0xfe, 0x75, 0xf4,
	0x82, 0xdb, 0xfc, 0x14, 0xfe, 0xb8, 0xd9, 0x21, 0x73, 0xbb, 0x43, 0xe6, 0x9f, 0x1d, 0x32, 0xbf,
	0xef, 0x91, 0xb1, 0xdd, 0x23, 0xe3, 0xd7, 0x1e, 0x19, 0x9f, 0xde, 0x70, 0xa1, 0x3e, 0x97, 0x11,
	0x8e, 0x21, 0x25, 0x6d, 0xfc, 0x38, 0xa1, 0x91, 0xec, 0x1e, 0x64, 0x35, 0x79, 0x45, 0xbe, 0x9d,
	0xfd, 0x09, 0x6a, 0x9d, 0x33, 0x19, 0x5d, 0xeb, 0xdb, 0xbf, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0xba, 0xa9, 0x3b, 0xfc, 0xb6, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PoolToGauges != nil {
		{
			size, err := m.PoolToGauges.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.DistrInfo != nil {
		{
			size, err := m.DistrInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.LockableDurations) > 0 {
		for iNdEx := len(m.LockableDurations) - 1; iNdEx >= 0; iNdEx-- {
			n, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.LockableDurations[iNdEx], dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.LockableDurations[iNdEx]):])
			if err != nil {
				return 0, err
			}
			i -= n
			i = encodeVarintGenesis(dAtA, i, uint64(n))
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.LockableDurations) > 0 {
		for _, e := range m.LockableDurations {
			l = github_com_gogo_protobuf_types.SizeOfStdDuration(e)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.DistrInfo != nil {
		l = m.DistrInfo.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.PoolToGauges != nil {
		l = m.PoolToGauges.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockableDurations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LockableDurations = append(m.LockableDurations, time.Duration(0))
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&(m.LockableDurations[len(m.LockableDurations)-1]), dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistrInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DistrInfo == nil {
				m.DistrInfo = &DistrInfo{}
			}
			if err := m.DistrInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolToGauges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PoolToGauges == nil {
				m.PoolToGauges = &PoolToGauges{}
			}
			if err := m.PoolToGauges.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
