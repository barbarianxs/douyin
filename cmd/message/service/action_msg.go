package service

import (
	"context"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/message/dal/db"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/message/pack"
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
	MessageModel := &db.Message{
		ToUserId:   req.ToUserId,
		FromUserId:  req.FromUserId,
		Content: req.Content,
	}
	return db.CreateMessage(s.ctx, []*db.Message{MessageModel})
}