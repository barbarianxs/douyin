
package pack

import (
	"errors"
	// "time"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/relation"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *relation.BaseResp {
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

func baseResp(err errno.ErrNo) *relation.BaseResp {
	return &relation.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
