package db

import (
	"context"
	"github.com/YANGJUNYAN0715/douyin/tree/li/pkg/consts"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"index:idx_username,unique;type:varchar(40);not null" json:"username"`
	Password string `gorm:"type:varchar(256);not null" json:"password"`
	// FavoriteVideos []Video `gorm:"many2many:user_favorite_videos" json:"favorite_videos"`
	FollowingCount int `gorm:"default:0" json:"following_count"`
	FollowerCount  int `gorm:"default:0" json:"follower_count"`
}

func (u *User) TableName() string {
	return consts.UserTableName
}

// GetUserById multiple get list of user info
func GetUserById(ctx context.Context, userID int64) (*User, error) {
	var res *User
	if userID == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id = ?", userID).Find(&res).Error; err != nil {
		return nil, err
	}
	log.Println("***db***")
	log.Println(res)
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
