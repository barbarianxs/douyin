package service

import (
	"context"
<<<<<<< HEAD
<<<<<<< HEAD
	"log"
	// "time"
=======

>>>>>>> origin/guo
=======
	"log"
	// "time"
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/relation/dal/db"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/message/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/relation"
)
type ActionMsgService struct {
	ctx context.Context
}

func NewActionMsgService(ctx context.Context) *ActionMsgService {
	return &ActionMsgService{ctx: ctx}
}

// Create Message
<<<<<<< HEAD
<<<<<<< HEAD
func (s *ActionMsgService) ActionMsg(req *relation.MessageActionRequest) error {
=======
func (s *ActionMsgService) MGetActionMsg(req *relation.MessageActionRequest) error {
>>>>>>> origin/guo
=======
func (s *ActionMsgService) ActionMsg(req *relation.MessageActionRequest) error {
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	MessageModel := &db.Message{
		ToUserId:   req.ToUserId,
		FromUserId:  req.FromUserId,
		Content: req.Content,
<<<<<<< HEAD
<<<<<<< HEAD

	}
	
	
	log.Println(req.FromUserId, "------------------------", req.ToUserId)
	return db.CreateMessage(s.ctx, MessageModel)
=======
	}
	return db.CreateMessage(s.ctx, []*db.Message{MessageModel})
>>>>>>> origin/guo
=======

	}
	
	
	log.Println(req.FromUserId, "------------------------", req.ToUserId)
	return db.CreateMessage(s.ctx, MessageModel)
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
}