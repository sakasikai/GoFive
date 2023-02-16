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

type CreateUserService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

func (s *CreateUserService) CreateUser(req *gofive.CreateUserRequest) error {
	users, err := db.QueryUserByName(s.ctx, req.UserName)
	if err != nil {
		return err
	}

	// db 存在，是重复创建
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	// db 不存在，才会创建
	// hd5加密的password，存入db
	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}

	passWord := fmt.Sprintf("%x", h.Sum(nil))
	return db.CreateUser(s.ctx, []*db.User{{
		UserName: req.UserName,
		Password: passWord,
	}}) // double brackets
}
