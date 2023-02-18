package pack

import (
	"errors"
	gofive "github.com/sakasikai/GoFive/kitex_gen/GoFive"
	"github.com/sakasikai/GoFive/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *gofive.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *gofive.BaseResp {
	return &gofive.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
