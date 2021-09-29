// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: service.proto

package accounts

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	pipes "github.com/nochte/pipelinr-protocol/protobuf/pipes"
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

// Api Endpoints for Accounts service

func NewAccountsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Accounts service

type AccountsService interface {
	// These are unauthenticated, wide-open
	Signup(ctx context.Context, in *SignupRequest, opts ...client.CallOption) (*pipes.GenericResponse, error)
	ConfirmSignup(ctx context.Context, in *ConfirmSignupRequest, opts ...client.CallOption) (*LoginResponse, error)
	Login(ctx context.Context, in *UsernameLoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	ForgotPassword(ctx context.Context, in *UsernameLoginRequest, opts ...client.CallOption) (*pipes.GenericResponse, error)
	ConfirmForgotPassword(ctx context.Context, in *ConfirmSignupRequest, opts ...client.CallOption) (*LoginResponse, error)
	ValidateJWT(ctx context.Context, in *JWTValidationRequest, opts ...client.CallOption) (*LoginResponse, error)
	CreateAPIKey(ctx context.Context, in *CreateAPIKeyRequest, opts ...client.CallOption) (*CreateAPIKeyResponse, error)
	// NOTE: ListAPIKeys will not return the actual Key value
	ListAPIKeys(ctx context.Context, in *pipes.Null, opts ...client.CallOption) (*ListAPIKeyResponse, error)
	DeleteAPIKey(ctx context.Context, in *pipes.Xid, opts ...client.CallOption) (*pipes.GenericResponse, error)
	ValidateAPIKey(ctx context.Context, in *ValidateAPIKeyRequestResponse, opts ...client.CallOption) (*ValidateAPIKeyRequestResponse, error)
}

type accountsService struct {
	c    client.Client
	name string
}

func NewAccountsService(name string, c client.Client) AccountsService {
	return &accountsService{
		c:    c,
		name: name,
	}
}

func (c *accountsService) Signup(ctx context.Context, in *SignupRequest, opts ...client.CallOption) (*pipes.GenericResponse, error) {
	req := c.c.NewRequest(c.name, "Accounts.Signup", in)
	out := new(pipes.GenericResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsService) ConfirmSignup(ctx context.Context, in *ConfirmSignupRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "Accounts.ConfirmSignup", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsService) Login(ctx context.Context, in *UsernameLoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "Accounts.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsService) ForgotPassword(ctx context.Context, in *UsernameLoginRequest, opts ...client.CallOption) (*pipes.GenericResponse, error) {
	req := c.c.NewRequest(c.name, "Accounts.ForgotPassword", in)
	out := new(pipes.GenericResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsService) ConfirmForgotPassword(ctx context.Context, in *ConfirmSignupRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "Accounts.ConfirmForgotPassword", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsService) ValidateJWT(ctx context.Context, in *JWTValidationRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "Accounts.ValidateJWT", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsService) CreateAPIKey(ctx context.Context, in *CreateAPIKeyRequest, opts ...client.CallOption) (*CreateAPIKeyResponse, error) {
	req := c.c.NewRequest(c.name, "Accounts.CreateAPIKey", in)
	out := new(CreateAPIKeyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsService) ListAPIKeys(ctx context.Context, in *pipes.Null, opts ...client.CallOption) (*ListAPIKeyResponse, error) {
	req := c.c.NewRequest(c.name, "Accounts.ListAPIKeys", in)
	out := new(ListAPIKeyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsService) DeleteAPIKey(ctx context.Context, in *pipes.Xid, opts ...client.CallOption) (*pipes.GenericResponse, error) {
	req := c.c.NewRequest(c.name, "Accounts.DeleteAPIKey", in)
	out := new(pipes.GenericResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsService) ValidateAPIKey(ctx context.Context, in *ValidateAPIKeyRequestResponse, opts ...client.CallOption) (*ValidateAPIKeyRequestResponse, error) {
	req := c.c.NewRequest(c.name, "Accounts.ValidateAPIKey", in)
	out := new(ValidateAPIKeyRequestResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Accounts service

type AccountsHandler interface {
	// These are unauthenticated, wide-open
	Signup(context.Context, *SignupRequest, *pipes.GenericResponse) error
	ConfirmSignup(context.Context, *ConfirmSignupRequest, *LoginResponse) error
	Login(context.Context, *UsernameLoginRequest, *LoginResponse) error
	ForgotPassword(context.Context, *UsernameLoginRequest, *pipes.GenericResponse) error
	ConfirmForgotPassword(context.Context, *ConfirmSignupRequest, *LoginResponse) error
	ValidateJWT(context.Context, *JWTValidationRequest, *LoginResponse) error
	CreateAPIKey(context.Context, *CreateAPIKeyRequest, *CreateAPIKeyResponse) error
	// NOTE: ListAPIKeys will not return the actual Key value
	ListAPIKeys(context.Context, *pipes.Null, *ListAPIKeyResponse) error
	DeleteAPIKey(context.Context, *pipes.Xid, *pipes.GenericResponse) error
	ValidateAPIKey(context.Context, *ValidateAPIKeyRequestResponse, *ValidateAPIKeyRequestResponse) error
}

func RegisterAccountsHandler(s server.Server, hdlr AccountsHandler, opts ...server.HandlerOption) error {
	type accounts interface {
		Signup(ctx context.Context, in *SignupRequest, out *pipes.GenericResponse) error
		ConfirmSignup(ctx context.Context, in *ConfirmSignupRequest, out *LoginResponse) error
		Login(ctx context.Context, in *UsernameLoginRequest, out *LoginResponse) error
		ForgotPassword(ctx context.Context, in *UsernameLoginRequest, out *pipes.GenericResponse) error
		ConfirmForgotPassword(ctx context.Context, in *ConfirmSignupRequest, out *LoginResponse) error
		ValidateJWT(ctx context.Context, in *JWTValidationRequest, out *LoginResponse) error
		CreateAPIKey(ctx context.Context, in *CreateAPIKeyRequest, out *CreateAPIKeyResponse) error
		ListAPIKeys(ctx context.Context, in *pipes.Null, out *ListAPIKeyResponse) error
		DeleteAPIKey(ctx context.Context, in *pipes.Xid, out *pipes.GenericResponse) error
		ValidateAPIKey(ctx context.Context, in *ValidateAPIKeyRequestResponse, out *ValidateAPIKeyRequestResponse) error
	}
	type Accounts struct {
		accounts
	}
	h := &accountsHandler{hdlr}
	return s.Handle(s.NewHandler(&Accounts{h}, opts...))
}

type accountsHandler struct {
	AccountsHandler
}

func (h *accountsHandler) Signup(ctx context.Context, in *SignupRequest, out *pipes.GenericResponse) error {
	return h.AccountsHandler.Signup(ctx, in, out)
}

func (h *accountsHandler) ConfirmSignup(ctx context.Context, in *ConfirmSignupRequest, out *LoginResponse) error {
	return h.AccountsHandler.ConfirmSignup(ctx, in, out)
}

func (h *accountsHandler) Login(ctx context.Context, in *UsernameLoginRequest, out *LoginResponse) error {
	return h.AccountsHandler.Login(ctx, in, out)
}

func (h *accountsHandler) ForgotPassword(ctx context.Context, in *UsernameLoginRequest, out *pipes.GenericResponse) error {
	return h.AccountsHandler.ForgotPassword(ctx, in, out)
}

func (h *accountsHandler) ConfirmForgotPassword(ctx context.Context, in *ConfirmSignupRequest, out *LoginResponse) error {
	return h.AccountsHandler.ConfirmForgotPassword(ctx, in, out)
}

func (h *accountsHandler) ValidateJWT(ctx context.Context, in *JWTValidationRequest, out *LoginResponse) error {
	return h.AccountsHandler.ValidateJWT(ctx, in, out)
}

func (h *accountsHandler) CreateAPIKey(ctx context.Context, in *CreateAPIKeyRequest, out *CreateAPIKeyResponse) error {
	return h.AccountsHandler.CreateAPIKey(ctx, in, out)
}

func (h *accountsHandler) ListAPIKeys(ctx context.Context, in *pipes.Null, out *ListAPIKeyResponse) error {
	return h.AccountsHandler.ListAPIKeys(ctx, in, out)
}

func (h *accountsHandler) DeleteAPIKey(ctx context.Context, in *pipes.Xid, out *pipes.GenericResponse) error {
	return h.AccountsHandler.DeleteAPIKey(ctx, in, out)
}

func (h *accountsHandler) ValidateAPIKey(ctx context.Context, in *ValidateAPIKeyRequestResponse, out *ValidateAPIKeyRequestResponse) error {
	return h.AccountsHandler.ValidateAPIKey(ctx, in, out)
}