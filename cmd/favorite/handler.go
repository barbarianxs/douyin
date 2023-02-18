package main

import (
	"context"
	"github.com/YANGJUNYAN0715/douyin/tree/li/cmd/favorite/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/li/cmd/favorite/service"
	"github.com/YANGJUNYAN0715/douyin/tree/li/kitex_gen/favorite"
	"github.com/YANGJUNYAN0715/douyin/tree/li/pkg/errno"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	resp = new(favorite.FavoriteActionResponse)

	if len(req.Token) == 0 || req.VideoId == 0 || req.ActionType == 0 {
		resp = pack.BuildFavoriteBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		resp = pack.BuildFavoriteBaseResp(err)
		return resp, nil
	}
	resp = pack.BuildFavoriteBaseResp(errno.Success)
	return resp, nil
	return
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	resp = new(favorite.FavoriteListResponse)

	if req.UserId == 0 {
		resp = pack.BuildFavoriteListBaseResp(errno.ParamErr)
		return resp, nil
	}

	videoList, err := service.NewFavoriteListService(ctx).FavoriteList(req)
	if err != nil {
		resp = pack.BuildFavoriteListBaseResp(err)
		return resp, nil
	}

	resp = pack.BuildFavoriteListBaseResp(errno.Success)
	resp.VideoList = videoList
	return resp, nil
}
