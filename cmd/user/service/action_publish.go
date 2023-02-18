
package service

import (
	"context"
	// "crypto/md5"
	// "fmt"
	// "io"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/user/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/user/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
)

type PublishActionService struct {
	ctx context.Context
}

// NewPublishActionService new CreateUserService
func NewPublishActionService(ctx context.Context) *PublishActionService {
	return &PublishActionService{ctx: ctx}
}

/// PublishActionService query user info
func (s *PublishActionService) PublishAction(req *user.PublishActionRequest) error{
	

	
	VideoModel := &db.Video{
		AuthorID:   req.UserId,
		PlayUrl:  req.FileUrl,
		CoverUrl: req.CoverUrl,
		Title: req.Title,
	}
	return db.CreateVideo(s.ctx, []*db.Video{VideoModel})
}