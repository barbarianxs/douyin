package main

import (
	"context"
	"time"

	"github.com/YANGJUNYAN0715/douyin/tree/zhao/cmd/comment/dal/db"
	comment "github.com/YANGJUNYAN0715/douyin/tree/zhao/kitex_gen/comment"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/pkg/errno"
	"github.com/cloudwego/kitex/tool/internal_pkg/log"
)

// CommentSrvImpl implements the last service interface defined in the IDL.
type CommentSrvImpl struct{}

func packErr1(err error) *comment.DouyinCommentActionResponse {
	msg := err.Error()
	return &comment.DouyinCommentActionResponse{StatusCode: errno.CommentError, StatusMsg: &msg}
}

func packErr2(err error) *comment.DouyinCommentListResponse {
	msg := err.Error()
	return &comment.DouyinCommentListResponse{StatusCode: errno.CommentError, StatusMsg: &msg}
}

// TODO rpc
func getUser(id int) *comment.User {
	return &comment.User{Id: 1, Name: "dsa", FollowCount: 1, FollowerCount: 1, IsFollow: true}
}

// TODO service
// CommentAction implements the CommentSrvImpl interface.
func (s *CommentSrvImpl) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	// TODO: Your code here...
	log.Info("get comment action req", *req)
	resp = new(comment.DouyinCommentActionResponse)

	//TODO check video id
	if req.ActionType == 1 {
		cmt := &db.Comment{UserId: int(req.UserId), Content: *req.CommentText,
			VideoId: int(req.VideoId), IsValid: true, CreateTime: time.Now().String()}
		if err := db.CreateComment(ctx, cmt); err != nil {
			return packErr1(err), nil
		}
		return &comment.DouyinCommentActionResponse{StatusCode: errno.SuccessCode,
			Comment: &comment.Comment{Id: int64(cmt.ID), User: getUser(cmt.UserId), Content: cmt.Content, CreateDate: cmt.CreateTime}}, nil
	} else if req.ActionType == 2 {
		cmt := &db.Comment{ID: uint(*req.CommentId)}
		tmp, err := db.SelectComment(ctx, int(cmt.ID))
		if err != nil {
			return packErr1(err), nil
		}
		if tmp == nil {
			return &comment.DouyinCommentActionResponse{StatusCode: errno.CommentNotFound}, nil
		}
		if err := db.DeleteComment(ctx, cmt); err != nil {
			return packErr1(err), nil
		}
		return &comment.DouyinCommentActionResponse{StatusCode: errno.SuccessCode,
			Comment: &comment.Comment{Id: int64(tmp.ID), User: getUser(cmt.UserId), Content: tmp.Content, CreateDate: tmp.CreateTime}}, nil
	} else {
		msg := "err"
		return &comment.DouyinCommentActionResponse{StatusCode: errno.ActionTypeErrCode, StatusMsg: &msg}, err.New
	}
}

// CommentList implements the CommentSrvImpl interface.
func (s *CommentSrvImpl) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	log.Info("get comment list req", *req)
	// TODO: Your code here...
	resp = new(comment.DouyinCommentListResponse)
	cmts, err := db.QueryComments(ctx, int(req.VideoId))
	if err != nil {
		return packErr2(err), nil
	}
	res := []*comment.Comment{}
	for _, c := range cmts {
		tmp := &comment.Comment{Id: int64(c.ID), Content: c.Content,
			CreateDate: c.CreateTime, User: getUser(c.UserId)}
		res = append(res, tmp)
	}
	return &comment.DouyinCommentListResponse{StatusCode: errno.SuccessCode, CommentList: res}, nil
}
