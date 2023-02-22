package service

import (
	"context"
	"log"
	// "time"
	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/relation/dal/db"
	// "github.com/YANGJUNYAN0715/douyin/tree/main/cmd/message/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/relation"
)
type ActionMsgService struct {
	ctx context.Context
}

func NewActionMsgService(ctx context.Context) *ActionMsgService {
	return &ActionMsgService{ctx: ctx}
}

// Create Message
func (s *ActionMsgService) ActionMsg(req *relation.MessageActionRequest) error {
	MessageModel := &db.Message{
		ToUserId:   req.ToUserId,
		FromUserId:  req.FromUserId,
		Content: req.Content,

	}
	
	
	log.Println(req.FromUserId, "------------------------", req.ToUserId)
	return db.CreateMessage(s.ctx, MessageModel)
}