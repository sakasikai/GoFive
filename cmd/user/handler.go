package main

import (
	"context"
	"fmt"
	"github.com/sakasikai/GoFive/cmd/user/pack"
	"github.com/sakasikai/GoFive/cmd/user/service"
	gofive "github.com/sakasikai/GoFive/kitex_gen/GoFive"
	"github.com/sakasikai/GoFive/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *gofive.CreateUserRequest) (resp *gofive.CreateUserResponse, err error) {
	resp = new(gofive.CreateUserResponse)
	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// QueryUserByID implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUserByID(ctx context.Context, req *gofive.QueryUserByIDRequest) (resp *gofive.QueryUserResponse, err error) {
	resp = new(gofive.QueryUserResponse)
	u := service.NewQueryUserService(ctx).QueryUserByID(req)
	resp.User = &gofive.User{
		Id:            int64(u.ID),
		Name:          u.UserName,
		FollowCount:   u.FollowerCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
	resp.BaseResp = pack.BaseResp(errno.Success)
	return resp, nil
}

// QueryUserByName implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUserByName(ctx context.Context, req *gofive.QueryUserByNameRequest) (resp *gofive.QueryUserResponse, err error) {
	resp = new(gofive.QueryUserResponse)
	u := service.NewQueryUserService(ctx).QueryUserByName(req)
	resp.User = &gofive.User{
		Id:            int64(u.ID),
		Name:          u.UserName,
		FollowCount:   u.FollowerCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
	resp.BaseResp = pack.BaseResp(errno.Success)
	return resp, nil
	return
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *gofive.CheckUserRequest) (resp *gofive.CheckUserResponse, err error) {
	// TODO: Your code here...
	fmt.Println("ok")
	return
}
