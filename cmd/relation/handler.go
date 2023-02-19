package main

import (
	"context"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/relation/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/relation/service"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/relation"


	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *relation.MessageChatRequest) (resp *relation.MessageChatResponse, err error) {
	// TODO: Your code here...
	return
}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *relation.MessageActionRequest) (resp *relation.MessageActionResponse, err error) {
	// TODO: Your code here...
	return
}
