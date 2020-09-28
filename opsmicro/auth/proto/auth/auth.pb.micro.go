// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: auth.proto

package com_ops_auth_service_auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Auth service

func NewAuthEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Auth service

type AuthService interface {
	Login(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	ValidateAccessToken(ctx context.Context, in *ValidateTokenRequest, opts ...client.CallOption) (*ValidateTokenResponse, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) Login(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.Login", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ValidateAccessToken(ctx context.Context, in *ValidateTokenRequest, opts ...client.CallOption) (*ValidateTokenResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.ValidateAccessToken", in)
	out := new(ValidateTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	Login(context.Context, *Request, *Response) error
	ValidateAccessToken(context.Context, *ValidateTokenRequest, *ValidateTokenResponse) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type auth interface {
		Login(ctx context.Context, in *Request, out *Response) error
		ValidateAccessToken(ctx context.Context, in *ValidateTokenRequest, out *ValidateTokenResponse) error
	}
	type Auth struct {
		auth
	}
	h := &authHandler{hdlr}
	return s.Handle(s.NewHandler(&Auth{h}, opts...))
}

type authHandler struct {
	AuthHandler
}

func (h *authHandler) Login(ctx context.Context, in *Request, out *Response) error {
	return h.AuthHandler.Login(ctx, in, out)
}

func (h *authHandler) ValidateAccessToken(ctx context.Context, in *ValidateTokenRequest, out *ValidateTokenResponse) error {
	return h.AuthHandler.ValidateAccessToken(ctx, in, out)
}
