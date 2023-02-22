
package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/user/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/user"
	"github.com/YANGJUNYAN0715/douyin/tree/main/pkg/errno"
)

type LoginUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewLoginUserService(ctx context.Context) *LoginUserService {
	return &LoginUserService{ctx: ctx}
}

/// LoginUser Login user info
func (s *LoginUserService) LoginUser(req *user.LoginUserRequest) (int64, error) {
	fmt.Println("---------------------------kitex test----------------------------")
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