package main

import (
	"context"
	"log"

	feed "github.com/YANGJUNYAN0715/douyin/tree/zhao/kitex_gen/feed"
	//"github.com/YANGJUNYAN0715/douyin/tree/zhao/cmd/feed/jwt"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/cmd/feed/service"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/pkg/errno"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/cmd/feed/pack"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetUserFeed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetUserFeed(ctx context.Context, req *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	var uid int64 = 0
	if req.Token != "" {
		claim, err := Jwt.ParseToken(req.Token)
		if err != nil {
			resp = pack.BuildVideoResp(errno.Token2UserIdErr)
			return resp, nil
		} else {
			uid = claim.Id
		}
	}
	//uid = req.Userid
	log.Println("--------uid------------")
	log.Println(uid)
	vis, nextTime, err := service.NewGetUserFeedService(ctx).GetUserFeed(req, uid)
	if err != nil {
		resp = pack.BuildVideoResp(err)
		return resp, nil
	}

	resp = pack.BuildVideoResp(errno.Success)
	resp.VideoList = vis
	resp.NextTime = nextTime
	return resp, nil
}
