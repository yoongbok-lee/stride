// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/stakeibc/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the stakeibc module's genesis state.
type GenesisState struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PortId string `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	// list of zones that are registered by the protocol
	HostZoneList     []HostZone     `protobuf:"bytes,5,rep,name=host_zone_list,json=hostZoneList,proto3" json:"host_zone_list"`
	EpochTrackerList []EpochTracker `protobuf:"bytes,10,rep,name=epoch_tracker_list,json=epochTrackerList,proto3" json:"epoch_tracker_list"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea81129ed6fb77a, []int{0}
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

func (m *GenesisState) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *GenesisState) GetHostZoneList() []HostZone {
	if m != nil {
		return m.HostZoneList
	}
	return nil
}

func (m *GenesisState) GetEpochTrackerList() []EpochTracker {
	if m != nil {
		return m.EpochTrackerList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "stride.stakeibc.GenesisState")
}

func init() { proto.RegisterFile("stride/stakeibc/genesis.proto", fileDescriptor_dea81129ed6fb77a) }

var fileDescriptor_dea81129ed6fb77a = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcd, 0x4e, 0xc2, 0x40,
	0x14, 0x85, 0x5b, 0x18, 0x4a, 0x19, 0x88, 0x36, 0x8d, 0x09, 0x48, 0xa4, 0x10, 0xdd, 0xb0, 0xb1,
	0x8d, 0x18, 0x7d, 0x00, 0x12, 0xa2, 0x56, 0x16, 0x0a, 0xae, 0xd8, 0x34, 0x6d, 0x99, 0xb4, 0x0d,
	0xc2, 0x34, 0x9d, 0xab, 0x51, 0x9f, 0xc2, 0x95, 0xcf, 0xc4, 0x92, 0xa5, 0x2b, 0x63, 0xe0, 0x45,
	0x4c, 0x3b, 0xe3, 0x5f, 0xd9, 0xcd, 0x9d, 0xf3, 0xe5, 0xbb, 0x27, 0x17, 0xb7, 0x18, 0x24, 0xd1,
	0x94, 0x58, 0x0c, 0xdc, 0x19, 0x89, 0x3c, 0xdf, 0x0a, 0xc8, 0x82, 0xb0, 0x88, 0x99, 0x71, 0x42,
	0x81, 0xea, 0xbb, 0x3c, 0x36, 0xbf, 0xe3, 0xe6, 0x5e, 0x40, 0x03, 0x9a, 0x65, 0x56, 0xfa, 0xe2,
	0x58, 0xf3, 0x20, 0x6f, 0x89, 0xdd, 0xc4, 0x9d, 0x0b, 0x49, 0xb3, 0x9d, 0x4f, 0x43, 0xca, 0xc0,
	0x79, 0xa1, 0x0b, 0x22, 0x80, 0xa3, 0x3c, 0x40, 0x62, 0xea, 0x87, 0x0e, 0x24, 0xae, 0x3f, 0x23,
	0x09, 0x87, 0x0e, 0xdf, 0x0a, 0xb8, 0x76, 0xc1, 0xcb, 0x8d, 0xc1, 0x05, 0xa2, 0x9f, 0x61, 0x85,
	0xaf, 0x69, 0xc8, 0x1d, 0xb9, 0x5b, 0xed, 0xd5, 0xcd, 0x5c, 0x59, 0xf3, 0x26, 0x8b, 0xfb, 0x68,
	0xf9, 0xd1, 0x96, 0x46, 0x02, 0xd6, 0xeb, 0xb8, 0x1c, 0xd3, 0x04, 0x9c, 0x68, 0xda, 0x28, 0x74,
	0xe4, 0x6e, 0x65, 0xa4, 0xa4, 0xe3, 0xd5, 0x54, 0x1f, 0xe0, 0x9d, 0x9f, 0x62, 0xce, 0x7d, 0xc4,
	0xa0, 0x51, 0xea, 0x14, 0xbb, 0xd5, 0xde, 0xfe, 0x96, 0xf7, 0x92, 0x32, 0x98, 0xd0, 0x05, 0x11,
	0xe6, 0x5a, 0x28, 0xe6, 0x61, 0xc4, 0x40, 0xbf, 0xc5, 0xfa, 0xbf, 0xfa, 0x5c, 0x85, 0x33, 0x55,
	0x6b, 0x4b, 0x35, 0x48, 0xd1, 0x3b, 0x4e, 0x0a, 0x9d, 0x46, 0xfe, 0xfc, 0xa5, 0x4a, 0x1b, 0xa9,
	0x45, 0x0d, 0xd9, 0x48, 0x45, 0x5a, 0xc9, 0x46, 0xaa, 0xa2, 0x95, 0x6d, 0xa4, 0x56, 0x34, 0x6c,
	0x23, 0xb5, 0xaa, 0xd5, 0xfa, 0xd7, 0xcb, 0xb5, 0x21, 0xaf, 0xd6, 0x86, 0xfc, 0xb9, 0x36, 0xe4,
	0xd7, 0x8d, 0x21, 0xad, 0x36, 0x86, 0xf4, 0xbe, 0x31, 0xa4, 0xc9, 0x49, 0x10, 0x41, 0xf8, 0xe0,
	0x99, 0x3e, 0x9d, 0x5b, 0xe3, 0x6c, 0xf1, 0xf1, 0xd0, 0xf5, 0x98, 0x25, 0xce, 0xfd, 0x78, 0x6e,
	0x3d, 0xfd, 0xde, 0x1c, 0x9e, 0x63, 0xc2, 0x3c, 0x25, 0x3b, 0xf6, 0xe9, 0x57, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x53, 0xc7, 0x78, 0x6e, 0x18, 0x02, 0x00, 0x00,
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
	if len(m.EpochTrackerList) > 0 {
		for iNdEx := len(m.EpochTrackerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EpochTrackerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.HostZoneList) > 0 {
		for iNdEx := len(m.HostZoneList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HostZoneList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
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
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.HostZoneList) > 0 {
		for _, e := range m.HostZoneList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.EpochTrackerList) > 0 {
		for _, e := range m.EpochTrackerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
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
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostZoneList", wireType)
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
			m.HostZoneList = append(m.HostZoneList, HostZone{})
			if err := m.HostZoneList[len(m.HostZoneList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochTrackerList", wireType)
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
			m.EpochTrackerList = append(m.EpochTrackerList, EpochTracker{})
			if err := m.EpochTrackerList[len(m.EpochTrackerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
