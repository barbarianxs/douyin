
package pack

import (
	"errors"
	// "time"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/interact"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
)
// BuildBaseResp build baseResp from error
<<<<<<< HEAD
func BuildBaseResp(err error) *interact.BaseResp {
=======
func BuildBaseResp(err error) *user.BaseResp {
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
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

<<<<<<< HEAD
func baseResp(err errno.ErrNo) *interact.BaseResp {
	return &interact.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
// BuildFavoriteBaseResp build interact baseResp from error
func BuildFavoriteBaseResp(err error) *interact.FavoriteActionResponse {
	if err == nil {
		return interactbaseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return interactbaseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return interactbaseResp(s)
}

func interactbaseResp(err errno.ErrNo) *interact.FavoriteActionResponse {
	return &interact.FavoriteActionResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}

func BuildFavoriteListBaseResp(err error) *interact.FavoriteListResponse {
	if err == nil {
		return interactListbaseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return interactListbaseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return interactListbaseResp(s)
}

func interactListbaseResp(err errno.ErrNo) *interact.FavoriteListResponse {
	return &interact.FavoriteListResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
=======
func baseResp(err errno.ErrNo) *user.BaseResp {
	return &user.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
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
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
