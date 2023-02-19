package main

import (
	"context"
	"fmt"
	"time"

	"github.com/YANGJUNYAN0715/douyin/tree/zhao/cmd/comment/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/cmd/comment/rpc"
	comment "github.com/YANGJUNYAN0715/douyin/tree/zhao/kitex_gen/comment"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/kitex_gen/user"
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
	return &comment.DouyinCommentListResponse{StatusCode: errno.SuccessCode, StatusMsg: &msg,
		CommentList: []*comment.Comment{}}
}

func getUser(ctx context.Context, id int, token string) (*comment.User, error) {
	resp, err := rpc.Info(ctx, &user.DouyinUserRequest{Token: token, UserId: int64(id)})
	if err != nil {
		fmt.Println(err)
		panic(err)
		return nil, err
	}
	return &comment.User{Id: resp.Id, Name: resp.Name, FollowCount: resp.FollowCount,
		FollowerCount: resp.FollowerCount, IsFollow: resp.IsFollow}, nil
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
		user, err := getUser(ctx, cmt.UserId, req.Token)
		if err != nil {
			return nil, err
		}
		return &comment.DouyinCommentActionResponse{StatusCode: errno.SuccessCode,
			Comment: &comment.Comment{Id: int64(cmt.ID), User: user, Content: cmt.Content, CreateDate: cmt.CreateTime}}, nil
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
		user, err := getUser(ctx, tmp.UserId, req.Token)
		if err != nil {
			return nil, err
		}
		return &comment.DouyinCommentActionResponse{StatusCode: errno.SuccessCode,
			Comment: &comment.Comment{Id: int64(tmp.ID), User: user, Content: tmp.Content, CreateDate: tmp.CreateTime}}, nil
	} else {
		msg := "err"
		return &comment.DouyinCommentActionResponse{StatusCode: errno.ActionTypeErrCode, StatusMsg: &msg}, nil
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
		fmt.Println("get user , id = ", c.UserId)
		user, err := getUser(ctx, c.UserId, req.Token)
		if err != nil {
			return nil, err
		}
		tmp := &comment.Comment{Id: int64(c.ID), Content: c.Content,
			CreateDate: c.CreateTime, User: user}
		res = append(res, tmp)
	}
	return &comment.DouyinCommentListResponse{StatusCode: errno.SuccessCode, CommentList: res}, nil
}
