package service

import (
	"context"
	"github.com/sakasikai/GoFive/cmd/user/dal/db"
	gofive "github.com/sakasikai/GoFive/kitex_gen/GoFive"
)

type QueryUserService struct {
	ctx context.Context
}

func NewQueryUserService(ctx context.Context) *QueryUserService {
	return &QueryUserService{ctx: ctx}
}

func (s *QueryUserService) QueryUserByID(req *gofive.QueryUserByIDRequest) *db.User {
	user, err := db.QueryUserByID(s.ctx, req.UserId)
	if err != nil {
		return nil
	}
	return user[0]
}

func (s *QueryUserService) QueryUserByName(req *gofive.QueryUserByNameRequest) *db.User {
	user, err := db.QueryUserByName(s.ctx, req.UserName)
	if err != nil {
		return nil
	}
	return user[0]
}
