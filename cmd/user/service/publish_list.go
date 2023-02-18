
package service

import (
	"context"
	// "crypto/md5"
	// "fmt"
	// "io"
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
	
	videos := pack.Videos(videoModels)
	
	return videos, nil
}