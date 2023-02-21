package main

import (
	"context"
<<<<<<< HEAD
	"fmt"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/interact/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/interact/service"
	"time"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/interact/dal/db"
	"github.com/cloudwego/kitex/tool/internal_pkg/log"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
	interact "github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/interact"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/interact/rpc"
=======
	// "fmt"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/interact/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/interact/service"




	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
	interact "github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/interact"
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
)

// InteractServiceImpl implements the last service interface defined in the IDL.
type InteractServiceImpl struct{}

<<<<<<< HEAD

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *InteractServiceImpl) FavoriteAction(ctx context.Context, req *interact.FavoriteActionRequest) (resp *interact.FavoriteActionResponse, err error) {
	resp = new(interact.FavoriteActionResponse)

	if len(req.Token) == 0 || req.VideoId == 0 || req.ActionType == 0 {
		resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg
=======
// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *interact.FavoriteActionRequest) (resp *interact.FavoriteActionResponse, err error) {
	resp = new(interact.FavoriteActionResponse)

	if len(req.Token) == 0 || req.VideoId == 0 || req.ActionType == 0 {
		resp = pack.BuildFavoriteBaseResp(errno.ParamErr)
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
		return resp, nil
	}

	err = service.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
<<<<<<< HEAD
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}
	resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg
=======
		resp = pack.BuildFavoriteBaseResp(err)
		return resp, nil
	}
	resp = pack.BuildFavoriteBaseResp(errno.Success)
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	return resp, nil
	return
}

// FavoriteList implements the FavoriteServiceImpl interface.
<<<<<<< HEAD
func (s *InteractServiceImpl) FavoriteList(ctx context.Context, req *interact.FavoriteListRequest) (resp *interact.FavoriteListResponse, err error) {
	resp = new(interact.FavoriteListResponse)

	if req.UserId == 0 {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
=======
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *interact.FavoriteListRequest) (resp *interact.FavoriteListResponse, err error) {
	resp = new(interact.FavoriteListResponse)

	if req.UserId == 0 {
		resp = pack.BuildFavoriteListBaseResp(errno.ParamErr)
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
		return resp, nil
	}

	videoList, err := service.NewFavoriteListService(ctx).FavoriteList(req)
	if err != nil {
<<<<<<< HEAD
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
=======
		resp = pack.BuildFavoriteListBaseResp(err)
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
		return resp, nil
	}

	resp = pack.BuildFavoriteListBaseResp(errno.Success)
	resp.VideoList = videoList
	return resp, nil
}

<<<<<<< HEAD

func packErr1(err error) *interact.CommentActionResponse {
	msg := err.Error()
	return &interact.CommentActionResponse{StatusCode: errno.CommentError, StatusMsg: msg}
=======
// CommentSrvImpl implements the last service interface defined in the IDL.
type CommentSrvImpl struct{}

func packErr1(err error) *interact.CommentActionResponse {
	msg := err.Error()
	return &interact.CommentActionResponse{StatusCode: errno.CommentError, StatusMsg: &msg}
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
}

func packErr2(err error) *interact.CommentListResponse {
	msg := err.Error()
<<<<<<< HEAD
	return &interact.CommentListResponse{StatusCode: errno.SuccessCode, StatusMsg: msg,
=======
	return &interact.CommentListResponse{StatusCode: errno.SuccessCode, StatusMsg: &msg,
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
		CommentList: []*interact.Comment{}}
}

func getUser(ctx context.Context, id int, token string) (*interact.User, error) {
<<<<<<< HEAD
	resp, err := rpc.Info(ctx, &user.UserInfoRequest{Token: token, UserId: int64(id)})
=======
	resp, err := rpc.Info(ctx, &user.UserRequest{Token: token, UserId: int64(id)})
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	if err != nil {
		fmt.Println(err)
		panic(err)
		return nil, err
	}
	return &interact.User{Id: resp.Id, Name: resp.Name, FollowCount: resp.FollowCount,
		FollowerCount: resp.FollowerCount, IsFollow: resp.IsFollow}, nil
}

// TODO service
// CommentAction implements the CommentSrvImpl interface.
<<<<<<< HEAD
func (s *InteractServiceImpl) CommentAction(ctx context.Context, req *interact.CommentActionRequest) (resp *interact.CommentActionResponse, err error) {
=======
func (s *CommentSrvImpl) CommentAction(ctx context.Context, req *interact.CommentActionRequest) (resp *interact.CommentActionResponse, err error) {
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	// TODO: Your code here...
	log.Info("get interact action req", *req)
	resp = new(interact.CommentActionResponse)

	//TODO check video id
	if req.ActionType == 1 {
<<<<<<< HEAD
		cmt := &db.Comment{UserId: int(req.UserId), Content: req.CommentText,
=======
		cmt := &db.Comment{UserId: int(req.UserId), Content: *req.CommentText,
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
			VideoId: int(req.VideoId), IsValid: true, CreateTime: time.Now().String()}
		if err := db.CreateComment(ctx, cmt); err != nil {
			return packErr1(err), nil
		}
		user, err := getUser(ctx, cmt.UserId, req.Token)
		if err != nil {
			return nil, err
		}
		return &interact.CommentActionResponse{StatusCode: errno.SuccessCode,
			Comment: &interact.Comment{Id: int64(cmt.ID), User: user, Content: cmt.Content, CreateDate: cmt.CreateTime}}, nil
	} else if req.ActionType == 2 {
<<<<<<< HEAD
		cmt := &db.Comment{ID: uint(req.CommentId)}
=======
		cmt := &db.Comment{ID: uint(*req.CommentId)}
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
		tmp, err := db.SelectComment(ctx, int(cmt.ID))
		if err != nil {
			return packErr1(err), nil
		}
		if tmp == nil {
			return &interact.CommentActionResponse{StatusCode: errno.CommentNotFound}, nil
		}
		if err := db.DeleteComment(ctx, cmt); err != nil {
			return packErr1(err), nil
		}
		user, err := getUser(ctx, tmp.UserId, req.Token)
		if err != nil {
			return nil, err
		}
		return &interact.CommentActionResponse{StatusCode: errno.SuccessCode,
			Comment: &interact.Comment{Id: int64(tmp.ID), User: user, Content: tmp.Content, CreateDate: tmp.CreateTime}}, nil
	} else {
		msg := "err"
<<<<<<< HEAD
		return &interact.CommentActionResponse{StatusCode: errno.ActionTypeErrCode, StatusMsg: msg}, nil
=======
		return &interact.CommentActionResponse{StatusCode: errno.ActionTypeErrCode, StatusMsg: &msg}, nil
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	}
}

// CommentList implements the CommentSrvImpl interface.
<<<<<<< HEAD
func (s *InteractServiceImpl) CommentList(ctx context.Context, req *interact.CommentListRequest) (resp *interact.CommentListResponse, err error) {
=======
func (s *CommentSrvImpl) CommentList(ctx context.Context, req *interact.CommentListRequest) (resp *interact.CommentListResponse, err error) {
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	log.Info("get interact list req", *req)
	// TODO: Your code here...
	resp = new(interact.CommentListResponse)
	cmts, err := db.QueryComments(ctx, int(req.VideoId))
	if err != nil {
		return packErr2(err), nil
	}
	res := []*interact.Comment{}
	for _, c := range cmts {
		fmt.Println("get user , id = ", c.UserId)
		user, err := getUser(ctx, c.UserId, req.Token)
		if err != nil {
			return nil, err
		}
		tmp := &interact.Comment{Id: int64(c.ID), Content: c.Content,
			CreateDate: c.CreateTime, User: user}
		res = append(res, tmp)
	}
	return &interact.CommentListResponse{StatusCode: errno.SuccessCode, CommentList: res}, nil
}
