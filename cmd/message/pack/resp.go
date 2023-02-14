
package pack

import (
	"errors"
	"time"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/message"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *message.BaseResp {
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

func baseResp(err errno.ErrNo) *message.BaseResp {
	return &message.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}
