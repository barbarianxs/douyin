package service

import (
	"context"
	"time"
	"log"
	// "crypto/md5"
	// "fmt"
	// "io"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/user/pack"
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
	log.Println("----------------------kitex feed--------------------------------------")
	videos, err := db.MGetVideos(s.ctx, LIMIT, req.LatestTime)
	log.Println("-------------req.LatestTime----------")
	log.Println(req.LatestTime)
	log.Println(videos[0])
	if err != nil {
		return vis, nextTime, err
	}

	if len(videos) == 0 {
		nextTime = time.Now().UnixMilli()
		return vis, nextTime, nil
	} else {
		nextTime = videos[len(videos)-1].UpdatedAt.UnixMilli()
	}
	log.Println("-------------req.nextTime----------")
	log.Println(nextTime)
	// if vis, err = db.BuildVideos(s.ctx, videos, &req.UserId); err != nil {
	// 	nextTime = time.Now().UnixMilli()
	// 	return vis, nextTime, err
	// }
	//查询视频作者信息
	nextTime = time.Now().UnixMilli()
	pack_videos := make([]*user.Video, 0) 
	for index, val := range videos{
		users, err := db.QueryUserInfo(s.ctx, videos[index].AuthorID)
		u := users[0]
		if err != nil{
			return nil, 0, err
		}
		
		if temp := pack.Video(val, u); temp != nil{
			pack_videos = append(pack_videos, temp)

		}
		

	}

	return pack_videos, nextTime, nil
}