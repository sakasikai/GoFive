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
	//
	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[constants.IdentityKey].(float64))

	var uidParam struct {
		id int64 `json:"user_id"`
	}

	if err := c.Bind(&uidParam); err != nil {
		SendUserInfoResponse(c, errno.ConvertErr(err), nil)
	}

	if uidParam.id != userID {
		SendUserInfoResponse(c, errno.ParamErr, nil)
	}

	req := &gofive.QueryUserByIDRequest{UserId: userID}
	users, err := rpc.QueryUsersByID(context.Background(), req)
	SendUserInfoResponse(c, err, users)
}
