
package service

import (
	"context"
	// "crypto/md5"
	// "fmt"
	// "io"

	"github.com/YANGJUNYAN0715/douyin/tree/zhao/cmd/user/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/kitex_gen/user"
	// "github.com/YANGJUNYAN0715/douyin/tree/zhao/pkg/errno"
)

type GetUserByIdService struct {
	ctx context.Context
}

// NewGetUserByIdService new GetUserByIdService
func NewGetUserByIdService(ctx context.Context) *GetUserByIdService {
	return &GetUserByIdService{ctx: ctx}
}

// get user info.
func (s *GetUserByIdService) GetUserById(req *user.DouyinUserRequest) (*db.User,error) {
	info, err := db.GetUserById(s.ctx, req.UserId)
	if err != nil {
		return info,err
	}
	return info,nil
}
