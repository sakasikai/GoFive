package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/sakasikai/GoFive/cmd/user/dal/db"
	gofive "github.com/sakasikai/GoFive/kitex_gen/GoFive"
	"github.com/sakasikai/GoFive/pkg/errno"
	"io"
)

type CheckUserService struct {
	ctx context.Context
}

func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{ctx: ctx}
}

func (s *CheckUserService) CheckUser(req *gofive.CheckUserRequest) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	users, err := db.QueryUserByName(s.ctx, req.UserName)
	if err != nil {
		return 0, err
	}

	if len(users) == 0 {
		return 0, errno.AuthorizationFailedErr
	}

	u := users[0]
	if u.Password != passWord {
		return 0, errno.AuthorizationFailedErr
	}
	return int64(u.ID), nil
}
