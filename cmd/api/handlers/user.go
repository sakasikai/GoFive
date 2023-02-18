package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"github.com/sakasikai/GoFive/cmd/api/rpc"
	gofive "github.com/sakasikai/GoFive/kitex_gen/GoFive"
	"github.com/sakasikai/GoFive/pkg/constants"
	"github.com/sakasikai/GoFive/pkg/errno"
)

func UserInfo(ctx context.Context, c *app.RequestContext) {
	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[constants.IdentityKey].(float64))

	users, err := rpc.QueryUsersByID(context.Background(), &gofive.QueryUserByIDRequest{UserId: userID})

	if err != nil || len(users) == 0 {
		SendUserInfoResponse(c, errno.ConvertErr(err), nil)
	}

	SendUserInfoResponse(c, errno.ConvertErr(errno.Success), users[0].User)
}
