
package service

import (
	"context"
	// "crypto/md5"
	// "fmt"
	// "io"
<<<<<<< HEAD
<<<<<<< HEAD
	// "log"
=======
>>>>>>> origin/guo
=======
	// "log"
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/user/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/user/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
)

type PublishListService struct {
	ctx context.Context
}

// NewPublishListService new CreateUserService
func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{ctx: ctx}
}

/// PublishListService query user info
func (s *PublishListService) PublishList(req *user.PublishListRequest) ([]*user.Video, error){
	videoModels, err := db.MGetVideosOfUserIDList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	users, err := db.QueryUserInfo(s.ctx, req.UserId)
	u := users[0]
	videos := pack.Videos(videoModels, u)
<<<<<<< HEAD
<<<<<<< HEAD
	// log.Println(videos[0].PlayUrl)
=======
	
>>>>>>> origin/guo
=======
	// log.Println(videos[0].PlayUrl)
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	return videos, nil
}