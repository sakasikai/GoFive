package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/sakasikai/GoFive/cmd/api/rpc"
	gofive "github.com/sakasikai/GoFive/kitex_gen/GoFive"
	"github.com/sakasikai/GoFive/pkg/errno"
)

func Register(ctx context.Context, c *app.RequestContext) {
	var registerVar UserParam
	if err := c.Bind(&registerVar); err != nil {
		SendRegisterResponse(c, errno.ConvertErr(err), -1, "nil")
		return
	}

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		//println(registerVar.UserName, "or", registerVar.PassWord)
		SendRegisterResponse(c, errno.ParamErr, -1, "nil")
		return
	}

	// userName, passWord 都不空
	err := rpc.CreateUser(context.Background(), &gofive.CreateUserRequest{
		UserName: registerVar.UserName,
		Password: registerVar.PassWord,
	})

	if err != nil {
		SendRegisterResponse(c, errno.ConvertErr(err), -1, "nil")
		return
	}

	resp, err := rpc.QueryUsersByName(context.Background(), &gofive.QueryUserByNameRequest{
		UserName: registerVar.UserName,
	})

	if err != nil {
		SendRegisterResponse(c, errno.ConvertErr(err), -1, "nil")
		return
	}

	response := resp[0]

	// flow same as loginHandler
	token, _, err := jwtAuthMiddleware.TokenGenerator(response.User.Id)

	if err != nil {
		SendRegisterResponse(c, errno.ConvertErr(err), -1, "nil")
		return
	}

	SendRegisterResponse(c, errno.Success, response.User.Id, token)
}
