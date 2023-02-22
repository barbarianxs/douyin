package service

import (
	"context"
	"time"
	"log"
	// "crypto/md5"
	// "fmt"
	// "io"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/user/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
)

const (
	LIMIT = 30 // 单次返回最大视频数
)

type GetUserFeedService struct {
	ctx context.Context
}

// NewGetUserByIdService new GetUserByIdService
func NewGetUserFeedService(ctx context.Context) *GetUserFeedService {
	return &GetUserFeedService{ctx: ctx}
}

// get user info.
func (s *GetUserFeedService) GetUserFeed(req *user.FeedRequest) (vis []*user.Video, nextTime int64, err error) {
	log("----------------------kitex feed--------------------------------------")
	videos, err := db.MGetVideos(s.ctx, LIMIT, req.LatestTime)
	log.Println("-------------req.LatestTime----------")
	log.Println(req.LatestTime)
	if err != nil {
		return vis, nextTime, err
	}

	if len(videos) == 0 {
		nextTime = time.Now().UnixMilli()
		return vis, nextTime, nil
	} else {
		nextTime = videos[len(videos)-1].UpdatedAt.UnixMilli()
	}

	if vis, err = db.BuildVideos(s.ctx, videos, &req.UserId); err != nil {
		nextTime = time.Now().UnixMilli()
		return vis, nextTime, err
	}

	return vis, nextTime, nil
}