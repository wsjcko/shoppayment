// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shoppayment.proto

package go_micro_service_shop_payment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PaymentInfo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PaymentName          string   `protobuf:"bytes,2,opt,name=payment_name,json=paymentName,proto3" json:"payment_name,omitempty"`
	PaymentSid           string   `protobuf:"bytes,3,opt,name=payment_sid,json=paymentSid,proto3" json:"payment_sid,omitempty"`
	PaymentStatus        bool     `protobuf:"varint,4,opt,name=payment_status,json=paymentStatus,proto3" json:"payment_status,omitempty"`
	PaymentImage         string   `protobuf:"bytes,5,opt,name=Payment_image,json=PaymentImage,proto3" json:"Payment_image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentInfo) Reset()         { *m = PaymentInfo{} }
func (m *PaymentInfo) String() string { return proto.CompactTextString(m) }
func (*PaymentInfo) ProtoMessage()    {}
func (*PaymentInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_997f580f5dec5927, []int{0}
}
func (m *PaymentInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PaymentInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PaymentInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PaymentInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentInfo.Merge(m, src)
}
func (m *PaymentInfo) XXX_Size() int {
	return m.Size()
}
func (m *PaymentInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentInfo proto.InternalMessageInfo

func (m *PaymentInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PaymentInfo) GetPaymentName() string {
	if m != nil {
		return m.PaymentName
	}
	return ""
}

func (m *PaymentInfo) GetPaymentSid() string {
	if m != nil {
		return m.PaymentSid
	}
	return ""
}

func (m *PaymentInfo) GetPaymentStatus() bool {
	if m != nil {
		return m.PaymentStatus
	}
	return false
}

func (m *PaymentInfo) GetPaymentImage() string {
	if m != nil {
		return m.PaymentImage
	}
	return ""
}

type PaymentID struct {
	PaymentId            int64    `protobuf:"varint,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentID) Reset()         { *m = PaymentID{} }
func (m *PaymentID) String() string { return proto.CompactTextString(m) }
func (*PaymentID) ProtoMessage()    {}
func (*PaymentID) Descriptor() ([]byte, []int) {
	return fileDescriptor_997f580f5dec5927, []int{1}
}
func (m *PaymentID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PaymentID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PaymentID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PaymentID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentID.Merge(m, src)
}
func (m *PaymentID) XXX_Size() int {
	return m.Size()
}
func (m *PaymentID) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentID.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentID proto.InternalMessageInfo

func (m *PaymentID) GetPaymentId() int64 {
	if m != nil {
		return m.PaymentId
	}
	return 0
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_997f580f5dec5927, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type All struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *All) Reset()         { *m = All{} }
func (m *All) String() string { return proto.CompactTextString(m) }
func (*All) ProtoMessage()    {}
func (*All) Descriptor() ([]byte, []int) {
	return fileDescriptor_997f580f5dec5927, []int{3}
}
func (m *All) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *All) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_All.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *All) XXX_Merge(src proto.Message) {
	xxx_messageInfo_All.Merge(m, src)
}
func (m *All) XXX_Size() int {
	return m.Size()
}
func (m *All) XXX_DiscardUnknown() {
	xxx_messageInfo_All.DiscardUnknown(m)
}

var xxx_messageInfo_All proto.InternalMessageInfo

type PaymentAll struct {
	PaymentInfo          []*PaymentInfo `protobuf:"bytes,1,rep,name=payment_info,json=paymentInfo,proto3" json:"payment_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PaymentAll) Reset()         { *m = PaymentAll{} }
func (m *PaymentAll) String() string { return proto.CompactTextString(m) }
func (*PaymentAll) ProtoMessage()    {}
func (*PaymentAll) Descriptor() ([]byte, []int) {
	return fileDescriptor_997f580f5dec5927, []int{4}
}
func (m *PaymentAll) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PaymentAll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PaymentAll.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PaymentAll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentAll.Merge(m, src)
}
func (m *PaymentAll) XXX_Size() int {
	return m.Size()
}
func (m *PaymentAll) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentAll.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentAll proto.InternalMessageInfo

func (m *PaymentAll) GetPaymentInfo() []*PaymentInfo {
	if m != nil {
		return m.PaymentInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*PaymentInfo)(nil), "go.micro.service.shop.payment.PaymentInfo")
	proto.RegisterType((*PaymentID)(nil), "go.micro.service.shop.payment.PaymentID")
	proto.RegisterType((*Response)(nil), "go.micro.service.shop.payment.Response")
	proto.RegisterType((*All)(nil), "go.micro.service.shop.payment.All")
	proto.RegisterType((*PaymentAll)(nil), "go.micro.service.shop.payment.PaymentAll")
}

func init() { proto.RegisterFile("shoppayment.proto", fileDescriptor_997f580f5dec5927) }

var fileDescriptor_997f580f5dec5927 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0xae, 0xd2, 0x40,
	0x18, 0xc5, 0x3b, 0x14, 0x0c, 0xfd, 0x0a, 0x08, 0xb3, 0x6a, 0x88, 0xd4, 0x3a, 0xc6, 0x58, 0x59,
	0x74, 0x81, 0x4f, 0x50, 0xd2, 0x98, 0xb0, 0xd0, 0x98, 0x12, 0x57, 0x2e, 0x9a, 0xca, 0x0c, 0x30,
	0xc9, 0xb4, 0xd3, 0xd0, 0x6a, 0xc2, 0x73, 0xb8, 0xf1, 0x25, 0x7c, 0x0f, 0x97, 0x3e, 0x82, 0xc1,
	0x17, 0xb9, 0x69, 0xef, 0x4c, 0xef, 0x5d, 0x5d, 0xfe, 0xec, 0xe6, 0x3b, 0xf9, 0xf5, 0x3b, 0x67,
	0x4e, 0x3a, 0x30, 0x29, 0xf7, 0xb2, 0x28, 0xd2, 0x63, 0xc6, 0xf2, 0x2a, 0x28, 0x0e, 0xb2, 0x92,
	0x78, 0xb6, 0x93, 0x41, 0xc6, 0x37, 0x07, 0x19, 0x94, 0xec, 0xf0, 0x83, 0x6f, 0x58, 0x50, 0x33,
	0x81, 0x82, 0xc8, 0x6f, 0x04, 0xf6, 0xe7, 0xfb, 0xf3, 0x2a, 0xdf, 0x4a, 0x3c, 0x82, 0x0e, 0xa7,
	0x0e, 0xf2, 0x90, 0x6f, 0xc6, 0x1d, 0x4e, 0xf1, 0x2b, 0x18, 0x28, 0x34, 0xc9, 0xd3, 0x8c, 0x39,
	0x1d, 0x0f, 0xf9, 0x56, 0x6c, 0x2b, 0xed, 0x53, 0x9a, 0x31, 0xfc, 0x12, 0xf4, 0x98, 0x94, 0x9c,
	0x3a, 0x66, 0x43, 0x80, 0x92, 0xd6, 0x9c, 0xe2, 0x37, 0x30, 0x6a, 0x81, 0x2a, 0xad, 0xbe, 0x97,
	0x4e, 0xd7, 0x43, 0x7e, 0x3f, 0x1e, 0x6a, 0xa6, 0x11, 0xf1, 0x6b, 0x18, 0xaa, 0x24, 0x09, 0xcf,
	0xd2, 0x1d, 0x73, 0x7a, 0xcd, 0xa6, 0x81, 0x8e, 0x57, 0x6b, 0x64, 0x0e, 0x96, 0x9e, 0x23, 0x3c,
	0x03, 0x6d, 0x93, 0xb4, 0xa1, 0x2d, 0xa5, 0xac, 0x28, 0x79, 0x01, 0xfd, 0x98, 0x95, 0x85, 0xcc,
	0x4b, 0x86, 0xc7, 0x60, 0x66, 0xe5, 0xae, 0x61, 0xac, 0xb8, 0x3e, 0x92, 0x1e, 0x98, 0xa1, 0x10,
	0xe4, 0x2b, 0x80, 0x5a, 0x18, 0x0a, 0x81, 0x3f, 0x3e, 0x5c, 0x97, 0xe7, 0x5b, 0xe9, 0x20, 0xcf,
	0xf4, 0xed, 0xc5, 0x3c, 0x78, 0xb2, 0xc4, 0xe0, 0x51, 0x81, 0x6d, 0x35, 0xf5, 0xb0, 0xf8, 0xd9,
	0x05, 0x7b, 0xbd, 0x97, 0x85, 0x02, 0x30, 0x05, 0x08, 0x29, 0xd5, 0xd3, 0x15, 0x6b, 0xa7, 0xfe,
	0x85, 0x6c, 0x44, 0x0c, 0xbc, 0x85, 0xe1, 0x97, 0x82, 0xa6, 0x15, 0xbb, 0xc5, 0xe8, 0xed, 0x19,
	0x56, 0x37, 0x4a, 0x0c, 0xbc, 0x87, 0x49, 0xc4, 0x04, 0x6b, 0x7d, 0x96, 0xc7, 0x55, 0x84, 0x2f,
	0x0e, 0x7a, 0x8d, 0x13, 0x87, 0xe7, 0x1f, 0x78, 0x4e, 0x6f, 0xf3, 0xb9, 0xe2, 0xf6, 0xc4, 0xc0,
	0x29, 0x8c, 0x6a, 0xab, 0x50, 0x08, 0xdd, 0x1e, 0x39, 0xf3, 0x7d, 0x28, 0xc4, 0xf4, 0xdd, 0x65,
	0x1e, 0xf5, 0x0f, 0x67, 0x2c, 0xc7, 0x7f, 0x4e, 0x2e, 0xfa, 0x7b, 0x72, 0xd1, 0xbf, 0x93, 0x8b,
	0x7e, 0xfd, 0x77, 0x8d, 0x6f, 0xcf, 0x9a, 0xb7, 0xfa, 0xfe, 0x2e, 0x00, 0x00, 0xff, 0xff, 0xba,
	0xbd, 0x31, 0x23, 0xc0, 0x03, 0x00, 0x00,
}

func (m *PaymentInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PaymentInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PaymentInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.PaymentImage) > 0 {
		i -= len(m.PaymentImage)
		copy(dAtA[i:], m.PaymentImage)
		i = encodeVarintShoppayment(dAtA, i, uint64(len(m.PaymentImage)))
		i--
		dAtA[i] = 0x2a
	}
	if m.PaymentStatus {
		i--
		if m.PaymentStatus {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.PaymentSid) > 0 {
		i -= len(m.PaymentSid)
		copy(dAtA[i:], m.PaymentSid)
		i = encodeVarintShoppayment(dAtA, i, uint64(len(m.PaymentSid)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PaymentName) > 0 {
		i -= len(m.PaymentName)
		copy(dAtA[i:], m.PaymentName)
		i = encodeVarintShoppayment(dAtA, i, uint64(len(m.PaymentName)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintShoppayment(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PaymentID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PaymentID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PaymentID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.PaymentId != 0 {
		i = encodeVarintShoppayment(dAtA, i, uint64(m.PaymentId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintShoppayment(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *All) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *All) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *All) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *PaymentAll) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PaymentAll) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PaymentAll) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.PaymentInfo) > 0 {
		for iNdEx := len(m.PaymentInfo) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PaymentInfo[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintShoppayment(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintShoppayment(dAtA []byte, offset int, v uint64) int {
	offset -= sovShoppayment(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PaymentInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovShoppayment(uint64(m.Id))
	}
	l = len(m.PaymentName)
	if l > 0 {
		n += 1 + l + sovShoppayment(uint64(l))
	}
	l = len(m.PaymentSid)
	if l > 0 {
		n += 1 + l + sovShoppayment(uint64(l))
	}
	if m.PaymentStatus {
		n += 2
	}
	l = len(m.PaymentImage)
	if l > 0 {
		n += 1 + l + sovShoppayment(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PaymentID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PaymentId != 0 {
		n += 1 + sovShoppayment(uint64(m.PaymentId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovShoppayment(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *All) Size() (n int) {
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

func (m *PaymentAll) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PaymentInfo) > 0 {
		for _, e := range m.PaymentInfo {
			l = e.Size()
			n += 1 + l + sovShoppayment(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovShoppayment(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozShoppayment(x uint64) (n int) {
	return sovShoppayment(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PaymentInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShoppayment
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
			return fmt.Errorf("proto: PaymentInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PaymentInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShoppayment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShoppayment
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
				return ErrInvalidLengthShoppayment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShoppayment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PaymentName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentSid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShoppayment
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
				return ErrInvalidLengthShoppayment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShoppayment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PaymentSid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentStatus", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShoppayment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PaymentStatus = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentImage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShoppayment
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
				return ErrInvalidLengthShoppayment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShoppayment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PaymentImage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShoppayment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShoppayment
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
func (m *PaymentID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShoppayment
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
			return fmt.Errorf("proto: PaymentID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PaymentID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentId", wireType)
			}
			m.PaymentId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShoppayment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PaymentId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShoppayment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShoppayment
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
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShoppayment
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
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShoppayment
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
				return ErrInvalidLengthShoppayment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShoppayment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShoppayment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShoppayment
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
func (m *All) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShoppayment
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
			return fmt.Errorf("proto: All: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: All: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipShoppayment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShoppayment
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
func (m *PaymentAll) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShoppayment
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
			return fmt.Errorf("proto: PaymentAll: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PaymentAll: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShoppayment
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
				return ErrInvalidLengthShoppayment
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthShoppayment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PaymentInfo = append(m.PaymentInfo, &PaymentInfo{})
			if err := m.PaymentInfo[len(m.PaymentInfo)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShoppayment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShoppayment
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
func skipShoppayment(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowShoppayment
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
					return 0, ErrIntOverflowShoppayment
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
					return 0, ErrIntOverflowShoppayment
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
				return 0, ErrInvalidLengthShoppayment
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupShoppayment
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthShoppayment
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthShoppayment        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowShoppayment          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupShoppayment = fmt.Errorf("proto: unexpected end of group")
)
