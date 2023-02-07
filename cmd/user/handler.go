package main

import (
	"context"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/user/pack"
	userservice "github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/user/service"
	user "github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user/userservice"
	
	
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// LoginUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) LoginUser(ctx context.Context, req *user.LoginUserRequest) (resp *user.LoginUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.LoginUserResponse)

	if err = req.IsValid(); err != nil {
		
		return resp, nil
	}
	
	err = service.NewCreateUserService(ctx).LoginUserService(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	return
}

// LogoutUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) LogoutUser(ctx context.Context, req *user.LogoutUserRequest) (resp *user.LogoutUserResponse, err error) {
	// TODO: Your code here...
	return
}

// RegisterUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) RegisterUser(ctx context.Context, req *user.RegisterUserRequest) (resp *user.RegisterUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.RegisterUserResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	err = service.NewCreateUserService(ctx).RegisterUserService(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}
