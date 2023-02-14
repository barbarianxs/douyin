package service

import (
	"context"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/message/dal"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/message/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/message"
)

type ChatMsgService struct {
	ctx context.Context
}

func NewChatMsgService(ctx context.Context) *ChatMsgService {
	return &ChatMsgService{ctx: ctx}
}

func (s *ChatMsgService) MGetChatMsg(req *message.MessageChatRequest, to_user_id int64) ([]*message.Message, error) {
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
	modelMessages, err := dal.MGetChatMsg(s.ctx, myID, req.ToUserId)
	if err != nil {
		return nil, err
	}
	return pack.Messages(modelMessages), nil
}