package service

import (
	"context"
	// "errors"
	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/interact/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/interact"
	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/interact/pack"
	// "github.com/YANGJUNYAN0715/douyin/tree/main/pkg/consts"
	// "github.com/YANGJUNYAN0715/douyin/tree/main/pkg/errno"
	// "github.com/YANGJUNYAN0715/douyin/tree/main/pkg/jwt"
	// "log"
	"errors"
	"time"
)


type CommentActionService struct {
	ctx context.Context
}

// NewCommentActionService new CommentActionService
func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}

// CommentAction implement the mainke and unmainke operations
func (s *CommentActionService) CommentAction(req *interact.CommentActionRequest) (*interact.Comment, error) {

	// log.Info("get interact action req", *req)
	// resp = new(interact.CommentActionResponse)

	//TODO check video id
	if req.ActionType == 1 {
		cmt := &db.Comment{UserId: int64(req.UserId), Content: req.CommentText,
			VideoId: int64(req.VideoId), IsValid: true, CreateDate: time.Now().Format("06-01")}
		if err := db.CreateComment(s.ctx, cmt); err != nil {
			return nil, err
		}
		user, err := db.QueryUserInfo(s.ctx, cmt.UserId)
		
		if err != nil {
			return nil, err
		}
		u := user[0]
		
		comment := pack.Comment(cmt, u)
		return comment, nil
		// return &interact.Comment{Id: int64(cmt.ID), User: user, Content: cmt.Content, CreateDate: cmt.CreateDate}, nil
	} else if req.ActionType == 2 {
		cmt := &db.Comment{ID: int64(req.CommentId)}
		tmp, err := db.SelectComment(s.ctx, int64(cmt.ID))
		if err != nil {
			return nil, err
		}
		if tmp == nil {
			return nil, err
		}
		if err := db.DeleteComment(s.ctx, cmt); err != nil {
			return nil, err
		}
		user, err := db.QueryUserInfo(s.ctx, tmp.UserId)
		if err != nil {
			return nil, err
		}
		u := user[0]
		
		
		comment := pack.Comment(tmp, u)
		return comment, nil
		// return &interact.CommentActionResponse{StatusCode: errno.SuccessCode,
		// 	Comment: &interact.Comment{Id: int64(tmp.ID), User: user, Content: tmp.Content, CreateDate: tmp.CreateDate}}, nil
	} else {
		// msg := "err"
		err := errors.New("ActionTypeErrCode")
		return nil, err
		// return &interact.CommentActionResponse{StatusCode: errno.ActionTypeErrCode, StatusMsg: msg}, nil
	}
}
