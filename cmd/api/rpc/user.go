package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	gofive "github.com/sakasikai/GoFive/kitex_gen/GoFive"
	"github.com/sakasikai/GoFive/kitex_gen/GoFive/userservice"
	"github.com/sakasikai/GoFive/pkg/constants"
	"github.com/sakasikai/GoFive/pkg/errno"
	"time"
)

var userClient userservice.Client

func InitUserRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(5*time.Minute),          // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),                            // resolver
	)

	if err != nil {
		panic(err)
	}
	userClient = c
}

func CreateUser(ctx context.Context, req *gofive.CreateUserRequest) error {
	r, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	if r.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(r.BaseResp.StatusCode, r.BaseResp.StatusMsg)
	}

	return nil
}

func QueryUsersByID(ctx context.Context, req *gofive.QueryUserByIDRequest) ([]*gofive.QueryUserResponse, error) {
	resp, err := userClient.QueryUserByID(ctx, req)
	if err != nil {
		return nil, err
	}

	return []*gofive.QueryUserResponse{resp}, nil
}

func QueryUsersByName(ctx context.Context, req *gofive.QueryUserByNameRequest) ([]*gofive.QueryUserResponse, error) {
	resp, err := userClient.QueryUserByName(ctx, req)
	if err != nil {
		return nil, err
	}

	return []*gofive.QueryUserResponse{resp}, nil
}
