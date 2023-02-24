
package pack

import (
	"errors"
	// "time"
	// "log"
	"github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/user"
	"github.com/YANGJUNYAN0715/douyin/tree/main/pkg/errno"
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
	// log.Println("++++++++++++=====",err.ErrCode,"================================")
	// log.Println("++++++++++++=====",err.ErrMsg,"================================")
	return &user.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}

// func BuildBaseVideoResp(err error) *user.BaseResp {
// 	if err == nil {
// 		return baseResp(errno.Success)
// 	}

// 	e := errno.ErrNo{}
// 	if errors.As(err, &e) {
// 		return baseResp(e)
// 	}

// 	s := errno.ServiceErr.WithMessage(err.Error())
// 	return baseVideoResp(s)
// }

// func baseVideoResp(err errno.ErrNo) *user.BaseResp {
// 	return &user.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
// }
