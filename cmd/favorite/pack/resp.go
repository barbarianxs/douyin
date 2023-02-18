package pack

import (
	"errors"
	"github.com/YANGJUNYAN0715/douyin/tree/li/kitex_gen/favorite"
	"github.com/YANGJUNYAN0715/douyin/tree/li/pkg/errno"
)

// BuildFavoriteBaseResp build favorite baseResp from error
func BuildFavoriteBaseResp(err error) *favorite.FavoriteActionResponse {
	if err == nil {
		return favoritebaseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return favoritebaseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return favoritebaseResp(s)
}

func favoritebaseResp(err errno.ErrNo) *favorite.FavoriteActionResponse {
	return &favorite.FavoriteActionResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}

func BuildFavoriteListBaseResp(err error) *favorite.FavoriteListResponse {
	if err == nil {
		return favoriteListbaseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return favoriteListbaseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return favoriteListbaseResp(s)
}

func favoriteListbaseResp(err errno.ErrNo) *favorite.FavoriteListResponse {
	return &favorite.FavoriteListResponse{StatusCode: err.ErrCode, StatusMsg: &(err.ErrMsg)}
}
