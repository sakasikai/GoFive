package pack

import (
	gofive "github.com/sakasikai/GoFive/kitex_gen/GoFive"
	"github.com/sakasikai/GoFive/pkg/errno"
)

func BaseResp(err errno.ErrNo) *gofive.BaseResp {
	return &gofive.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
