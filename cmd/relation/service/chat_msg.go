package service

import (
	"context"

	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/relation/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/relation/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/relation"
)

type ChatMsgService struct {
	ctx context.Context
}

func NewChatMsgService(ctx context.Context) *ChatMsgService {
	return &ChatMsgService{ctx: ctx}
}

func (s *ChatMsgService) MGetChatMsg(req *relation.MessageChatRequest) ([]*relation.Message, error) {
	messageModels, err := db.MGetMessages(s.ctx, req.FromUserId, req.ToUserId)
	if err != nil {
		return nil, err
	}
	
	messages := pack.Messages(messageModels)
	
	return messages, nil
	
}