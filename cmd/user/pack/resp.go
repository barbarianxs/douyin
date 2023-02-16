
package pack

import (
	"errors"

	"github.com/YANGJUNYAN0715/douyin/tree/zhao/kitex_gen/user"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/pkg/errno"
)
// 官方接口 仅使用douyinregisterresp的打包格式
// BuildDouyinUserRegisterResponse build DouyinUserRegisterResponse from error
func BuildDouyinUserRegisterResponse(err error) *user.DouyinUserRegisterResponse {
	if err == nil {
		return DouyinUserRegisterResponse(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return DouyinUserRegisterResponse(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return DouyinUserRegisterResponse(s)
}

func DouyinUserRegisterResponse(err errno.ErrNo) *user.DouyinUserRegisterResponse {
	return &user.DouyinUserRegisterResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

// user info response 
func BuildDouyinUserResponse(err error) *user.DouyinUserResponse {
	if err == nil {
		return DouyinUserResponse(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return DouyinUserResponse(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return DouyinUserResponse(s)
}

func DouyinUserResponse(err errno.ErrNo) *user.DouyinUserResponse {
	return &user.DouyinUserResponse{
		StatusCode: int32(err.ErrCode),
		StatusMsg: err.ErrMsg,
	}
}