package service

import (
	"context"
	"errors"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/interact/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/interact/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/interact"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/consts"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/jwt"
	"sync"
	"log"
)

type CommentListService struct {
	ctx context.Context
}

// NewCommentListService new CommentListService
func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{ctx: ctx}
}

// CommentList get video information that users guoke
func (s *CommentListService) CommentList(req *interact.CommentListRequest) ([]*interact.Video, error) {
	log.Info("get interact list req", *req)
	// TODO: Your code here...
	resp = new(interact.CommentListResponse)
	cmts, err := db.QueryComments(s.ctx, int64(req.VideoId))
	if err != nil {
		return packErr2(err), nil
	}
	res := []*interact.Comment{}
	for _, c := range cmts {
		fmt.Println("get user , id = ", c.UserId)
		user, err := db.QueryUserInfo(s.ctx, c.UserId)
		if err != nil {
			return nil, err
		}
		tmp := &interact.Comment{Id: int64(c.ID), Content: c.Content,
			CreateDate: c.CreateDate, User: user}
		res = append(res, tmp)
	}
	return &interact.CommentListResponse{StatusCode: errno.SuccessCode, CommentList: res}, nil



}