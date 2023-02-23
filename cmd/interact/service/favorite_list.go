package service

import (
	"context"
	"errors"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/interact/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/interact/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/interact"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/consts"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/jwt"
	"sync"
	"log"
)

type FavoriteListService struct {
	ctx context.Context
}

// NewFavoriteListService new FavoriteListService
func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{ctx: ctx}
}

// FavoriteList get video information that users guoke
func (s *FavoriteListService) FavoriteList(req *interact.FavoriteListRequest) ([]*interact.Video, error) {
	//获取用户id
	// Jwt := jwt.NewJWT([]byte(consts.SecretKey))
	// req.UserId, _ := Jwt.CheckToken(req.Token)
	log.Println("1===============================",req.UserId,"==================================")
	//检查用户是否存在
	user, err := db.QueryUserByIds(s.ctx, []int64{req.UserId})
	if err != nil {
		return nil, err
	}
	if len(user) == 0 {
		return nil, errors.New("user not exist")
	}

	//获取目标用户的点赞视频id号
	videoIds, err := db.QueryFavoriteById(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	log.Println("2===============================",videoIds[0],"==================================")
	//获取点赞视频的信息
	videoData, err := db.QueryVideoByVideoIds(s.ctx, videoIds)
	if err != nil {
		return nil, err
	}
	log.Println("3===============================",videoData,"==================================")
	//获取点赞视频的用户id号
	userIds := make([]int64, 0)
	for _, video := range videoData {
		userIds = append(userIds, video.AuthorID)
	}
	log.Println("4===============================",userIds,"==================================")
	//获取点赞视频的用户信息
	users, err := db.QueryUserByIds(s.ctx, userIds)
	if err != nil {
		return nil, err
	}
	userMap := make(map[int64]*db.User)
	for _, user := range users {
		userMap[int64(user.ID)] = user
	}

	var interactMap map[int64]*db.Favorite
	var relationMap map[int64]*db.Relation
	//if user not logged in
	if req.UserId == -1 {
		interactMap = nil
		relationMap = nil
	} else {
		var wg sync.WaitGroup
		wg.Add(2)
		var interactErr, relationErr error
		//获取点赞信息
		go func() {
			defer wg.Done()
			interactMap, err = db.QueryFavoriteByIds(s.ctx, req.UserId, videoIds)
			if err != nil {
				interactErr = err
				return
			}
		}()
		//获取关注信息
		go func() {
			defer wg.Done()
			relationMap, err = db.QueryRelationByIds(s.ctx, req.UserId, userIds)
			if err != nil {
				relationErr = err
				return
			}
		}()
		wg.Wait()
		if interactErr != nil {
			return nil, interactErr
		}
		if relationErr != nil {
			return nil, relationErr
		}

	}
	log.Println("4===============================",userIds,"==================================")
	videoList := pack.VideoList(req.UserId, videoData, userMap, interactMap, relationMap)
	return videoList, nil

}