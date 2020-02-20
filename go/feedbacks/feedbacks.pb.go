// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/feedbacks.proto

package feedbacks

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ResponseOk struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseOk) Reset()         { *m = ResponseOk{} }
func (m *ResponseOk) String() string { return proto.CompactTextString(m) }
func (*ResponseOk) ProtoMessage()    {}
func (*ResponseOk) Descriptor() ([]byte, []int) {
	return fileDescriptor_ada70ba69b487c57, []int{0}
}

func (m *ResponseOk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseOk.Unmarshal(m, b)
}
func (m *ResponseOk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseOk.Marshal(b, m, deterministic)
}
func (m *ResponseOk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseOk.Merge(m, src)
}
func (m *ResponseOk) XXX_Size() int {
	return xxx_messageInfo_ResponseOk.Size(m)
}
func (m *ResponseOk) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseOk.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseOk proto.InternalMessageInfo

func (m *ResponseOk) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type RequestApp struct {
	ClientId             int64    `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	AppVersion           string   `protobuf:"bytes,2,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	Rate                 string   `protobuf:"bytes,3,opt,name=rate,proto3" json:"rate,omitempty"`
	Reason               []int32  `protobuf:"varint,4,rep,packed,name=reason,proto3" json:"reason,omitempty"`
	Comment              string   `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
	Date                 string   `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestApp) Reset()         { *m = RequestApp{} }
func (m *RequestApp) String() string { return proto.CompactTextString(m) }
func (*RequestApp) ProtoMessage()    {}
func (*RequestApp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ada70ba69b487c57, []int{1}
}

func (m *RequestApp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestApp.Unmarshal(m, b)
}
func (m *RequestApp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestApp.Marshal(b, m, deterministic)
}
func (m *RequestApp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestApp.Merge(m, src)
}
func (m *RequestApp) XXX_Size() int {
	return xxx_messageInfo_RequestApp.Size(m)
}
func (m *RequestApp) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestApp.DiscardUnknown(m)
}

var xxx_messageInfo_RequestApp proto.InternalMessageInfo

func (m *RequestApp) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *RequestApp) GetAppVersion() string {
	if m != nil {
		return m.AppVersion
	}
	return ""
}

func (m *RequestApp) GetRate() string {
	if m != nil {
		return m.Rate
	}
	return ""
}

func (m *RequestApp) GetReason() []int32 {
	if m != nil {
		return m.Reason
	}
	return nil
}

func (m *RequestApp) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *RequestApp) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

type RequestStore struct {
	ClientId             int64    `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Rate                 string   `protobuf:"bytes,2,opt,name=rate,proto3" json:"rate,omitempty"`
	Reason               []int32  `protobuf:"varint,3,rep,packed,name=reason,proto3" json:"reason,omitempty"`
	Comment              string   `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
	City                 string   `protobuf:"bytes,5,opt,name=city,proto3" json:"city,omitempty"`
	Date                 string   `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestStore) Reset()         { *m = RequestStore{} }
func (m *RequestStore) String() string { return proto.CompactTextString(m) }
func (*RequestStore) ProtoMessage()    {}
func (*RequestStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_ada70ba69b487c57, []int{2}
}

func (m *RequestStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestStore.Unmarshal(m, b)
}
func (m *RequestStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestStore.Marshal(b, m, deterministic)
}
func (m *RequestStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestStore.Merge(m, src)
}
func (m *RequestStore) XXX_Size() int {
	return xxx_messageInfo_RequestStore.Size(m)
}
func (m *RequestStore) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestStore.DiscardUnknown(m)
}

var xxx_messageInfo_RequestStore proto.InternalMessageInfo

func (m *RequestStore) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *RequestStore) GetRate() string {
	if m != nil {
		return m.Rate
	}
	return ""
}

func (m *RequestStore) GetReason() []int32 {
	if m != nil {
		return m.Reason
	}
	return nil
}

func (m *RequestStore) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *RequestStore) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *RequestStore) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

type RequestOrder struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	ClientId             int64    `protobuf:"varint,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	PaymentType          string   `protobuf:"bytes,3,opt,name=payment_type,json=paymentType,proto3" json:"payment_type,omitempty"`
	DeliveryType         string   `protobuf:"bytes,4,opt,name=delivery_type,json=deliveryType,proto3" json:"delivery_type,omitempty"`
	Rate                 string   `protobuf:"bytes,5,opt,name=rate,proto3" json:"rate,omitempty"`
	Reason               []int32  `protobuf:"varint,6,rep,packed,name=reason,proto3" json:"reason,omitempty"`
	Comment              string   `protobuf:"bytes,7,opt,name=comment,proto3" json:"comment,omitempty"`
	DateOrder            string   `protobuf:"bytes,8,opt,name=date_order,json=dateOrder,proto3" json:"date_order,omitempty"`
	DateRate             string   `protobuf:"bytes,9,opt,name=date_rate,json=dateRate,proto3" json:"date_rate,omitempty"`
	StoreName            string   `protobuf:"bytes,10,opt,name=store_name,json=storeName,proto3" json:"store_name,omitempty"`
	TransportCompany     string   `protobuf:"bytes,11,opt,name=transport_company,json=transportCompany,proto3" json:"transport_company,omitempty"`
	Country              string   `protobuf:"bytes,12,opt,name=country,proto3" json:"country,omitempty"`
	City                 string   `protobuf:"bytes,13,opt,name=city,proto3" json:"city,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestOrder) Reset()         { *m = RequestOrder{} }
func (m *RequestOrder) String() string { return proto.CompactTextString(m) }
func (*RequestOrder) ProtoMessage()    {}
func (*RequestOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_ada70ba69b487c57, []int{3}
}

func (m *RequestOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestOrder.Unmarshal(m, b)
}
func (m *RequestOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestOrder.Marshal(b, m, deterministic)
}
func (m *RequestOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestOrder.Merge(m, src)
}
func (m *RequestOrder) XXX_Size() int {
	return xxx_messageInfo_RequestOrder.Size(m)
}
func (m *RequestOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestOrder.DiscardUnknown(m)
}

var xxx_messageInfo_RequestOrder proto.InternalMessageInfo

func (m *RequestOrder) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *RequestOrder) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *RequestOrder) GetPaymentType() string {
	if m != nil {
		return m.PaymentType
	}
	return ""
}

func (m *RequestOrder) GetDeliveryType() string {
	if m != nil {
		return m.DeliveryType
	}
	return ""
}

func (m *RequestOrder) GetRate() string {
	if m != nil {
		return m.Rate
	}
	return ""
}

func (m *RequestOrder) GetReason() []int32 {
	if m != nil {
		return m.Reason
	}
	return nil
}

func (m *RequestOrder) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *RequestOrder) GetDateOrder() string {
	if m != nil {
		return m.DateOrder
	}
	return ""
}

func (m *RequestOrder) GetDateRate() string {
	if m != nil {
		return m.DateRate
	}
	return ""
}

func (m *RequestOrder) GetStoreName() string {
	if m != nil {
		return m.StoreName
	}
	return ""
}

func (m *RequestOrder) GetTransportCompany() string {
	if m != nil {
		return m.TransportCompany
	}
	return ""
}

func (m *RequestOrder) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *RequestOrder) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func init() {
	proto.RegisterType((*ResponseOk)(nil), "feedbacks.ResponseOk")
	proto.RegisterType((*RequestApp)(nil), "feedbacks.RequestApp")
	proto.RegisterType((*RequestStore)(nil), "feedbacks.RequestStore")
	proto.RegisterType((*RequestOrder)(nil), "feedbacks.RequestOrder")
}

func init() { proto.RegisterFile("proto/feedbacks.proto", fileDescriptor_ada70ba69b487c57) }

var fileDescriptor_ada70ba69b487c57 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x49, 0xd3, 0xa6, 0xc9, 0x34, 0x8b, 0xc0, 0xd2, 0x82, 0x61, 0xb5, 0xa2, 0x14, 0x09,
	0x55, 0x42, 0x5a, 0x24, 0x10, 0x07, 0x8e, 0xc0, 0x69, 0x0f, 0xb0, 0x52, 0x40, 0x1c, 0xb8, 0x44,
	0x6e, 0x32, 0xa0, 0x68, 0x9b, 0xd8, 0xd8, 0xee, 0x4a, 0x79, 0x05, 0xde, 0x80, 0x3b, 0xaf, 0xc0,
	0xfb, 0x21, 0x8f, 0xeb, 0x96, 0xee, 0xb6, 0x7b, 0xf3, 0xff, 0xcf, 0x8c, 0xfd, 0x8d, 0xed, 0x81,
	0x63, 0xa5, 0xa5, 0x95, 0x2f, 0xbf, 0x23, 0xd6, 0x0b, 0x51, 0x5d, 0x9a, 0x33, 0xd2, 0x2c, 0xdb,
	0x18, 0xb3, 0xe7, 0x00, 0x05, 0x1a, 0x25, 0x3b, 0x83, 0x17, 0x97, 0x8c, 0xc3, 0xd8, 0xac, 0xaa,
	0x0a, 0x8d, 0xe1, 0xd1, 0x34, 0x9a, 0xa7, 0x45, 0x90, 0xb3, 0x3f, 0x91, 0x4b, 0xfc, 0xb9, 0x42,
	0x63, 0xdf, 0x29, 0xc5, 0x4e, 0x20, 0xab, 0x96, 0x0d, 0x76, 0xb6, 0x6c, 0x6a, 0x4a, 0x8d, 0x8b,
	0xd4, 0x1b, 0xe7, 0x35, 0x7b, 0x02, 0x13, 0xa1, 0x54, 0x79, 0x85, 0xda, 0x34, 0xb2, 0xe3, 0x83,
	0x69, 0x34, 0xcf, 0x0a, 0x10, 0x4a, 0x7d, 0xf5, 0x0e, 0x63, 0x30, 0xd4, 0xc2, 0x22, 0x8f, 0x29,
	0x42, 0x6b, 0xf6, 0x00, 0x12, 0x8d, 0xc2, 0xc8, 0x8e, 0x0f, 0xa7, 0xf1, 0x7c, 0x54, 0xac, 0x95,
	0x43, 0xaa, 0x64, 0xdb, 0x62, 0x67, 0xf9, 0x88, 0xd2, 0x83, 0x74, 0xbb, 0xd4, 0x6e, 0x97, 0xc4,
	0xef, 0xe2, 0xd6, 0xb3, 0xdf, 0x11, 0xe4, 0x6b, 0xcc, 0xcf, 0x56, 0x6a, 0xbc, 0x1d, 0x34, 0x70,
	0x0c, 0xf6, 0x72, 0xc4, 0x87, 0x38, 0x86, 0x37, 0x38, 0xaa, 0xc6, 0xf6, 0x6b, 0x3c, 0x5a, 0xef,
	0x65, 0xfb, 0x15, 0x6f, 0xd8, 0x2e, 0x74, 0x8d, 0x9a, 0x3d, 0x82, 0x54, 0xba, 0x45, 0x40, 0xcb,
	0x8a, 0x31, 0xe9, 0xf3, 0x7a, 0x17, 0x7b, 0x70, 0x0d, 0xfb, 0x29, 0xe4, 0x4a, 0xf4, 0xee, 0xec,
	0xd2, 0xf6, 0x2a, 0x5c, 0xe3, 0x64, 0xed, 0x7d, 0xe9, 0x15, 0xb2, 0x67, 0x70, 0x54, 0xe3, 0xb2,
	0xb9, 0x42, 0xdd, 0xfb, 0x1c, 0xcf, 0x9c, 0x07, 0x93, 0x92, 0x42, 0xfb, 0xa3, 0xbd, 0xed, 0x27,
	0x87, 0xda, 0x1f, 0xef, 0xb6, 0x7f, 0x0a, 0xe0, 0xda, 0x2b, 0x09, 0x9d, 0xa7, 0x14, 0xcc, 0x9c,
	0xe3, 0x9b, 0x3c, 0x01, 0x12, 0x25, 0x9d, 0x94, 0x51, 0x34, 0x75, 0x46, 0xe1, 0x4e, 0x3b, 0x05,
	0x30, 0xee, 0x99, 0xca, 0x4e, 0xb4, 0xc8, 0xc1, 0xd7, 0x92, 0xf3, 0x49, 0xb4, 0xc8, 0x5e, 0xc0,
	0x7d, 0xab, 0x45, 0x67, 0x94, 0xd4, 0xb6, 0xac, 0x64, 0xab, 0x44, 0xd7, 0xf3, 0x09, 0x65, 0xdd,
	0xdb, 0x04, 0x3e, 0x78, 0xdf, 0x13, 0xae, 0x3a, 0xab, 0x7b, 0x9e, 0x07, 0x42, 0x92, 0x9b, 0x07,
	0x3a, 0xda, 0x3e, 0xd0, 0xab, 0xbf, 0x11, 0x24, 0x1f, 0xe5, 0xa2, 0x59, 0x22, 0x7b, 0x03, 0xb1,
	0xfb, 0xd2, 0xc7, 0x67, 0xdb, 0x31, 0xd9, 0xfe, 0xf4, 0xc7, 0xbb, 0x76, 0x98, 0x94, 0xd9, 0x1d,
	0xf6, 0x16, 0x46, 0xfe, 0x8b, 0x3d, 0xbc, 0x59, 0x48, 0x81, 0x5b, 0x4b, 0xfd, 0xe5, 0xec, 0x29,
	0xa5, 0xc0, 0xc1, 0xd2, 0xf7, 0x77, 0xbf, 0xe5, 0x3f, 0xfe, 0x1b, 0xe8, 0x45, 0x42, 0x13, 0xfd,
	0xfa, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0xa9, 0xfe, 0x12, 0xea, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MobileClient is the client API for Mobile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MobileClient interface {
	App(ctx context.Context, in *RequestApp, opts ...grpc.CallOption) (*ResponseOk, error)
	Store(ctx context.Context, in *RequestStore, opts ...grpc.CallOption) (*ResponseOk, error)
	Order(ctx context.Context, in *RequestOrder, opts ...grpc.CallOption) (*ResponseOk, error)
}

type mobileClient struct {
	cc *grpc.ClientConn
}

func NewMobileClient(cc *grpc.ClientConn) MobileClient {
	return &mobileClient{cc}
}

func (c *mobileClient) App(ctx context.Context, in *RequestApp, opts ...grpc.CallOption) (*ResponseOk, error) {
	out := new(ResponseOk)
	err := c.cc.Invoke(ctx, "/feedbacks.Mobile/App", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) Store(ctx context.Context, in *RequestStore, opts ...grpc.CallOption) (*ResponseOk, error) {
	out := new(ResponseOk)
	err := c.cc.Invoke(ctx, "/feedbacks.Mobile/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) Order(ctx context.Context, in *RequestOrder, opts ...grpc.CallOption) (*ResponseOk, error) {
	out := new(ResponseOk)
	err := c.cc.Invoke(ctx, "/feedbacks.Mobile/Order", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MobileServer is the server API for Mobile service.
type MobileServer interface {
	App(context.Context, *RequestApp) (*ResponseOk, error)
	Store(context.Context, *RequestStore) (*ResponseOk, error)
	Order(context.Context, *RequestOrder) (*ResponseOk, error)
}

// UnimplementedMobileServer can be embedded to have forward compatible implementations.
type UnimplementedMobileServer struct {
}

func (*UnimplementedMobileServer) App(ctx context.Context, req *RequestApp) (*ResponseOk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method App not implemented")
}
func (*UnimplementedMobileServer) Store(ctx context.Context, req *RequestStore) (*ResponseOk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (*UnimplementedMobileServer) Order(ctx context.Context, req *RequestOrder) (*ResponseOk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Order not implemented")
}

func RegisterMobileServer(s *grpc.Server, srv MobileServer) {
	s.RegisterService(&_Mobile_serviceDesc, srv)
}

func _Mobile_App_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestApp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServer).App(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedbacks.Mobile/App",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServer).App(ctx, req.(*RequestApp))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mobile_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestStore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedbacks.Mobile/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServer).Store(ctx, req.(*RequestStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mobile_Order_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServer).Order(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedbacks.Mobile/Order",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServer).Order(ctx, req.(*RequestOrder))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mobile_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feedbacks.Mobile",
	HandlerType: (*MobileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "App",
			Handler:    _Mobile_App_Handler,
		},
		{
			MethodName: "Store",
			Handler:    _Mobile_Store_Handler,
		},
		{
			MethodName: "Order",
			Handler:    _Mobile_Order_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/feedbacks.proto",
}