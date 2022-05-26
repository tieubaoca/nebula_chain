// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: launchpad/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// GenesisState defines the launchpad module's genesis state.
type GenesisState struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// initial Project at genesis
	Projects []*Project `protobuf:"bytes,2,rep,name=projects,proto3" json:"projects,omitempty"`
	// global project id. The start of global project id will start from the number of project at genesis.
	GlobalProjectIdStart uint64 `protobuf:"varint,3,opt,name=global_project_id_start,json=globalProjectIdStart,proto3" json:"global_project_id_start,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_555ea8fa49503c1a, []int{0}
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

func (m *GenesisState) GetProjects() []*Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

func (m *GenesisState) GetGlobalProjectIdStart() uint64 {
	if m != nil {
		return m.GlobalProjectIdStart
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "nghuyenthevinh2000.nebula.launchpad.GenesisState")
}

func init() { proto.RegisterFile("launchpad/genesis.proto", fileDescriptor_555ea8fa49503c1a) }

var fileDescriptor_555ea8fa49503c1a = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0x49, 0x2c, 0xcd,
	0x4b, 0xce, 0x28, 0x48, 0x4c, 0xd1, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0xce, 0x2c, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x52, 0xce, 0x4b, 0xcf, 0x28, 0xad, 0x4c, 0xcd, 0x2b, 0xc9, 0x48, 0x2d,
	0xcb, 0xcc, 0xcb, 0x30, 0x32, 0x30, 0x30, 0xd0, 0xcb, 0x4b, 0x4d, 0x2a, 0xcd, 0x49, 0xd4, 0x83,
	0x6b, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xab, 0xd7, 0x07, 0xb1, 0x20, 0x5a, 0xa5, 0xc4,
	0x10, 0x66, 0x16, 0x24, 0x16, 0x25, 0xe6, 0x42, 0x8d, 0x94, 0x42, 0xb2, 0xab, 0xa0, 0x28, 0x3f,
	0x2b, 0x35, 0xb9, 0x04, 0x22, 0xa1, 0x74, 0x8b, 0x91, 0x8b, 0xc7, 0x1d, 0x62, 0x7b, 0x70, 0x49,
	0x62, 0x49, 0xaa, 0x90, 0x27, 0x17, 0x1b, 0x44, 0xa7, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91,
	0xb6, 0x1e, 0x11, 0xae, 0xd1, 0x0b, 0x00, 0x6b, 0x71, 0x62, 0x39, 0x71, 0x4f, 0x9e, 0x21, 0x08,
	0x6a, 0x80, 0x90, 0x07, 0x17, 0x07, 0xd4, 0xb2, 0x62, 0x09, 0x26, 0x05, 0x66, 0x0d, 0x6e, 0x23,
	0x1d, 0xe2, 0x0c, 0x83, 0x68, 0x0a, 0x82, 0xeb, 0x16, 0x32, 0xe5, 0x12, 0x4f, 0xcf, 0xc9, 0x4f,
	0x4a, 0xcc, 0x89, 0x87, 0x0a, 0xc5, 0x67, 0xa6, 0xc4, 0x17, 0x97, 0x24, 0x16, 0x95, 0x48, 0x30,
	0x2b, 0x30, 0x6a, 0xb0, 0x04, 0x89, 0x40, 0xa4, 0xa1, 0x3a, 0x3d, 0x53, 0x82, 0x41, 0x72, 0x4e,
	0x01, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7,
	0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x96, 0x9e, 0x59, 0x92,
	0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x8f, 0xe9, 0x24, 0x7d, 0x88, 0x93, 0xf4, 0x2b, 0xf4,
	0x11, 0xc1, 0x56, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x0e, 0x35, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x45, 0x37, 0x74, 0xc5, 0xbc, 0x01, 0x00, 0x00,
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
	if m.GlobalProjectIdStart != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.GlobalProjectIdStart))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Projects) > 0 {
		for iNdEx := len(m.Projects) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Projects[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
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
	if len(m.Projects) > 0 {
		for _, e := range m.Projects {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.GlobalProjectIdStart != 0 {
		n += 1 + sovGenesis(uint64(m.GlobalProjectIdStart))
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
				return fmt.Errorf("proto: wrong wireType = %d for field Projects", wireType)
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
			m.Projects = append(m.Projects, &Project{})
			if err := m.Projects[len(m.Projects)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalProjectIdStart", wireType)
			}
			m.GlobalProjectIdStart = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GlobalProjectIdStart |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
