package main

import (
	"context"

	"github.com/YANGJUNYAN0715/douyin/tree/li/cmd/user/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/li/cmd/user/service"
	"github.com/YANGJUNYAN0715/douyin/tree/li/kitex_gen/user"
	"github.com/YANGJUNYAN0715/douyin/tree/li/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// 没有用kitex插件生成检查代码，所以没有IsVlid这个函数， 需要自己实现IsValid的判断逻辑，
// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(user.DouyinUserRegisterResponse)
	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = pack.BuildDouyinUserRegisterResponse(errno.ParamErr)
		return resp, nil
	}
	err = service.NewRegisterService(ctx).Register(req)
	if err != nil {
		resp = pack.BuildDouyinUserRegisterResponse(err)
		return resp, nil
	}
	//注册成功，没有设置注册成功后直接登录
	resp = pack.BuildDouyinUserRegisterResponse(errno.Success)
	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(user.DouyinUserRegisterResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = pack.BuildDouyinUserRegisterResponse(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewLoginService(ctx).Login(req)
	if err != nil {
		resp = pack.BuildDouyinUserRegisterResponse(err)
		return resp, nil
	}

	// login success
	resp = pack.BuildDouyinUserRegisterResponse(errno.Success)
	resp.UserId = uid
	return resp, nil
}

// GetUserById implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserById(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.DouyinUserResponse)
	if req.UserId == 0 {
		resp = pack.BuildDouyinUserResponse(errno.ParamErr)
		return resp, nil
	}
	info, err := service.NewGetUserByIdService(ctx).GetUserById(req)
	if err != nil {
		resp = pack.BuildDouyinUserResponse(err)
		return resp, nil
	}

	resp = pack.BuildDouyinUserResponse(errno.Success)
	resp.User = pack.BuildUser(info)
	return resp, nil
	return
}
