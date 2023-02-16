package service

import (
	"context"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/message/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/message/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/message"
)

type ChatMsgService struct {
	ctx context.Context
}

func NewChatMsgService(ctx context.Context) *ChatMsgService {
	return &ChatMsgService{ctx: ctx}
}

func (s *ChatMsgService) MGetChatMsg(req *message.MessageChatRequest) ([]*message.Message, error) {
	messageModels, err := db.MGetMessages(s.ctx, req.FromUserId, req.ToUserId)
	if err != nil {
		return nil, err
	}
	
	messages := pack.Messages(messageModels)
	
	return messages, nil
	
}