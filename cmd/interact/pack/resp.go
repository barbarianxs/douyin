
package pack

import (
	"errors"
	// "time"

	"github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/interact"
	"github.com/YANGJUNYAN0715/douyin/tree/main/pkg/errno"
)
// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *interact.BaseResp {
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