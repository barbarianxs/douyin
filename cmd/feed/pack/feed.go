package pack

import (
	"errors"

	"github.com/YANGJUNYAN0715/douyin/tree/zhao/kitex_gen/feed"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/pkg/errno"
)

func BuildVideoResp(err error) *feed.DouyinFeedResponse {
	if err == nil {
		//return videoResp(errno.Success)
		return videoResp()
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		//return videoResp(e)
		return videoResp()
	}

	//s := errno.ErrUnknown.WithMessage(err.Error())
	//return videoResp(s)
	return videoResp()
}

func videoResp() *feed.DouyinFeedResponse {
	return &feed.DouyinFeedResponse{StatusCode: int32(0), StatusMsg: "123"}
}