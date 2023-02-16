package handlers

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sakasikai/GoFive/pkg/errno"
)

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type LoginResp struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_msg"`
	UserId  int64  `json:"user_id"`
	Token   string `json:"token"`
}

type UserInfoResp struct {
	Code    int64       `json:"status_code"`
	Message string      `json:"status_msg"`
	Data    interface{} `json:"data"`
}

// SendResponse pack response
func SendLoginResponse(c *app.RequestContext, err error, userId int64, token string) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, LoginResp{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		UserId:  userId,
		Token:   token,
	})
}

func SendUserInfoResponse(c *app.RequestContext, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, UserInfoResp{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}
