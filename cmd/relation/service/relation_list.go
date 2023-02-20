
package service

import (
	"context"
	// "crypto/md5"
	// "fmt"
	// "io"
	// "log"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/relation/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/relation"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
)

type RelationListService struct {
	ctx context.Context
}

// NewRelationListService new RelationListService
func NewRelationListService(ctx context.Context) *RelationListService {
	return &RelationListService{ctx: ctx}
}

// 查找关注列表
func (s *RelationListService) RelationFollowList(req *relation.RelationFollowListRequest) ([]*relation.User, error) {
	users, err := db.RelationFollowList(s.ctx, req.UserId)
	if err != nil{
		return nil,err
	}
	// log.Println("relation-service")
	// log.Println(users)
	return users,nil
}

// 查找粉丝列表 
func (s *RelationListService) RelationFollowerList(req *relation.RelationFollowerListRequest)  ([]*relation.User, error) {
	users, err := db.RelationFollowerList(s.ctx, req.UserId)
	if err != nil{
		return nil,err
	}
	return users,nil
}

// 查找好友列表  💦先用粉丝列表代替，返回为user包装得到的FriendUser
func (s *RelationListService) RelationFriendList(req *relation.RelationFriendListRequest)  ([]*relation.FriendUser, error) {
	users, err := db.RelationFriendList(s.ctx, req.UserId)
	if err != nil{
		return nil,err
	}
	return users,nil
}
