package db

import (
	"context"
	"github.com/YANGJUNYAN0715/douyin/tree/main/pkg/consts"
)

type Comment struct {
	ID         int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	UserId     int64  `gorm:"column:user_id;NOT NULL"`
	VideoId    int64	`gorm:"column:video_id;NOT NULL"`
	ParentId   int64
	IsValid    bool	  `gorm:"column:is_valid;default:1"`
	Content    string	`gorm:"column:content;NOT NULL"`
	CreateDate string  `gorm:"column:create_date;type:varchar(32);not null"`
}

func (u *Comment) TableName() string {
	return consts.CommentTableName
}

// CreateUser create user info
func CreateComment(ctx context.Context, arr *Comment) error {
	return DB.WithContext(ctx).Create(arr).Error
}

// CreateUser create user info
func SelectComment(ctx context.Context, id int64) (*Comment, error) {
	cmt := &Comment{ID: int64(id)}
	return cmt, DB.Model(&Comment{}).WithContext(ctx).First(cmt).Error
}

// CreateUser create user info
func DeleteComment(ctx context.Context, arr *Comment) error {
	return DB.Model(&Comment{}).WithContext(ctx).Where("id = ?", arr.ID).Update("is_valid", false).Error
}

func deleteAll(ctx context.Context, videoId int64) error {
	return DB.Model(&Comment{}).WithContext(ctx).Where("video_id = ?", videoId).
		Delete(&Comment{}).Error
}

// QueryUser query list of user info
func QueryComments(ctx context.Context, videoId int64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := DB.WithContext(ctx).Order("created_at desc").
		Where("video_id = ? and is_valid = 1", videoId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
