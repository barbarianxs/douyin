package service

import (
	"context"
	"errors"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/interact/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/interact"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/consts"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/jwt"
)

type FavoriteActionService struct {
	ctx context.Context
}

// NewFavoriteActionService new FavoriteActionService
func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

// FavoriteAction implement the guoke and unguoke operations
func (s *FavoriteActionService) FavoriteAction(req *interact.FavoriteActionRequest) error {
	// Jwt := jwt.NewJWT([]byte(consts.SecretKey))
	// claim, err := Jwt.ParseToken(req.Token)
	// if err != nil {
	// 	return err
	// }
	// req.UserId := claim.Id

	videos, err := db.QueryVideoByVideoIds(s.ctx, []int64{req.VideoId})
	if err != nil {
		return err
	}
	if len(videos) == 0 {
		return errors.New("video not exist")
	}

	//若ActionType（操作类型）等于1，则向interact表创建一条记录，同时向video表的目标video增加点赞数
	//若ActionType等于2，则向interact表删除一条记录，同时向video表的目标video减少点赞数
	//若ActionType不等于1和2，则返回错误
	if req.ActionType == consts.Like {
		interact := &db.Favorite{
			UserId:  req.UserId,
			VideoId: req.VideoId,
		}

		err := db.CreateFavorite(s.ctx, interact, req.VideoId)
		if err != nil {
			return err
		}
	}
	if req.ActionType == consts.Unlike {
		err := db.DeleteFavorite(s.ctx, req.UserId, req.VideoId)
		if err != nil {
			return err
		}

	}
	if req.ActionType != consts.Like && req.ActionType != consts.Unlike {
		return errors.New("action type no equal 1 and 2")
	}
	return nil
}