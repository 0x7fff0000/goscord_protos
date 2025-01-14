// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: sso/v1/sso.proto

package ssoconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/0x7fff0000/goscord-protos/gen/go/sso/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// AuthServiceName is the fully-qualified name of the AuthService service.
	AuthServiceName = "sso.v1.AuthService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AuthServiceRegisterProcedure is the fully-qualified name of the AuthService's Register RPC.
	AuthServiceRegisterProcedure = "/sso.v1.AuthService/Register"
	// AuthServiceLoginProcedure is the fully-qualified name of the AuthService's Login RPC.
	AuthServiceLoginProcedure = "/sso.v1.AuthService/Login"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	authServiceServiceDescriptor        = v1.File_sso_v1_sso_proto.Services().ByName("AuthService")
	authServiceRegisterMethodDescriptor = authServiceServiceDescriptor.Methods().ByName("Register")
	authServiceLoginMethodDescriptor    = authServiceServiceDescriptor.Methods().ByName("Login")
)

// AuthServiceClient is a client for the sso.v1.AuthService service.
type AuthServiceClient interface {
	Register(context.Context, *connect.Request[v1.RegisterRequest]) (*connect.Response[v1.RegisterResponse], error)
	Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error)
}

// NewAuthServiceClient constructs a client for the sso.v1.AuthService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AuthServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &authServiceClient{
		register: connect.NewClient[v1.RegisterRequest, v1.RegisterResponse](
			httpClient,
			baseURL+AuthServiceRegisterProcedure,
			connect.WithSchema(authServiceRegisterMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		login: connect.NewClient[v1.LoginRequest, v1.LoginResponse](
			httpClient,
			baseURL+AuthServiceLoginProcedure,
			connect.WithSchema(authServiceLoginMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// authServiceClient implements AuthServiceClient.
type authServiceClient struct {
	register *connect.Client[v1.RegisterRequest, v1.RegisterResponse]
	login    *connect.Client[v1.LoginRequest, v1.LoginResponse]
}

// Register calls sso.v1.AuthService.Register.
func (c *authServiceClient) Register(ctx context.Context, req *connect.Request[v1.RegisterRequest]) (*connect.Response[v1.RegisterResponse], error) {
	return c.register.CallUnary(ctx, req)
}

// Login calls sso.v1.AuthService.Login.
func (c *authServiceClient) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	return c.login.CallUnary(ctx, req)
}

// AuthServiceHandler is an implementation of the sso.v1.AuthService service.
type AuthServiceHandler interface {
	Register(context.Context, *connect.Request[v1.RegisterRequest]) (*connect.Response[v1.RegisterResponse], error)
	Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error)
}

// NewAuthServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthServiceHandler(svc AuthServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	authServiceRegisterHandler := connect.NewUnaryHandler(
		AuthServiceRegisterProcedure,
		svc.Register,
		connect.WithSchema(authServiceRegisterMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	authServiceLoginHandler := connect.NewUnaryHandler(
		AuthServiceLoginProcedure,
		svc.Login,
		connect.WithSchema(authServiceLoginMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/sso.v1.AuthService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AuthServiceRegisterProcedure:
			authServiceRegisterHandler.ServeHTTP(w, r)
		case AuthServiceLoginProcedure:
			authServiceLoginHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAuthServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthServiceHandler struct{}

func (UnimplementedAuthServiceHandler) Register(context.Context, *connect.Request[v1.RegisterRequest]) (*connect.Response[v1.RegisterResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("sso.v1.AuthService.Register is not implemented"))
}

func (UnimplementedAuthServiceHandler) Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("sso.v1.AuthService.Login is not implemented"))
}
