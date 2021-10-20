// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/slashing/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgUnjail defines the Msg/Unjail request type
type MsgUnjail struct {
	ValidatorAddr string `protobuf:"bytes,1,opt,name=validator_addr,json=validatorAddr,proto3" json:"address"`
}

func (m *MsgUnjail) Reset()         { *m = MsgUnjail{} }
func (m *MsgUnjail) String() string { return proto.CompactTextString(m) }
func (*MsgUnjail) ProtoMessage()    {}
func (*MsgUnjail) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c5611c0c4a59d9d, []int{0}
}
func (m *MsgUnjail) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnjail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnjail.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnjail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnjail.Merge(m, src)
}
func (m *MsgUnjail) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnjail) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnjail.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnjail proto.InternalMessageInfo

// MsgUnjailResponse defines the Msg/Unjail response type
type MsgUnjailResponse struct {
}

func (m *MsgUnjailResponse) Reset()         { *m = MsgUnjailResponse{} }
func (m *MsgUnjailResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUnjailResponse) ProtoMessage()    {}
func (*MsgUnjailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c5611c0c4a59d9d, []int{1}
}
func (m *MsgUnjailResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnjailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnjailResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnjailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnjailResponse.Merge(m, src)
}
func (m *MsgUnjailResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnjailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnjailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnjailResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUnjail)(nil), "cosmos.slashing.v1beta1.MsgUnjail")
	proto.RegisterType((*MsgUnjailResponse)(nil), "cosmos.slashing.v1beta1.MsgUnjailResponse")
}

func init() { proto.RegisterFile("cosmos/slashing/v1beta1/tx.proto", fileDescriptor_3c5611c0c4a59d9d) }

var fileDescriptor_3c5611c0c4a59d9d = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x2f, 0xce, 0x49, 0x2c, 0xce, 0xc8, 0xcc, 0x4b, 0xd7, 0x2f, 0x33, 0x4c, 0x4a,
	0x2d, 0x49, 0x34, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x87, 0xa8,
	0xd0, 0x83, 0xa9, 0xd0, 0x83, 0xaa, 0x90, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xab, 0xd1, 0x07,
	0xb1, 0x20, 0xca, 0xa5, 0x24, 0x21, 0xca, 0xe3, 0x21, 0x12, 0x50, 0xbd, 0x60, 0x8e, 0x52, 0x22,
	0x17, 0xa7, 0x6f, 0x71, 0x7a, 0x68, 0x5e, 0x56, 0x62, 0x66, 0x8e, 0x90, 0x17, 0x17, 0x5f, 0x59,
	0x62, 0x4e, 0x66, 0x4a, 0x62, 0x49, 0x7e, 0x51, 0x7c, 0x62, 0x4a, 0x4a, 0x91, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0xa7, 0x93, 0xf2, 0xab, 0x7b, 0xf2, 0xec, 0x20, 0x7e, 0x6a, 0x71, 0xf1, 0xa5, 0x2d,
	0xba, 0x22, 0x50, 0x13, 0x1c, 0x21, 0x22, 0xc1, 0x25, 0x45, 0x99, 0x79, 0xe9, 0x41, 0xbc, 0x70,
	0xad, 0x20, 0x71, 0x2b, 0x8e, 0x8e, 0x05, 0xf2, 0x0c, 0x33, 0x16, 0xc8, 0x33, 0x2a, 0x09, 0x73,
	0x09, 0xc2, 0xad, 0x08, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x35, 0x8a, 0xe7, 0x62, 0xf6,
	0x2d, 0x4e, 0x17, 0x8a, 0xe0, 0x62, 0x83, 0xda, 0xad, 0xa4, 0x87, 0xc3, 0x4f, 0x7a, 0x70, 0xcd,
	0x52, 0x5a, 0x84, 0xd5, 0xc0, 0x2c, 0x70, 0xf2, 0x5e, 0xf1, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39,
	0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63,
	0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x74, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92,
	0xf3, 0x73, 0xa1, 0xe1, 0x01, 0xa5, 0x74, 0x8b, 0x53, 0xb2, 0xf5, 0x2b, 0x10, 0x41, 0x5f, 0x52,
	0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x0e, 0x2c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9f,
	0x32, 0xcd, 0x3e, 0x9a, 0x01, 0x00, 0x00,
}

func (this *MsgUnjail) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgUnjail)
	if !ok {
		that2, ok := that.(MsgUnjail)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ValidatorAddr != that1.ValidatorAddr {
		return false
	}
	return true
}
func (this *MsgUnjailResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgUnjailResponse)
	if !ok {
		that2, ok := that.(MsgUnjailResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// Unjail defines a method for unjailing a jailed validator, thus returning
	// them into the bonded validator set, so they can begin receiving provisions
	// and rewards again.
	Unjail(ctx context.Context, in *MsgUnjail, opts ...grpc.CallOption) (*MsgUnjailResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Unjail(ctx context.Context, in *MsgUnjail, opts ...grpc.CallOption) (*MsgUnjailResponse, error) {
	out := new(MsgUnjailResponse)
	err := c.cc.Invoke(ctx, "/cosmos.slashing.v1beta1.Msg/Unjail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Unjail defines a method for unjailing a jailed validator, thus returning
	// them into the bonded validator set, so they can begin receiving provisions
	// and rewards again.
	Unjail(context.Context, *MsgUnjail) (*MsgUnjailResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Unjail(ctx context.Context, req *MsgUnjail) (*MsgUnjailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unjail not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Unjail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnjail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Unjail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.slashing.v1beta1.Msg/Unjail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Unjail(ctx, req.(*MsgUnjail))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.slashing.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unjail",
			Handler:    _Msg_Unjail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/slashing/v1beta1/tx.proto",
}

func (m *MsgUnjail) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnjail) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnjail) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddr) > 0 {
		i -= len(m.ValidatorAddr)
		copy(dAtA[i:], m.ValidatorAddr)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ValidatorAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUnjailResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnjailResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnjailResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUnjail) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddr)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUnjailResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUnjail) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUnjail: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnjail: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUnjailResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUnjailResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnjailResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
