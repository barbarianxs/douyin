package db

import (
	"context"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/pkg/consts"
)

type Comment struct {
	ID         uint `gorm:"primarykey"`
	UserId     int
	VideoId    int
	ParentId   int
	IsValid    bool
	Content    string
	CreateTime string
}

func (u *Comment) TableName() string {
	return consts.CommentTableName
}

// CreateUser create user info
func CreateComment(ctx context.Context, arr *Comment) error {
	return DB.WithContext(ctx).Create(arr).Error
}

// CreateUser create user info
func SelectComment(ctx context.Context, id int) (*Comment, error) {
	cmt := &Comment{ID: uint(id)}
	return cmt, DB.Model(&Comment{}).WithContext(ctx).First(cmt).Error
}

// CreateUser create user info
func DeleteComment(ctx context.Context, arr *Comment) error {
	return DB.Model(&Comment{}).WithContext(ctx).Where("id = ?", arr.ID).Update("is_valid", false).Error
}

func deleteAll(ctx context.Context, videoId int) error {
	return DB.Model(&Comment{}).WithContext(ctx).Where("video_id = ?", videoId).
		Delete(&Comment{}).Error
}

// QueryUser query list of user info
func QueryComments(ctx context.Context, videoId int) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := DB.WithContext(ctx).Order("create_time desc").
		Where("video_id = ? and is_valid = 1", videoId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
