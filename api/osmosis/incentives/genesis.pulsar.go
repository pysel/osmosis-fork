// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package incentives

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_GenesisState_2_list)(nil)

type _GenesisState_2_list struct {
	list *[]*Gauge
}

func (x *_GenesisState_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Gauge)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Gauge)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_2_list) AppendMutable() protoreflect.Value {
	v := new(Gauge)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_2_list) NewElement() protoreflect.Value {
	v := new(Gauge)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_2_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_GenesisState_3_list)(nil)

type _GenesisState_3_list struct {
	list *[]*durationpb.Duration
}

func (x *_GenesisState_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*durationpb.Duration)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*durationpb.Duration)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_3_list) AppendMutable() protoreflect.Value {
	v := new(durationpb.Duration)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_3_list) NewElement() protoreflect.Value {
	v := new(durationpb.Duration)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_GenesisState                    protoreflect.MessageDescriptor
	fd_GenesisState_params             protoreflect.FieldDescriptor
	fd_GenesisState_gauges             protoreflect.FieldDescriptor
	fd_GenesisState_lockable_durations protoreflect.FieldDescriptor
	fd_GenesisState_last_gauge_id      protoreflect.FieldDescriptor
)

func init() {
	file_osmosis_incentives_genesis_proto_init()
	md_GenesisState = File_osmosis_incentives_genesis_proto.Messages().ByName("GenesisState")
	fd_GenesisState_params = md_GenesisState.Fields().ByName("params")
	fd_GenesisState_gauges = md_GenesisState.Fields().ByName("gauges")
	fd_GenesisState_lockable_durations = md_GenesisState.Fields().ByName("lockable_durations")
	fd_GenesisState_last_gauge_id = md_GenesisState.Fields().ByName("last_gauge_id")
}

var _ protoreflect.Message = (*fastReflection_GenesisState)(nil)

type fastReflection_GenesisState GenesisState

func (x *GenesisState) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GenesisState)(x)
}

func (x *GenesisState) slowProtoReflect() protoreflect.Message {
	mi := &file_osmosis_incentives_genesis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GenesisState_messageType fastReflection_GenesisState_messageType
var _ protoreflect.MessageType = fastReflection_GenesisState_messageType{}

type fastReflection_GenesisState_messageType struct{}

func (x fastReflection_GenesisState_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GenesisState)(nil)
}
func (x fastReflection_GenesisState_messageType) New() protoreflect.Message {
	return new(fastReflection_GenesisState)
}
func (x fastReflection_GenesisState_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GenesisState
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GenesisState) Descriptor() protoreflect.MessageDescriptor {
	return md_GenesisState
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GenesisState) Type() protoreflect.MessageType {
	return _fastReflection_GenesisState_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GenesisState) New() protoreflect.Message {
	return new(fastReflection_GenesisState)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GenesisState) Interface() protoreflect.ProtoMessage {
	return (*GenesisState)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GenesisState) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Params != nil {
		value := protoreflect.ValueOfMessage(x.Params.ProtoReflect())
		if !f(fd_GenesisState_params, value) {
			return
		}
	}
	if len(x.Gauges) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_2_list{list: &x.Gauges})
		if !f(fd_GenesisState_gauges, value) {
			return
		}
	}
	if len(x.LockableDurations) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_3_list{list: &x.LockableDurations})
		if !f(fd_GenesisState_lockable_durations, value) {
			return
		}
	}
	if x.LastGaugeId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.LastGaugeId)
		if !f(fd_GenesisState_last_gauge_id, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_GenesisState) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "osmosis.incentives.GenesisState.params":
		return x.Params != nil
	case "osmosis.incentives.GenesisState.gauges":
		return len(x.Gauges) != 0
	case "osmosis.incentives.GenesisState.lockable_durations":
		return len(x.LockableDurations) != 0
	case "osmosis.incentives.GenesisState.last_gauge_id":
		return x.LastGaugeId != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: osmosis.incentives.GenesisState"))
		}
		panic(fmt.Errorf("message osmosis.incentives.GenesisState does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "osmosis.incentives.GenesisState.params":
		x.Params = nil
	case "osmosis.incentives.GenesisState.gauges":
		x.Gauges = nil
	case "osmosis.incentives.GenesisState.lockable_durations":
		x.LockableDurations = nil
	case "osmosis.incentives.GenesisState.last_gauge_id":
		x.LastGaugeId = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: osmosis.incentives.GenesisState"))
		}
		panic(fmt.Errorf("message osmosis.incentives.GenesisState does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GenesisState) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "osmosis.incentives.GenesisState.params":
		value := x.Params
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "osmosis.incentives.GenesisState.gauges":
		if len(x.Gauges) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_2_list{})
		}
		listValue := &_GenesisState_2_list{list: &x.Gauges}
		return protoreflect.ValueOfList(listValue)
	case "osmosis.incentives.GenesisState.lockable_durations":
		if len(x.LockableDurations) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_3_list{})
		}
		listValue := &_GenesisState_3_list{list: &x.LockableDurations}
		return protoreflect.ValueOfList(listValue)
	case "osmosis.incentives.GenesisState.last_gauge_id":
		value := x.LastGaugeId
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: osmosis.incentives.GenesisState"))
		}
		panic(fmt.Errorf("message osmosis.incentives.GenesisState does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "osmosis.incentives.GenesisState.params":
		x.Params = value.Message().Interface().(*Params)
	case "osmosis.incentives.GenesisState.gauges":
		lv := value.List()
		clv := lv.(*_GenesisState_2_list)
		x.Gauges = *clv.list
	case "osmosis.incentives.GenesisState.lockable_durations":
		lv := value.List()
		clv := lv.(*_GenesisState_3_list)
		x.LockableDurations = *clv.list
	case "osmosis.incentives.GenesisState.last_gauge_id":
		x.LastGaugeId = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: osmosis.incentives.GenesisState"))
		}
		panic(fmt.Errorf("message osmosis.incentives.GenesisState does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "osmosis.incentives.GenesisState.params":
		if x.Params == nil {
			x.Params = new(Params)
		}
		return protoreflect.ValueOfMessage(x.Params.ProtoReflect())
	case "osmosis.incentives.GenesisState.gauges":
		if x.Gauges == nil {
			x.Gauges = []*Gauge{}
		}
		value := &_GenesisState_2_list{list: &x.Gauges}
		return protoreflect.ValueOfList(value)
	case "osmosis.incentives.GenesisState.lockable_durations":
		if x.LockableDurations == nil {
			x.LockableDurations = []*durationpb.Duration{}
		}
		value := &_GenesisState_3_list{list: &x.LockableDurations}
		return protoreflect.ValueOfList(value)
	case "osmosis.incentives.GenesisState.last_gauge_id":
		panic(fmt.Errorf("field last_gauge_id of message osmosis.incentives.GenesisState is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: osmosis.incentives.GenesisState"))
		}
		panic(fmt.Errorf("message osmosis.incentives.GenesisState does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GenesisState) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "osmosis.incentives.GenesisState.params":
		m := new(Params)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "osmosis.incentives.GenesisState.gauges":
		list := []*Gauge{}
		return protoreflect.ValueOfList(&_GenesisState_2_list{list: &list})
	case "osmosis.incentives.GenesisState.lockable_durations":
		list := []*durationpb.Duration{}
		return protoreflect.ValueOfList(&_GenesisState_3_list{list: &list})
	case "osmosis.incentives.GenesisState.last_gauge_id":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: osmosis.incentives.GenesisState"))
		}
		panic(fmt.Errorf("message osmosis.incentives.GenesisState does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GenesisState) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in osmosis.incentives.GenesisState", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GenesisState) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_GenesisState) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GenesisState) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GenesisState)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Params != nil {
			l = options.Size(x.Params)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Gauges) > 0 {
			for _, e := range x.Gauges {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.LockableDurations) > 0 {
			for _, e := range x.LockableDurations {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.LastGaugeId != 0 {
			n += 1 + runtime.Sov(uint64(x.LastGaugeId))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GenesisState)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.LastGaugeId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.LastGaugeId))
			i--
			dAtA[i] = 0x20
		}
		if len(x.LockableDurations) > 0 {
			for iNdEx := len(x.LockableDurations) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.LockableDurations[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x1a
			}
		}
		if len(x.Gauges) > 0 {
			for iNdEx := len(x.Gauges) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Gauges[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x12
			}
		}
		if x.Params != nil {
			encoded, err := options.Marshal(x.Params)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*GenesisState)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Params == nil {
					x.Params = &Params{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Params); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Gauges", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Gauges = append(x.Gauges, &Gauge{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Gauges[len(x.Gauges)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LockableDurations", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.LockableDurations = append(x.LockableDurations, &durationpb.Duration{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.LockableDurations[len(x.LockableDurations)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LastGaugeId", wireType)
				}
				x.LastGaugeId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.LastGaugeId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: osmosis/incentives/genesis.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GenesisState defines the incentives module's various parameters when first
// initialized
type GenesisState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// params are all the parameters of the module
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	// gauges are all gauges that should exist at genesis
	Gauges []*Gauge `protobuf:"bytes,2,rep,name=gauges,proto3" json:"gauges,omitempty"`
	// lockable_durations are all lockup durations that gauges can be locked for
	// in order to recieve incentives
	LockableDurations []*durationpb.Duration `protobuf:"bytes,3,rep,name=lockable_durations,json=lockableDurations,proto3" json:"lockable_durations,omitempty"`
	// last_gauge_id is what the gauge number will increment from when creating
	// the next gauge after genesis
	LastGaugeId uint64 `protobuf:"varint,4,opt,name=last_gauge_id,json=lastGaugeId,proto3" json:"last_gauge_id,omitempty"`
}

func (x *GenesisState) Reset() {
	*x = GenesisState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_osmosis_incentives_genesis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState) ProtoMessage() {}

// Deprecated: Use GenesisState.ProtoReflect.Descriptor instead.
func (*GenesisState) Descriptor() ([]byte, []int) {
	return file_osmosis_incentives_genesis_proto_rawDescGZIP(), []int{0}
}

func (x *GenesisState) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *GenesisState) GetGauges() []*Gauge {
	if x != nil {
		return x.Gauges
	}
	return nil
}

func (x *GenesisState) GetLockableDurations() []*durationpb.Duration {
	if x != nil {
		return x.LockableDurations
	}
	return nil
}

func (x *GenesisState) GetLastGaugeId() uint64 {
	if x != nil {
		return x.LastGaugeId
	}
	return 0
}

var File_osmosis_incentives_genesis_proto protoreflect.FileDescriptor

var file_osmosis_incentives_genesis_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x65, 0x6e, 0x74,
	0x69, 0x76, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x65,
	0x6e, 0x74, 0x69, 0x76, 0x65, 0x73, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x76, 0x65, 0x73,
	0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x76, 0x65,
	0x73, 0x2f, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x02,
	0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x38,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x69,
	0x76, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00,
	0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x37, 0x0a, 0x06, 0x67, 0x61, 0x75, 0x67,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x76, 0x65, 0x73, 0x2e, 0x47, 0x61,
	0x75, 0x67, 0x65, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x06, 0x67, 0x61, 0x75, 0x67, 0x65,
	0x73, 0x12, 0x6f, 0x0a, 0x12, 0x6c, 0x6f, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x25, 0xc8, 0xde, 0x1f, 0x00, 0xf2, 0xde,
	0x1f, 0x19, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x6c, 0x6f, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x98, 0xdf, 0x1f, 0x01, 0x52,
	0x11, 0x6c, 0x6f, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x67, 0x61, 0x75, 0x67, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x47,
	0x61, 0x75, 0x67, 0x65, 0x49, 0x64, 0x42, 0xc7, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x76, 0x65,
	0x73, 0x42, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2d, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2f, 0x69,
	0x6e, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x76, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x4f, 0x49, 0x58, 0xaa,
	0x02, 0x12, 0x4f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2e, 0x49, 0x6e, 0x63, 0x65, 0x6e, 0x74,
	0x69, 0x76, 0x65, 0x73, 0xca, 0x02, 0x12, 0x4f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x5c, 0x49,
	0x6e, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x76, 0x65, 0x73, 0xe2, 0x02, 0x1e, 0x4f, 0x73, 0x6d, 0x6f,
	0x73, 0x69, 0x73, 0x5c, 0x49, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x76, 0x65, 0x73, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x4f, 0x73, 0x6d,
	0x6f, 0x73, 0x69, 0x73, 0x3a, 0x3a, 0x49, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x76, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_osmosis_incentives_genesis_proto_rawDescOnce sync.Once
	file_osmosis_incentives_genesis_proto_rawDescData = file_osmosis_incentives_genesis_proto_rawDesc
)

func file_osmosis_incentives_genesis_proto_rawDescGZIP() []byte {
	file_osmosis_incentives_genesis_proto_rawDescOnce.Do(func() {
		file_osmosis_incentives_genesis_proto_rawDescData = protoimpl.X.CompressGZIP(file_osmosis_incentives_genesis_proto_rawDescData)
	})
	return file_osmosis_incentives_genesis_proto_rawDescData
}

var file_osmosis_incentives_genesis_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_osmosis_incentives_genesis_proto_goTypes = []interface{}{
	(*GenesisState)(nil),        // 0: osmosis.incentives.GenesisState
	(*Params)(nil),              // 1: osmosis.incentives.Params
	(*Gauge)(nil),               // 2: osmosis.incentives.Gauge
	(*durationpb.Duration)(nil), // 3: google.protobuf.Duration
}
var file_osmosis_incentives_genesis_proto_depIdxs = []int32{
	1, // 0: osmosis.incentives.GenesisState.params:type_name -> osmosis.incentives.Params
	2, // 1: osmosis.incentives.GenesisState.gauges:type_name -> osmosis.incentives.Gauge
	3, // 2: osmosis.incentives.GenesisState.lockable_durations:type_name -> google.protobuf.Duration
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_osmosis_incentives_genesis_proto_init() }
func file_osmosis_incentives_genesis_proto_init() {
	if File_osmosis_incentives_genesis_proto != nil {
		return
	}
	file_osmosis_incentives_params_proto_init()
	file_osmosis_incentives_gauge_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_osmosis_incentives_genesis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_osmosis_incentives_genesis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_osmosis_incentives_genesis_proto_goTypes,
		DependencyIndexes: file_osmosis_incentives_genesis_proto_depIdxs,
		MessageInfos:      file_osmosis_incentives_genesis_proto_msgTypes,
	}.Build()
	File_osmosis_incentives_genesis_proto = out.File
	file_osmosis_incentives_genesis_proto_rawDesc = nil
	file_osmosis_incentives_genesis_proto_goTypes = nil
	file_osmosis_incentives_genesis_proto_depIdxs = nil
}
