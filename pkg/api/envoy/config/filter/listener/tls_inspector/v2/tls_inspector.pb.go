// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/listener/tls_inspector/v2/tls_inspector.proto

package envoy_config_filter_listener_tls_inspector_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type TlsInspector struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TlsInspector) Reset()         { *m = TlsInspector{} }
func (m *TlsInspector) String() string { return proto.CompactTextString(m) }
func (*TlsInspector) ProtoMessage()    {}
func (*TlsInspector) Descriptor() ([]byte, []int) {
	return fileDescriptor_74d5b8d995215fa0, []int{0}
}
func (m *TlsInspector) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TlsInspector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TlsInspector.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TlsInspector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TlsInspector.Merge(m, src)
}
func (m *TlsInspector) XXX_Size() int {
	return m.Size()
}
func (m *TlsInspector) XXX_DiscardUnknown() {
	xxx_messageInfo_TlsInspector.DiscardUnknown(m)
}

var xxx_messageInfo_TlsInspector proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TlsInspector)(nil), "envoy.config.filter.listener.tls_inspector.v2.TlsInspector")
}

func init() {
	proto.RegisterFile("envoy/config/filter/listener/tls_inspector/v2/tls_inspector.proto", fileDescriptor_74d5b8d995215fa0)
}

var fileDescriptor_74d5b8d995215fa0 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4c, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0xc9, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x2d, 0xd2, 0x2f, 0xc9, 0x29, 0x8e, 0xcf, 0xcc, 0x2b,
	0x2e, 0x48, 0x4d, 0x2e, 0xc9, 0x2f, 0xd2, 0x2f, 0x33, 0x42, 0x15, 0xd0, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xd2, 0x05, 0x1b, 0xa1, 0x07, 0x31, 0x42, 0x0f, 0x62, 0x84, 0x1e, 0xcc, 0x08, 0x3d,
	0x54, 0x1d, 0x65, 0x46, 0x52, 0x72, 0xa5, 0x29, 0x05, 0x89, 0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25,
	0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0xfa, 0xb9, 0x99, 0xe9, 0x45, 0x89, 0x25, 0xa9, 0x10, 0xe3,
	0x94, 0xf8, 0xb8, 0x78, 0x42, 0x72, 0x8a, 0x3d, 0x61, 0x5a, 0x9c, 0x66, 0x30, 0x9e, 0x78, 0x24,
	0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x9f, 0x66, 0xfc, 0xeb, 0x67, 0x35,
	0x11, 0x32, 0x82, 0xd8, 0x99, 0x5a, 0x51, 0x92, 0x9a, 0x57, 0x0c, 0x32, 0x04, 0x6a, 0x6f, 0x31,
	0x4e, 0x8b, 0x8d, 0xb9, 0xac, 0x33, 0xf3, 0xf5, 0xc0, 0xda, 0x0a, 0x8a, 0xf2, 0x2b, 0x2a, 0xf5,
	0x48, 0x72, 0xb5, 0x93, 0x20, 0xb2, 0x9b, 0x02, 0x40, 0x0e, 0x0d, 0x60, 0x4c, 0x62, 0x03, 0xbb,
	0xd8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x0d, 0x9a, 0xb9, 0x45, 0x01, 0x00, 0x00,
}

func (m *TlsInspector) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TlsInspector) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TlsInspector) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintTlsInspector(dAtA []byte, offset int, v uint64) int {
	offset -= sovTlsInspector(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TlsInspector) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTlsInspector(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTlsInspector(x uint64) (n int) {
	return sovTlsInspector(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TlsInspector) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTlsInspector
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
			return fmt.Errorf("proto: TlsInspector: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TlsInspector: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTlsInspector(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTlsInspector
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTlsInspector
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
func skipTlsInspector(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTlsInspector
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
					return 0, ErrIntOverflowTlsInspector
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
					return 0, ErrIntOverflowTlsInspector
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
				return 0, ErrInvalidLengthTlsInspector
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTlsInspector
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTlsInspector
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
				next, err := skipTlsInspector(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTlsInspector
				}
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
	ErrInvalidLengthTlsInspector = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTlsInspector   = fmt.Errorf("proto: integer overflow")
)
