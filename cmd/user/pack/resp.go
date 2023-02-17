
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

// func loginResp(err errno.ErrNo) *user.LoginUserResponse {
// 	return &user.LoginUserResponse{UserId: ,StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
// }
