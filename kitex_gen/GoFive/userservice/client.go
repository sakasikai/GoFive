// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	gofive "github.com/sakasikai/GoFive/kitex_gen/GoFive"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateUser(ctx context.Context, req *gofive.CreateUserRequest, callOptions ...callopt.Option) (r *gofive.CreateUserResponse, err error)
	QueryUserByID(ctx context.Context, req *gofive.QueryUserByIDRequest, callOptions ...callopt.Option) (r *gofive.QueryUserResponse, err error)
	QueryUserByName(ctx context.Context, req *gofive.QueryUserByNameRequest, callOptions ...callopt.Option) (r *gofive.QueryUserResponse, err error)
	CheckUser(ctx context.Context, req *gofive.CheckUserRequest, callOptions ...callopt.Option) (r *gofive.CheckUserResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) CreateUser(ctx context.Context, req *gofive.CreateUserRequest, callOptions ...callopt.Option) (r *gofive.CreateUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateUser(ctx, req)
}

func (p *kUserServiceClient) QueryUserByID(ctx context.Context, req *gofive.QueryUserByIDRequest, callOptions ...callopt.Option) (r *gofive.QueryUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryUserByID(ctx, req)
}

func (p *kUserServiceClient) QueryUserByName(ctx context.Context, req *gofive.QueryUserByNameRequest, callOptions ...callopt.Option) (r *gofive.QueryUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryUserByName(ctx, req)
}

func (p *kUserServiceClient) CheckUser(ctx context.Context, req *gofive.CheckUserRequest, callOptions ...callopt.Option) (r *gofive.CheckUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckUser(ctx, req)
}
