package main

import (
	"context"
	// "fmt"
	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/interact/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/interact/service"
	// "time"
	// "github.com/YANGJUNYAN0715/douyin/tree/main/cmd/interact/dal/db"
	// "github.com/cloudwego/kitex/tool/internal_pkg/log"
	// "github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/user"
	"github.com/YANGJUNYAN0715/douyin/tree/main/pkg/errno"
	interact "github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/interact"
	"log"
	// "github.com/YANGJUNYAN0715/douyin/tree/main/cmd/interact/rpc"
)

// InteractServiceImpl implements the last service interface defined in the IDL.
type InteractServiceImpl struct{}


// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *InteractServiceImpl) FavoriteAction(ctx context.Context, req *interact.FavoriteActionRequest) (resp *interact.FavoriteActionResponse, err error) {
	resp = new(interact.FavoriteActionResponse)

	if len(req.Token) == 0 || req.VideoId == 0 || req.ActionType == 0 {
		resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg
		return resp, nil
	}

	err = service.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}
	resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg
	return resp, nil
	
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *InteractServiceImpl) FavoriteList(ctx context.Context, req *interact.FavoriteListRequest) (resp *interact.FavoriteListResponse, err error) {
	resp = new(interact.FavoriteListResponse)

	if req.UserId == 0 {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}

	videoList, err := service.NewFavoriteListService(ctx).FavoriteList(req)
	if err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}

	resp = pack.BuildFavoriteListBaseResp(errno.Success)
	resp.VideoList = videoList
	return resp, nil
}
// CommentAction implements the CommentServiceImpl interface.
func (s *InteractServiceImpl) CommentAction(ctx context.Context, req *interact.CommentActionRequest) (resp *interact.CommentActionResponse, err error) {
	resp = new(interact.CommentActionResponse)

	if err = req.IsValid(); err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}


	comment, err := service.NewCommentActionService(ctx).CommentAction(req)
	if err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}
	resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg
	resp.Comment = comment
	return resp, nil
	
}


// CommentList implements the CommentSrvImpl interface.
func (s *InteractServiceImpl) CommentList(ctx context.Context, req *interact.CommentListRequest) (resp *interact.CommentListResponse, err error) {
	
	resp = new(interact.CommentListResponse)

	if req.UserId == 0 {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}
	log.Println("------------------------------kitex--commentlist-----------------------------------------------------")
	comment_list, err := service.NewCommentListService(ctx).CommentList(req)
	if err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}
	resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg

	resp.CommentList = comment_list
	return resp, nil
}


// func packErr1(err error) *interact.CommentActionResponse {
// 	msg := err.Error()
// 	return &interact.CommentActionResponse{StatusCode: errno.CommentError, StatusMsg: msg}
// }

// func packErr2(err error) *interact.CommentListResponse {
// 	msg := err.Error()
// 	return &interact.CommentListResponse{StatusCode: errno.SuccessCode, StatusMsg: msg,
// 		CommentList: []*interact.Comment{}}
// }

// func getUser(ctx context.Context, id int64, token string) (*interact.User, error) {
// 	resp, err := rpc.Info(ctx, &user.UserInfoRequest{Token: token, UserId: int64(id)})
// 	if err != nil {
// 		fmt.Println(err)
// 		panic(err)
// 		return nil, err
// 	}
// 	return &interact.User{Id: resp.Id, Name: resp.Name, FollowCount: resp.FollowCount,
// 		FollowerCount: resp.FollowerCount, IsFollow: resp.IsFollow}, nil
// }

// // TODO service
// // CommentAction implements the CommentSrvImpl interface.
// func (s *InteractServiceImpl) CommentAction(ctx context.Context, req *interact.CommentActionRequest) (resp *interact.CommentActionResponse, err error) {
// 	// TODO: Your code here...
// 	log.Info("get interact action req", *req)
// 	resp = new(interact.CommentActionResponse)

// 	//TODO check video id
// 	if req.ActionType == 1 {
// 		cmt := &db.Comment{UserId: int64(req.UserId), Content: req.CommentText,
// 			VideoId: int64(req.VideoId), IsValid: true, CreateDate: time.Now().Format("06-01")}
// 		if err := db.CreateComment(ctx, cmt); err != nil {
// 			return packErr1(err), nil
// 		}
// 		user, err := getUser(ctx, cmt.UserId, req.Token)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return &interact.CommentActionResponse{StatusCode: errno.SuccessCode,
// 			Comment: &interact.Comment{Id: int64(cmt.ID), User: user, Content: cmt.Content, CreateDate: cmt.CreateDate}}, nil
// 	} else if req.ActionType == 2 {
// 		cmt := &db.Comment{ID: int64(req.CommentId)}
// 		tmp, err := db.SelectComment(ctx, int64(cmt.ID))
// 		if err != nil {
// 			return packErr1(err), nil
// 		}
// 		if tmp == nil {
// 			return &interact.CommentActionResponse{StatusCode: errno.CommentNotFound}, nil
// 		}
// 		if err := db.DeleteComment(ctx, cmt); err != nil {
// 			return packErr1(err), nil
// 		}
// 		user, err := getUser(ctx, tmp.UserId, req.Token)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return &interact.CommentActionResponse{StatusCode: errno.SuccessCode,
// 			Comment: &interact.Comment{Id: int64(tmp.ID), User: user, Content: tmp.Content, CreateDate: tmp.CreateDate}}, nil
// 	} else {
// 		msg := "err"
// 		return &interact.CommentActionResponse{StatusCode: errno.ActionTypeErrCode, StatusMsg: msg}, nil
// 	}
// }

// // CommentList implements the CommentSrvImpl interface.
// func (s *InteractServiceImpl) CommentList(ctx context.Context, req *interact.CommentListRequest) (resp *interact.CommentListResponse, err error) {
// 	log.Info("get interact list req", *req)
// 	// TODO: Your code here...
// 	resp = new(interact.CommentListResponse)
// 	cmts, err := db.QueryComments(ctx, int64(req.VideoId))
// 	if err != nil {
// 		return packErr2(err), nil
// 	}
// 	res := []*interact.Comment{}
// 	for _, c := range cmts {
// 		fmt.Println("get user , id = ", c.UserId)
// 		user, err := getUser(ctx, c.UserId, req.Token)
// 		if err != nil {
// 			return nil, err
// 		}
// 		tmp := &interact.Comment{Id: int64(c.ID), Content: c.Content,
// 			CreateDate: c.CreateDate, User: user}
// 		res = append(res, tmp)
// 	}
// 	return &interact.CommentListResponse{StatusCode: errno.SuccessCode, CommentList: res}, nil
// }
