
package pack

import (
	"errors"
	// "time"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildLoginResp(err error) *user.LoginUserResponse {
	if err == nil {
		return LoginResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return LoginResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return LoginResp(s)
}

func LoginResp(err errno.ErrNo) *user.LoginUserResponse {
	return &user.LoginUserResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}

// BuildRegisterResp build baseResp from error
func BuildRegisterResp(err error) *user.RegisterUserResponse {
	if err == nil {
		return RegisterResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return RegisterResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return RegisterResp(s)
}

func RegisterResp(err errno.ErrNo) *user.RegisterUserResponse {
	return &user.RegisterUserResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
// func loginResp(err errno.ErrNo) *user.LoginUserResponse {
// 	return &user.LoginUserResponse{UserId: ,StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
// }
