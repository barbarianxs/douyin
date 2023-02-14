package main

import (
	"context"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/message/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/message/service"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/message"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageChatMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChatMessage(ctx context.Context, req *message.MessageChatRequest) (resp *message.MessageChatResponse, err error) {
	// TODO: Your code here...
	resp = new(message.MessageActionResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewChatMsgService(ctx).ChatMsgService(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.UserId = uid
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// MessageActionMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageActionMessage(ctx context.Context, req *message.MessageActionRequest) (resp *message.MessageActionResponse, err error) {
	// TODO: Your code here...
	resp = new(message.MessageActionResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewActionMsgService(ctx).MGetActionMsg(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
	
}
