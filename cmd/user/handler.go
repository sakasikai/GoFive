package main

import (
	"context"
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
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
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
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *gofive.CheckUserRequest) (resp *gofive.CheckUserResponse, err error) {
	resp = new(gofive.CheckUserResponse)

	if len(req.UserName) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserId = uid
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
