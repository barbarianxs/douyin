package service

import (
	"context"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/message/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/message/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/message"
)
type ActionMsgService struct {
	ctx context.Context
}

func NewActionMsgService(ctx context.Context) *ActionMsgService {
	return &ActionMsgService{ctx: ctx}
}
// Create Message
func (s *ActionMsgService) MGetActionMsg(req *message.MessageActionRequest, to_user_id int64) error {
	// 用户鉴权
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = 1001
		return &resp, nil
	}
	// 检查登录状态
	myID := claims.ID
	if myID == 0 {
		resp.StatusCode = 1002
		return &resp, nil
	}
	MessageModel := &db.message{
		ToUserId:  req.ToUserId,
		FromUserId:   myID,

		// ActionType: req.ActionType
		Content: req.Content,
	}
	return db.CreateMessage(s.ctx, []*db.Message{MessageModel})
}