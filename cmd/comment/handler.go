package main

import (
	"context"
	"time"

	"github.com/YANGJUNYAN0715/douyin/tree/zhao/cmd/api/biz/model/api"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/cmd/comment/dal/db"
	comment "github.com/YANGJUNYAN0715/douyin/tree/zhao/kitex_gen/comment"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/pkg/consts"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/pkg/errno"
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

// TODO service
// CommentAction implements the CommentSrvImpl interface.
func (s *CommentSrvImpl) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	// TODO: Your code here...
	resp = new(comment.DouyinCommentActionResponse)
	user := ctx.Value(consts.IdentityKey).(api.User)
	//TODO check video id
	if req.ActionType == 1 {
		if err := db.CreateComment(ctx, []*db.Comment{&db.Comment{UserId: int(user.ID), Content: *req.CommentText,
			VideoId: int(req.VideoId), IsValid: true, CreateTime: time.Now().String()}}); err != nil {
			return packErr1(err), nil
		}
	} else if req.ActionType == 2 {
		cmt := &db.Comment{ID: uint(*req.CommentId)}
		tmp, err := db.SelectComment(ctx, cmt)
		if err != nil {
			return packErr1(err), nil
		}
		if tmp == nil {
			return &comment.DouyinCommentActionResponse{StatusCode: errno.CommentNotFound}, nil
		}
		if err := db.DeleteComment(ctx, cmt); err != nil {
			return packErr1(err), nil
		}
	} else {
		return &comment.DouyinCommentActionResponse{StatusCode: errno.ActionTypeErrCode}, nil
	}
	return &comment.DouyinCommentActionResponse{StatusCode: errno.SuccessCode}, nil
}

// CommentList implements the CommentSrvImpl interface.
func (s *CommentSrvImpl) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	resp = new(comment.DouyinCommentListResponse)
	cmts, err := db.QueryComments(ctx, int(req.VideoId))
	if err != nil {
		return packErr2(err), nil
	}
	res := []*comment.Comment{}
	for _, c := range cmts {
		//TODO rpc
		tmp := &comment.Comment{Id: int64(c.ID), Content: c.Content,
			CreateDate: c.CreateTime, User: &comment.User{Id: 1, Name: "dsa", FollowCount: 1,
				FollowerCount: 1, IsFollow: true}}
		res = append(res, tmp)
	}
	return &comment.DouyinCommentListResponse{StatusCode: errno.SuccessCode, CommentList: res}, nil
}
