package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/YANGJUNYAN0715/douyin/tree/li/cmd/user/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/li/kitex_gen/user"
	"github.com/YANGJUNYAN0715/douyin/tree/li/pkg/errno"
)

type LoginService struct {
	ctx context.Context
}

// NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{
		ctx: ctx,
	}
}

// Login check user info
func (s *LoginService) Login(req *user.DouyinUserRegisterRequest) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.Username
	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.AuthorizationFailedErr
	}
	u := users[0]
	if u.Password != passWord {
		return 0, errno.AuthorizationFailedErr
	}
	return int64(u.ID), nil
}
