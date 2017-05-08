// Code generated by protoc-gen-gogo.
// source: trades.proto
// DO NOT EDIT!

/*
	Package trades is a generated protocol buffer package.

	It is generated from these files:
		trades.proto

	It has these top-level messages:
		Trade
		Coin
*/
package trades

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

type Trade struct {
	Tid    int64   `protobuf:"varint,1,opt,name=tid,proto3" json:"tid,omitempty"`
	Date   int64   `protobuf:"varint,2,opt,name=date,proto3" json:"date,omitempty"`
	Price  float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Amount float64 `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Type   string  `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
}

func (m *Trade) Reset()                    { *m = Trade{} }
func (m *Trade) String() string            { return proto.CompactTextString(m) }
func (*Trade) ProtoMessage()               {}
func (*Trade) Descriptor() ([]byte, []int) { return fileDescriptorTrades, []int{0} }

func (m *Trade) GetTid() int64 {
	if m != nil {
		return m.Tid
	}
	return 0
}

func (m *Trade) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *Trade) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Trade) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Trade) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

// Our address book file is just one of these.
type Coin struct {
	Trades []*Trade `protobuf:"bytes,1,rep,name=trades" json:"trades,omitempty"`
}

func (m *Coin) Reset()                    { *m = Coin{} }
func (m *Coin) String() string            { return proto.CompactTextString(m) }
func (*Coin) ProtoMessage()               {}
func (*Coin) Descriptor() ([]byte, []int) { return fileDescriptorTrades, []int{1} }

func (m *Coin) GetTrades() []*Trade {
	if m != nil {
		return m.Trades
	}
	return nil
}

func init() {
	proto.RegisterType((*Trade)(nil), "trades.Trade")
	proto.RegisterType((*Coin)(nil), "trades.Coin")
}
func (m *Trade) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Trade) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Tid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTrades(dAtA, i, uint64(m.Tid))
	}
	if m.Date != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTrades(dAtA, i, uint64(m.Date))
	}
	if m.Price != 0 {
		dAtA[i] = 0x19
		i++
		i = encodeFixed64Trades(dAtA, i, uint64(math.Float64bits(float64(m.Price))))
	}
	if m.Amount != 0 {
		dAtA[i] = 0x21
		i++
		i = encodeFixed64Trades(dAtA, i, uint64(math.Float64bits(float64(m.Amount))))
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintTrades(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	return i, nil
}

func (m *Coin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Coin) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Trades) > 0 {
		for _, msg := range m.Trades {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTrades(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Trades(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Trades(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTrades(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Trade) Size() (n int) {
	var l int
	_ = l
	if m.Tid != 0 {
		n += 1 + sovTrades(uint64(m.Tid))
	}
	if m.Date != 0 {
		n += 1 + sovTrades(uint64(m.Date))
	}
	if m.Price != 0 {
		n += 9
	}
	if m.Amount != 0 {
		n += 9
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovTrades(uint64(l))
	}
	return n
}

func (m *Coin) Size() (n int) {
	var l int
	_ = l
	if len(m.Trades) > 0 {
		for _, e := range m.Trades {
			l = e.Size()
			n += 1 + l + sovTrades(uint64(l))
		}
	}
	return n
}

func sovTrades(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTrades(x uint64) (n int) {
	return sovTrades(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Trade) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrades
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
			return fmt.Errorf("proto: Trade: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Trade: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tid", wireType)
			}
			m.Tid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrades
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			m.Date = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrades
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Date |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Price = float64(math.Float64frombits(v))
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Amount = float64(math.Float64frombits(v))
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrades
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
				return ErrInvalidLengthTrades
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTrades(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrades
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
func (m *Coin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrades
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
			return fmt.Errorf("proto: Coin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Coin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trades", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrades
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTrades
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Trades = append(m.Trades, &Trade{})
			if err := m.Trades[len(m.Trades)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTrades(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrades
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
func skipTrades(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTrades
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
					return 0, ErrIntOverflowTrades
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
					return 0, ErrIntOverflowTrades
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
				return 0, ErrInvalidLengthTrades
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTrades
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
				next, err := skipTrades(dAtA[start:])
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
	ErrInvalidLengthTrades = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTrades   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("trades.proto", fileDescriptorTrades) }

var fileDescriptorTrades = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x29, 0x4a, 0x4c,
	0x49, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0xf2, 0xb9, 0x58,
	0x43, 0x40, 0x2c, 0x21, 0x01, 0x2e, 0xe6, 0x92, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6,
	0x20, 0x10, 0x53, 0x48, 0x88, 0x8b, 0x25, 0x25, 0xb1, 0x24, 0x55, 0x82, 0x09, 0x2c, 0x04, 0x66,
	0x0b, 0x89, 0x70, 0xb1, 0x16, 0x14, 0x65, 0x26, 0xa7, 0x4a, 0x30, 0x2b, 0x30, 0x6a, 0x30, 0x06,
	0x41, 0x38, 0x42, 0x62, 0x5c, 0x6c, 0x89, 0xb9, 0xf9, 0xa5, 0x79, 0x25, 0x12, 0x2c, 0x60, 0x61,
	0x28, 0x0f, 0x64, 0x42, 0x49, 0x65, 0x41, 0xaa, 0x04, 0xab, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98,
	0xad, 0xa4, 0xcb, 0xc5, 0xe2, 0x9c, 0x9f, 0x99, 0x27, 0xa4, 0xca, 0x05, 0x75, 0x82, 0x04, 0xa3,
	0x02, 0xb3, 0x06, 0xb7, 0x11, 0xaf, 0x1e, 0xd4, 0x7d, 0x60, 0xe7, 0x04, 0x41, 0x25, 0x9d, 0x04,
	0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5,
	0x18, 0x92, 0xd8, 0xc0, 0x1e, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x25, 0x9a, 0xb0, 0x06,
	0xd0, 0x00, 0x00, 0x00,
}