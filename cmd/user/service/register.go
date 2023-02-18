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

type RegisterService struct {
	ctx context.Context
}

// NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Register create user info.
func (s *RegisterService) Register(req *user.DouyinUserRegisterRequest) error {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	return db.CreateUser(s.ctx, []*db.User{{
		Username: req.Username,
		Password: password,
	}})
}
