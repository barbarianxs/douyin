package main

import (
	"context"
	"fmt"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/user/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/user/service"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// LoginUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) LoginUser(ctx context.Context, req *user.LoginUserRequest) (resp *user.LoginUserResponse, err error) {
	// TODO: Your code here...
	
	resp = new(user.LoginUserResponse)

	if err = req.IsValid(); err != nil {
		resp = pack.BuildLoginResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewLoginUserService(ctx).LoginUser(req)
	if err != nil {
		resp = pack.BuildLoginResp(err)
		return resp, nil
	}

	resp = pack.BuildLoginResp(errno.Success)

	resp.UserId = uid
	
	return resp, nil

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
		resp = pack.BuildRegisterResp(errno.ParamErr)
		return resp, nil
	}
	err = service.NewRegisterUserService(ctx).RegisterUser(req)
	if err != nil {
		resp = pack.BuildRegisterResp(err)
		return resp, nil
	}

	resp = pack.BuildRegisterResp(errno.Success)

	// resp.UserId = uid

	return resp, nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	// TODO: Your code here...

	resp = new(user.UserInfoResponse)

	if err = req.IsValid(); err != nil {
		resp = pack.BuildUserInfoResp(errno.ParamErr)
		return resp, nil
	}

	user_info, err := service.NewUserInfoService(ctx).UserInfo(req)
	if err != nil {
		resp = pack.BuildUserInfoResp(err)
		return resp, nil
	}

	resp = pack.BuildUserInfoResp(errno.Success)

	resp.User = user_info


	return resp, nil
}
