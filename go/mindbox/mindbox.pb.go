// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/mindbox.proto

package mindbox

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type PingResponse struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1ddbb12eed38be6, []int{0}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type ParamsUserInformation struct {
	ClientId             int32    `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParamsUserInformation) Reset()         { *m = ParamsUserInformation{} }
func (m *ParamsUserInformation) String() string { return proto.CompactTextString(m) }
func (*ParamsUserInformation) ProtoMessage()    {}
func (*ParamsUserInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1ddbb12eed38be6, []int{1}
}

func (m *ParamsUserInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParamsUserInformation.Unmarshal(m, b)
}
func (m *ParamsUserInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParamsUserInformation.Marshal(b, m, deterministic)
}
func (m *ParamsUserInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamsUserInformation.Merge(m, src)
}
func (m *ParamsUserInformation) XXX_Size() int {
	return xxx_messageInfo_ParamsUserInformation.Size(m)
}
func (m *ParamsUserInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamsUserInformation.DiscardUnknown(m)
}

var xxx_messageInfo_ParamsUserInformation proto.InternalMessageInfo

func (m *ParamsUserInformation) GetClientId() int32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

type ResponseUserInformation struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	BonusTotal           int32    `protobuf:"varint,5,opt,name=bonus_total,json=bonusTotal,proto3" json:"bonus_total,omitempty"`
	BonusAvailable       int32    `protobuf:"varint,6,opt,name=bonus_available,json=bonusAvailable,proto3" json:"bonus_available,omitempty"`
	BonusBlocked         int32    `protobuf:"varint,7,opt,name=bonus_blocked,json=bonusBlocked,proto3" json:"bonus_blocked,omitempty"`
	TotalPaidAmount      float64  `protobuf:"fixed64,8,opt,name=total_paid_amount,json=totalPaidAmount,proto3" json:"total_paid_amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseUserInformation) Reset()         { *m = ResponseUserInformation{} }
func (m *ResponseUserInformation) String() string { return proto.CompactTextString(m) }
func (*ResponseUserInformation) ProtoMessage()    {}
func (*ResponseUserInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1ddbb12eed38be6, []int{2}
}

func (m *ResponseUserInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseUserInformation.Unmarshal(m, b)
}
func (m *ResponseUserInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseUserInformation.Marshal(b, m, deterministic)
}
func (m *ResponseUserInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseUserInformation.Merge(m, src)
}
func (m *ResponseUserInformation) XXX_Size() int {
	return xxx_messageInfo_ResponseUserInformation.Size(m)
}
func (m *ResponseUserInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseUserInformation.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseUserInformation proto.InternalMessageInfo

func (m *ResponseUserInformation) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *ResponseUserInformation) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *ResponseUserInformation) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ResponseUserInformation) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *ResponseUserInformation) GetBonusTotal() int32 {
	if m != nil {
		return m.BonusTotal
	}
	return 0
}

func (m *ResponseUserInformation) GetBonusAvailable() int32 {
	if m != nil {
		return m.BonusAvailable
	}
	return 0
}

func (m *ResponseUserInformation) GetBonusBlocked() int32 {
	if m != nil {
		return m.BonusBlocked
	}
	return 0
}

func (m *ResponseUserInformation) GetTotalPaidAmount() float64 {
	if m != nil {
		return m.TotalPaidAmount
	}
	return 0
}

type ParamsOrdersHistory struct {
	ClientId             int32    `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParamsOrdersHistory) Reset()         { *m = ParamsOrdersHistory{} }
func (m *ParamsOrdersHistory) String() string { return proto.CompactTextString(m) }
func (*ParamsOrdersHistory) ProtoMessage()    {}
func (*ParamsOrdersHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1ddbb12eed38be6, []int{3}
}

func (m *ParamsOrdersHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParamsOrdersHistory.Unmarshal(m, b)
}
func (m *ParamsOrdersHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParamsOrdersHistory.Marshal(b, m, deterministic)
}
func (m *ParamsOrdersHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamsOrdersHistory.Merge(m, src)
}
func (m *ParamsOrdersHistory) XXX_Size() int {
	return xxx_messageInfo_ParamsOrdersHistory.Size(m)
}
func (m *ParamsOrdersHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamsOrdersHistory.DiscardUnknown(m)
}

var xxx_messageInfo_ParamsOrdersHistory proto.InternalMessageInfo

func (m *ParamsOrdersHistory) GetClientId() int32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

type ResponseOrdersHistory struct {
	TotalCount           int32    `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Orders               []*Order `protobuf:"bytes,2,rep,name=orders,proto3" json:"orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseOrdersHistory) Reset()         { *m = ResponseOrdersHistory{} }
func (m *ResponseOrdersHistory) String() string { return proto.CompactTextString(m) }
func (*ResponseOrdersHistory) ProtoMessage()    {}
func (*ResponseOrdersHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1ddbb12eed38be6, []int{4}
}

func (m *ResponseOrdersHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseOrdersHistory.Unmarshal(m, b)
}
func (m *ResponseOrdersHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseOrdersHistory.Marshal(b, m, deterministic)
}
func (m *ResponseOrdersHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseOrdersHistory.Merge(m, src)
}
func (m *ResponseOrdersHistory) XXX_Size() int {
	return xxx_messageInfo_ResponseOrdersHistory.Size(m)
}
func (m *ResponseOrdersHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseOrdersHistory.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseOrdersHistory proto.InternalMessageInfo

func (m *ResponseOrdersHistory) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ResponseOrdersHistory) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

type Order struct {
	Id                    string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedDate           string   `protobuf:"bytes,2,opt,name=created_date,json=createdDate,proto3" json:"created_date,omitempty"`
	PaymentType           string   `protobuf:"bytes,3,opt,name=payment_type,json=paymentType,proto3" json:"payment_type,omitempty"`
	DiscountedTotalPrice  float64  `protobuf:"fixed64,4,opt,name=discounted_total_price,json=discountedTotalPrice,proto3" json:"discounted_total_price,omitempty"`
	PaymentAmount         float64  `protobuf:"fixed64,5,opt,name=payment_amount,json=paymentAmount,proto3" json:"payment_amount,omitempty"`
	AppliedDiscount       float64  `protobuf:"fixed64,6,opt,name=applied_discount,json=appliedDiscount,proto3" json:"applied_discount,omitempty"`
	AcquiredBalanceChange float64  `protobuf:"fixed64,7,opt,name=acquired_balance_change,json=acquiredBalanceChange,proto3" json:"acquired_balance_change,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1ddbb12eed38be6, []int{5}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetCreatedDate() string {
	if m != nil {
		return m.CreatedDate
	}
	return ""
}

func (m *Order) GetPaymentType() string {
	if m != nil {
		return m.PaymentType
	}
	return ""
}

func (m *Order) GetDiscountedTotalPrice() float64 {
	if m != nil {
		return m.DiscountedTotalPrice
	}
	return 0
}

func (m *Order) GetPaymentAmount() float64 {
	if m != nil {
		return m.PaymentAmount
	}
	return 0
}

func (m *Order) GetAppliedDiscount() float64 {
	if m != nil {
		return m.AppliedDiscount
	}
	return 0
}

func (m *Order) GetAcquiredBalanceChange() float64 {
	if m != nil {
		return m.AcquiredBalanceChange
	}
	return 0
}

func init() {
	proto.RegisterType((*PingResponse)(nil), "mindbox.PingResponse")
	proto.RegisterType((*ParamsUserInformation)(nil), "mindbox.ParamsUserInformation")
	proto.RegisterType((*ResponseUserInformation)(nil), "mindbox.ResponseUserInformation")
	proto.RegisterType((*ParamsOrdersHistory)(nil), "mindbox.ParamsOrdersHistory")
	proto.RegisterType((*ResponseOrdersHistory)(nil), "mindbox.ResponseOrdersHistory")
	proto.RegisterType((*Order)(nil), "mindbox.order")
}

func init() { proto.RegisterFile("proto/mindbox.proto", fileDescriptor_c1ddbb12eed38be6) }

var fileDescriptor_c1ddbb12eed38be6 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xdd, 0x52, 0xd3, 0x40,
	0x14, 0xc7, 0x49, 0xa0, 0x40, 0x4f, 0x4b, 0xab, 0x0b, 0x85, 0x4c, 0x51, 0xa8, 0x71, 0x54, 0xe4,
	0xa2, 0xcc, 0x20, 0xa3, 0xd7, 0x7c, 0x38, 0x23, 0x17, 0x4a, 0x27, 0x03, 0x37, 0xde, 0xc4, 0x4d,
	0x76, 0x29, 0x3b, 0x26, 0xbb, 0x71, 0xb3, 0x75, 0xec, 0xbb, 0xe8, 0x0b, 0xfa, 0x14, 0x4e, 0xce,
	0x6e, 0x8a, 0xad, 0x1f, 0x77, 0x39, 0xbf, 0xff, 0xf9, 0xea, 0x39, 0x7b, 0x0a, 0x9b, 0x85, 0x56,
	0x46, 0x1d, 0xe5, 0x42, 0xb2, 0x44, 0x7d, 0x1b, 0xa2, 0x45, 0xd6, 0x9c, 0xd9, 0xdf, 0x1d, 0x2b,
	0x35, 0xce, 0xf8, 0x11, 0xe2, 0x64, 0x72, 0x7b, 0xc4, 0xf3, 0xc2, 0x4c, 0xad, 0x57, 0x78, 0x08,
	0xed, 0x91, 0x90, 0xe3, 0x88, 0x97, 0x85, 0x92, 0x25, 0x27, 0x7d, 0x58, 0xd7, 0xee, 0x3b, 0xf0,
	0x06, 0xde, 0x41, 0x33, 0x9a, 0xd9, 0xe1, 0x09, 0xf4, 0x46, 0x54, 0xd3, 0xbc, 0xbc, 0x29, 0xb9,
	0xbe, 0x94, 0xb7, 0x4a, 0xe7, 0xd4, 0x08, 0x25, 0xc9, 0x2e, 0x34, 0xd3, 0x4c, 0x70, 0x69, 0x62,
	0xc1, 0x30, 0xaa, 0x11, 0xad, 0x5b, 0x70, 0xc9, 0xc2, 0xef, 0x3e, 0xec, 0xd4, 0xe9, 0x17, 0x03,
	0x1f, 0x03, 0xdc, 0x0a, 0x5d, 0x9a, 0x58, 0xd2, 0xbc, 0xae, 0xd7, 0x44, 0xf2, 0x81, 0xe6, 0xbc,
	0xca, 0x9b, 0xd1, 0x5a, 0xf5, 0x6d, 0x37, 0x15, 0x40, 0x71, 0x0b, 0x1a, 0x3c, 0xa7, 0x22, 0x0b,
	0x96, 0x51, 0xb0, 0x46, 0x45, 0x8b, 0x3b, 0x25, 0x79, 0xb0, 0x62, 0x29, 0x1a, 0x64, 0x1f, 0x5a,
	0x89, 0x92, 0x93, 0x32, 0x36, 0xca, 0xd0, 0x2c, 0x68, 0x60, 0x8b, 0x80, 0xe8, 0xba, 0x22, 0xe4,
	0x05, 0x74, 0xad, 0x03, 0xfd, 0x4a, 0x45, 0x46, 0x93, 0x8c, 0x07, 0xab, 0xe8, 0xd4, 0x41, 0x7c,
	0x5a, 0x53, 0xf2, 0x14, 0x36, 0xac, 0x63, 0x92, 0xa9, 0xf4, 0x33, 0x67, 0xc1, 0x1a, 0xba, 0xb5,
	0x11, 0x9e, 0x59, 0x46, 0x0e, 0xe1, 0x21, 0x16, 0x8a, 0x0b, 0x2a, 0x58, 0x4c, 0x73, 0x35, 0x91,
	0x26, 0x58, 0x1f, 0x78, 0x07, 0x5e, 0xd4, 0x45, 0x61, 0x44, 0x05, 0x3b, 0x45, 0x1c, 0x1e, 0xc3,
	0xa6, 0x1d, 0xea, 0x95, 0x66, 0x5c, 0x97, 0xef, 0x44, 0x69, 0x94, 0x9e, 0xfe, 0x7f, 0xa4, 0x9f,
	0xa0, 0x57, 0x4f, 0x74, 0x3e, 0x6a, 0x1f, 0x5a, 0xb6, 0x70, 0x8a, 0x25, 0x6d, 0x1c, 0x20, 0x3a,
	0xaf, 0x08, 0x79, 0x0e, 0xab, 0x0a, 0x23, 0x02, 0x7f, 0xb0, 0x7c, 0xd0, 0x3a, 0xee, 0x0c, 0xeb,
	0x47, 0x83, 0x38, 0x72, 0x6a, 0xf8, 0xc3, 0x87, 0x06, 0x7e, 0x92, 0x0e, 0xf8, 0xae, 0x83, 0x66,
	0xe4, 0x0b, 0x46, 0x9e, 0x40, 0x3b, 0xd5, 0x9c, 0x1a, 0xce, 0x62, 0x46, 0x4d, 0xbd, 0x96, 0x96,
	0x63, 0x17, 0xd4, 0xf0, 0xca, 0xa5, 0xa0, 0xd3, 0xbc, 0x6a, 0xde, 0x4c, 0x0b, 0xee, 0x16, 0xd4,
	0x72, 0xec, 0x7a, 0x5a, 0x70, 0x72, 0x02, 0xdb, 0x4c, 0x94, 0xd8, 0x25, 0x67, 0xb1, 0x1b, 0x96,
	0x16, 0xa9, 0xdd, 0x9b, 0x17, 0x6d, 0xdd, 0xab, 0xb8, 0xa0, 0x51, 0xa5, 0x91, 0x67, 0xd0, 0xa9,
	0x13, 0xbb, 0xa1, 0x36, 0xd0, 0x7b, 0xc3, 0x51, 0x3b, 0x52, 0xf2, 0x12, 0x1e, 0xd0, 0xa2, 0xc8,
	0x44, 0xd5, 0xa2, 0x4b, 0x83, 0xdb, 0xf4, 0xa2, 0xae, 0xe3, 0x17, 0x0e, 0x93, 0xd7, 0xb0, 0x43,
	0xd3, 0x2f, 0x13, 0xa1, 0x39, 0x8b, 0x13, 0x9a, 0x51, 0x99, 0xf2, 0x38, 0xbd, 0xa3, 0x72, 0xcc,
	0x71, 0xb1, 0x5e, 0xd4, 0xab, 0xe5, 0x33, 0xab, 0x9e, 0xa3, 0x78, 0xfc, 0xd3, 0x83, 0xb5, 0xf7,
	0x76, 0x72, 0xe4, 0x0d, 0xac, 0x54, 0x27, 0x44, 0xb6, 0x87, 0xf6, 0xd0, 0x86, 0xf5, 0xa1, 0x0d,
	0xdf, 0x56, 0x87, 0xd6, 0xef, 0xcd, 0x66, 0xfc, 0xfb, 0xa5, 0x85, 0x4b, 0xe4, 0x06, 0xba, 0x8b,
	0x07, 0xb1, 0x77, 0xef, 0xfb, 0xb7, 0x4b, 0xeb, 0x0f, 0x66, 0xfa, 0x3f, 0x4e, 0x2a, 0x5c, 0x22,
	0x57, 0xb0, 0x31, 0xff, 0x2a, 0x1e, 0x2d, 0x24, 0x9d, 0x53, 0xfb, 0x7b, 0x7f, 0xa4, 0x9c, 0xd3,
	0xc3, 0xa5, 0xb3, 0xf6, 0x47, 0x18, 0xcf, 0xfe, 0x5d, 0x92, 0x55, 0xfc, 0x79, 0xaf, 0x7e, 0x05,
	0x00, 0x00, 0xff, 0xff, 0xcc, 0x36, 0xec, 0x5b, 0x75, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MindboxClient is the client API for Mindbox service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MindboxClient interface {
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PingResponse, error)
	UserInformation(ctx context.Context, in *ParamsUserInformation, opts ...grpc.CallOption) (*ResponseUserInformation, error)
	OrdersHistory(ctx context.Context, in *ParamsOrdersHistory, opts ...grpc.CallOption) (*ResponseOrdersHistory, error)
}

type mindboxClient struct {
	cc *grpc.ClientConn
}

func NewMindboxClient(cc *grpc.ClientConn) MindboxClient {
	return &mindboxClient{cc}
}

func (c *mindboxClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/mindbox.Mindbox/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mindboxClient) UserInformation(ctx context.Context, in *ParamsUserInformation, opts ...grpc.CallOption) (*ResponseUserInformation, error) {
	out := new(ResponseUserInformation)
	err := c.cc.Invoke(ctx, "/mindbox.Mindbox/UserInformation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mindboxClient) OrdersHistory(ctx context.Context, in *ParamsOrdersHistory, opts ...grpc.CallOption) (*ResponseOrdersHistory, error) {
	out := new(ResponseOrdersHistory)
	err := c.cc.Invoke(ctx, "/mindbox.Mindbox/OrdersHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MindboxServer is the server API for Mindbox service.
type MindboxServer interface {
	Ping(context.Context, *empty.Empty) (*PingResponse, error)
	UserInformation(context.Context, *ParamsUserInformation) (*ResponseUserInformation, error)
	OrdersHistory(context.Context, *ParamsOrdersHistory) (*ResponseOrdersHistory, error)
}

// UnimplementedMindboxServer can be embedded to have forward compatible implementations.
type UnimplementedMindboxServer struct {
}

func (*UnimplementedMindboxServer) Ping(ctx context.Context, req *empty.Empty) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedMindboxServer) UserInformation(ctx context.Context, req *ParamsUserInformation) (*ResponseUserInformation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInformation not implemented")
}
func (*UnimplementedMindboxServer) OrdersHistory(ctx context.Context, req *ParamsOrdersHistory) (*ResponseOrdersHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrdersHistory not implemented")
}

func RegisterMindboxServer(s *grpc.Server, srv MindboxServer) {
	s.RegisterService(&_Mindbox_serviceDesc, srv)
}

func _Mindbox_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MindboxServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindbox.Mindbox/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MindboxServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mindbox_UserInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsUserInformation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MindboxServer).UserInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindbox.Mindbox/UserInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MindboxServer).UserInformation(ctx, req.(*ParamsUserInformation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mindbox_OrdersHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsOrdersHistory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MindboxServer).OrdersHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindbox.Mindbox/OrdersHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MindboxServer).OrdersHistory(ctx, req.(*ParamsOrdersHistory))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mindbox_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mindbox.Mindbox",
	HandlerType: (*MindboxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Mindbox_Ping_Handler,
		},
		{
			MethodName: "UserInformation",
			Handler:    _Mindbox_UserInformation_Handler,
		},
		{
			MethodName: "OrdersHistory",
			Handler:    _Mindbox_OrdersHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/mindbox.proto",
}