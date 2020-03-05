// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/meta.proto

package meta

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Mobile service

type MobileService interface {
	Contacts(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ResponseMobileAPIContacts, error)
	About(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ResponseMobileApiAbout, error)
	Faq(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ResponseFaq, error)
	Countries(ctx context.Context, in *ParamsCountries, opts ...client.CallOption) (*ResponseCountries, error)
}

type mobileService struct {
	c    client.Client
	name string
}

func NewMobileService(name string, c client.Client) MobileService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "meta"
	}
	return &mobileService{
		c:    c,
		name: name,
	}
}

func (c *mobileService) Contacts(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ResponseMobileAPIContacts, error) {
	req := c.c.NewRequest(c.name, "Mobile.Contacts", in)
	out := new(ResponseMobileAPIContacts)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileService) About(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ResponseMobileApiAbout, error) {
	req := c.c.NewRequest(c.name, "Mobile.About", in)
	out := new(ResponseMobileApiAbout)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileService) Faq(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ResponseFaq, error) {
	req := c.c.NewRequest(c.name, "Mobile.Faq", in)
	out := new(ResponseFaq)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileService) Countries(ctx context.Context, in *ParamsCountries, opts ...client.CallOption) (*ResponseCountries, error) {
	req := c.c.NewRequest(c.name, "Mobile.Countries", in)
	out := new(ResponseCountries)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Mobile service

type MobileHandler interface {
	Contacts(context.Context, *empty.Empty, *ResponseMobileAPIContacts) error
	About(context.Context, *empty.Empty, *ResponseMobileApiAbout) error
	Faq(context.Context, *empty.Empty, *ResponseFaq) error
	Countries(context.Context, *ParamsCountries, *ResponseCountries) error
}

func RegisterMobileHandler(s server.Server, hdlr MobileHandler, opts ...server.HandlerOption) error {
	type mobile interface {
		Contacts(ctx context.Context, in *empty.Empty, out *ResponseMobileAPIContacts) error
		About(ctx context.Context, in *empty.Empty, out *ResponseMobileApiAbout) error
		Faq(ctx context.Context, in *empty.Empty, out *ResponseFaq) error
		Countries(ctx context.Context, in *ParamsCountries, out *ResponseCountries) error
	}
	type Mobile struct {
		mobile
	}
	h := &mobileHandler{hdlr}
	return s.Handle(s.NewHandler(&Mobile{h}, opts...))
}

type mobileHandler struct {
	MobileHandler
}

func (h *mobileHandler) Contacts(ctx context.Context, in *empty.Empty, out *ResponseMobileAPIContacts) error {
	return h.MobileHandler.Contacts(ctx, in, out)
}

func (h *mobileHandler) About(ctx context.Context, in *empty.Empty, out *ResponseMobileApiAbout) error {
	return h.MobileHandler.About(ctx, in, out)
}

func (h *mobileHandler) Faq(ctx context.Context, in *empty.Empty, out *ResponseFaq) error {
	return h.MobileHandler.Faq(ctx, in, out)
}

func (h *mobileHandler) Countries(ctx context.Context, in *ParamsCountries, out *ResponseCountries) error {
	return h.MobileHandler.Countries(ctx, in, out)
}

// Client API for Stores service

type StoresService interface {
	All(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ResponseAllOfflineStoresInfo, error)
	ByID(ctx context.Context, in *ParamsOfflineStoreInfoByID, opts ...client.CallOption) (*ResponseOfflineStoreInfoByID, error)
	Cities(ctx context.Context, in *ParamsStoresCities, opts ...client.CallOption) (*ResponseStoresCities, error)
	ByCity(ctx context.Context, in *ParamsStoresByCity, opts ...client.CallOption) (*ResponseStoresByCity, error)
}

type storesService struct {
	c    client.Client
	name string
}

func NewStoresService(name string, c client.Client) StoresService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "meta"
	}
	return &storesService{
		c:    c,
		name: name,
	}
}

func (c *storesService) All(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ResponseAllOfflineStoresInfo, error) {
	req := c.c.NewRequest(c.name, "Stores.All", in)
	out := new(ResponseAllOfflineStoresInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storesService) ByID(ctx context.Context, in *ParamsOfflineStoreInfoByID, opts ...client.CallOption) (*ResponseOfflineStoreInfoByID, error) {
	req := c.c.NewRequest(c.name, "Stores.ByID", in)
	out := new(ResponseOfflineStoreInfoByID)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storesService) Cities(ctx context.Context, in *ParamsStoresCities, opts ...client.CallOption) (*ResponseStoresCities, error) {
	req := c.c.NewRequest(c.name, "Stores.Cities", in)
	out := new(ResponseStoresCities)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storesService) ByCity(ctx context.Context, in *ParamsStoresByCity, opts ...client.CallOption) (*ResponseStoresByCity, error) {
	req := c.c.NewRequest(c.name, "Stores.ByCity", in)
	out := new(ResponseStoresByCity)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Stores service

type StoresHandler interface {
	All(context.Context, *empty.Empty, *ResponseAllOfflineStoresInfo) error
	ByID(context.Context, *ParamsOfflineStoreInfoByID, *ResponseOfflineStoreInfoByID) error
	Cities(context.Context, *ParamsStoresCities, *ResponseStoresCities) error
	ByCity(context.Context, *ParamsStoresByCity, *ResponseStoresByCity) error
}

func RegisterStoresHandler(s server.Server, hdlr StoresHandler, opts ...server.HandlerOption) error {
	type stores interface {
		All(ctx context.Context, in *empty.Empty, out *ResponseAllOfflineStoresInfo) error
		ByID(ctx context.Context, in *ParamsOfflineStoreInfoByID, out *ResponseOfflineStoreInfoByID) error
		Cities(ctx context.Context, in *ParamsStoresCities, out *ResponseStoresCities) error
		ByCity(ctx context.Context, in *ParamsStoresByCity, out *ResponseStoresByCity) error
	}
	type Stores struct {
		stores
	}
	h := &storesHandler{hdlr}
	return s.Handle(s.NewHandler(&Stores{h}, opts...))
}

type storesHandler struct {
	StoresHandler
}

func (h *storesHandler) All(ctx context.Context, in *empty.Empty, out *ResponseAllOfflineStoresInfo) error {
	return h.StoresHandler.All(ctx, in, out)
}

func (h *storesHandler) ByID(ctx context.Context, in *ParamsOfflineStoreInfoByID, out *ResponseOfflineStoreInfoByID) error {
	return h.StoresHandler.ByID(ctx, in, out)
}

func (h *storesHandler) Cities(ctx context.Context, in *ParamsStoresCities, out *ResponseStoresCities) error {
	return h.StoresHandler.Cities(ctx, in, out)
}

func (h *storesHandler) ByCity(ctx context.Context, in *ParamsStoresByCity, out *ResponseStoresByCity) error {
	return h.StoresHandler.ByCity(ctx, in, out)
}
