// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proxy/vless/encoding/addons.proto

package encoding // import "v2ray.com/core/proxy/vless/encoding"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Addons struct {
	Flow                 string   `protobuf:"bytes,1,opt,name=Flow,proto3" json:"Flow,omitempty"`
	Seed                 []byte   `protobuf:"bytes,2,opt,name=Seed,proto3" json:"Seed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Addons) Reset()         { *m = Addons{} }
func (m *Addons) String() string { return proto.CompactTextString(m) }
func (*Addons) ProtoMessage()    {}
func (*Addons) Descriptor() ([]byte, []int) {
	return fileDescriptor_addons_715144385dbf650f, []int{0}
}
func (m *Addons) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Addons) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Addons.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Addons) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Addons.Merge(dst, src)
}
func (m *Addons) XXX_Size() int {
	return m.Size()
}
func (m *Addons) XXX_DiscardUnknown() {
	xxx_messageInfo_Addons.DiscardUnknown(m)
}

var xxx_messageInfo_Addons proto.InternalMessageInfo

func (m *Addons) GetFlow() string {
	if m != nil {
		return m.Flow
	}
	return ""
}

func (m *Addons) GetSeed() []byte {
	if m != nil {
		return m.Seed
	}
	return nil
}

func init() {
	proto.RegisterType((*Addons)(nil), "v2ray.core.proxy.vless.encoding.Addons")
}
func (m *Addons) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Addons) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Flow) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAddons(dAtA, i, uint64(len(m.Flow)))
		i += copy(dAtA[i:], m.Flow)
	}
	if len(m.Seed) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAddons(dAtA, i, uint64(len(m.Seed)))
		i += copy(dAtA[i:], m.Seed)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintAddons(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Addons) Size() (n int) {
	var l int
	_ = l
	l = len(m.Flow)
	if l > 0 {
		n += 1 + l + sovAddons(uint64(l))
	}
	l = len(m.Seed)
	if l > 0 {
		n += 1 + l + sovAddons(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAddons(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAddons(x uint64) (n int) {
	return sovAddons(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Addons) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAddons
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Addons: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Addons: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flow", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddons
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAddons
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Flow = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seed", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddons
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAddons
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seed = append(m.Seed[:0], dAtA[iNdEx:postIndex]...)
			if m.Seed == nil {
				m.Seed = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAddons(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAddons
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAddons(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAddons
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
					return 0, ErrIntOverflowAddons
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAddons
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthAddons
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAddons
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAddons(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAddons = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAddons   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("proxy/vless/encoding/addons.proto", fileDescriptor_addons_715144385dbf650f)
}

var fileDescriptor_addons_715144385dbf650f = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x28, 0xca, 0xaf,
	0xa8, 0xd4, 0x2f, 0xcb, 0x49, 0x2d, 0x2e, 0xd6, 0x4f, 0xcd, 0x4b, 0xce, 0x4f, 0xc9, 0xcc, 0x4b,
	0xd7, 0x4f, 0x4c, 0x49, 0xc9, 0xcf, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x2f,
	0x33, 0x2a, 0x4a, 0xac, 0xd4, 0x4b, 0xce, 0x2f, 0x4a, 0xd5, 0x03, 0xab, 0xd6, 0x03, 0xab, 0xd6,
	0x83, 0xa9, 0x56, 0x32, 0xe0, 0x62, 0x73, 0x04, 0x6b, 0x10, 0x12, 0xe2, 0x62, 0x71, 0xcb, 0xc9,
	0x2f, 0x97, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x41, 0x62, 0xc1, 0xa9, 0xa9, 0x29,
	0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x60, 0xb6, 0x53, 0xdd, 0x89, 0x47, 0x72, 0x8c, 0x17,
	0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe3, 0xb1, 0x1c, 0x03, 0x97, 0x72, 0x72, 0x7e,
	0xae, 0x1e, 0x01, 0x8b, 0x02, 0x18, 0xa3, 0x94, 0x61, 0x4a, 0x72, 0xf5, 0x41, 0xca, 0xf4, 0xb1,
	0xb9, 0x7e, 0x15, 0x93, 0x7c, 0x98, 0x51, 0x50, 0x62, 0xa5, 0x9e, 0x33, 0xc8, 0xa0, 0x00, 0xb0,
	0x41, 0x61, 0x60, 0x83, 0x5c, 0xa1, 0x2a, 0x92, 0xd8, 0xc0, 0x3e, 0x33, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x36, 0x32, 0x14, 0x7c, 0xfe, 0x00, 0x00, 0x00,
}
