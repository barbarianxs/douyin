
package service

import (
	"context"
	// "crypto/md5"
	// "fmt"
	// "io"
	// "log"
	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/user/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/user/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/user"
	// "github.com/YANGJUNYAN0715/douyin/tree/main/pkg/errno"
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
	// log.Println(videos[0].PlayUrl)
	return videos, nil
}