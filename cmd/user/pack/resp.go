
package pack

import (
	"errors"
	// "time"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
)
// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *user.BaseResp {
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

func baseResp(err errno.ErrNo) *user.BaseResp {
	return &user.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
<<<<<<< HEAD
=======
// // BuildLoginResp build baseResp from error
// func BuildLoginResp(err error) *user.LoginUserResponse {
// 	if err == nil {
// 		return LoginResp(errno.Success)
// 	}

// 	e := errno.ErrNo{}
// 	if errors.As(err, &e) {
// 		return LoginResp(e)
// 	}

// 	s := errno.ServiceErr.WithMessage(err.Error())
// 	return LoginResp(s)
// }

// func LoginResp(err errno.ErrNo) *user.LoginUserResponse {
// 	return &user.LoginUserResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
// }

// // BuildRegisterResp build baseResp from error
// func BuildRegisterResp(err error) *user.RegisterUserResponse {
// 	if err == nil {
// 		return RegisterResp(errno.Success)
// 	}

// 	e := errno.ErrNo{}
// 	if errors.As(err, &e) {
// 		return RegisterResp(e)
// 	}

// 	s := errno.ServiceErr.WithMessage(err.Error())
// 	return RegisterResp(s)
// }

// func RegisterResp(err errno.ErrNo) *user.RegisterUserResponse {
// 	return &user.RegisterUserResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
// }
// // BuildUserInfoResp build baseResp from error
// func BuildUserInfoResp(err error) *user.UserInfoResponse {
// 	if err == nil {
// 		return UserInfoResp(errno.Success)
// 	}

// 	e := errno.ErrNo{}
// 	if errors.As(err, &e) {
// 		return UserInfoResp(e)
// 	}

// 	s := errno.ServiceErr.WithMessage(err.Error())
// 	return UserInfoResp(s)
// }

// func UserInfoResp(err errno.ErrNo) *user.UserInfoResponse {
// 	return &user.UserInfoResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
// }
>>>>>>> origin/guo
