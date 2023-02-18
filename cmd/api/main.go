package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
	"github.com/sakasikai/GoFive/cmd/api/handlers"
	"github.com/sakasikai/GoFive/cmd/api/rpc"
	gofive "github.com/sakasikai/GoFive/kitex_gen/GoFive"
	"github.com/sakasikai/GoFive/pkg/constants"
	"github.com/sakasikai/GoFive/pkg/errno"
	"time"
)

func Init() {
	rpc.InitRPC()
}

func main() {
	Init()
	r := server.New(
		server.WithHostPorts("127.0.0.1:8080"),
		server.WithHandleMethodNotAllowed(true),
	)

	authMiddleware, _ := jwt.New(&jwt.HertzJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch e.(type) {
			case errno.ErrNo:
				return e.(errno.ErrNo).ErrMsg
			default:
				return e.Error()
			}
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			var loginVar handlers.UserParam
			if err := c.Bind(&loginVar); err != nil {
				c.JSON(consts.StatusOK, map[string]interface{}{
					"status_code": errno.ParamErr.ErrCode,
					"status_msg":  errno.ParamErr.ErrMsg,
				})
			}

			if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
				c.JSON(consts.StatusOK, map[string]interface{}{
					"status_code": errno.ParamErr.ErrCode,
					"status_msg":  errno.ParamErr.ErrMsg,
				})
			}

			users, _ := rpc.QueryUsersByName(
				context.Background(),
				&gofive.QueryUserByNameRequest{UserName: loginVar.UserName},
			)

			c.JSON(consts.StatusOK, map[string]interface{}{
				"status_code": errno.Success.ErrCode,
				"status_msg":  errno.Success.ErrMsg,
				"user_id":     users[0].User.Id,
				"token":       token,
			})
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"code":    errno.AuthorizationFailedErrCode,
				"message": message,
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVar handlers.UserParam
			if err := c.Bind(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			return rpc.CheckUser(context.Background(), &gofive.CheckUserRequest{UserName: loginVar.UserName, Password: loginVar.PassWord})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	// for token generation
	handlers.SetAuthMiddleware(authMiddleware)

	// todo recovery

	v := r.Group("/douyin")

	// basic
	v.POST("/user/register", handlers.Register)
	v.POST("/user/login/", authMiddleware.LoginHandler)

	v.Use(authMiddleware.MiddlewareFunc())
	{
		//v.GET("/feed/")
		v.GET("/user/", handlers.UserInfo)

		//v.POST("/publish/action/")
		//v.GET("/publish/list/")
	}

	// interact
	//socialize

	r.NoRoute(func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "no route")
	})
	r.NoMethod(func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "no method")
	})
	r.Spin()
}
