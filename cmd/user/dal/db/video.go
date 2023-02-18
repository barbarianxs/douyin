
package db

import (
	"context"
	// "fmt"
	"time"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/consts"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	ID            int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	AuthorID      int64     `gorm:"column:author_id;NOT NULL"`
	PublishTime   time.Time `gorm:"column:publish_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	FilePath      string    `gorm:"column:file_path;NOT NULL"`
	CoverPath     string    `gorm:"column:cover_path;NOT NULL"`
	FavoriteCount int64     `gorm:"column:favorite_count;default:0"`
	CommentCount  int64     `gorm:"column:comment_count;default:0"`
	Title         string    `gorm:"column:title;NOT NULL"`
}

func (v *Video) TableName() string {
	return consts.VideoTableName
}

// func UserGetFeed(ctx context.Context, latestTime *int64) ([]*Video, error) {
// 	return
// }

func UserPublishList(ctx context.Context, userID int64) ([]*Video, error) {
	// 获取视频列表
	var videoList []*Video
	conn := DB.WithContext(ctx).Model(&Video{}).Where("author_id = ?", userID).Find(&videoList)
	if err := conn.Error; err != nil {
		return nil, err
	}

	// 返回
	return videoList, nil
}

func CreateVideo(ctx context.Context, videos []*Video) error {
	if err := DB.WithContext(ctx).Create(videos).Error; err != nil {
		return err
	}
	return nil
}