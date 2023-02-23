package db

import (
	"context"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)
type Video struct {
	gorm.Model
	ID       int64   `gorm:"column:id;primary_key;AUTO_INCREMENT"`   
	AuthorID      int64     `gorm:"column:author_id;NOT NULL"`
	// PublishTime   time.Time `gorm:"column:publish_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	PlayUrl      string    `gorm:"column:play_url;NOT NULL"`
	CoverUrl     string    `gorm:"column:cover_url;NOT NULL"`
	FavoriteCount int64     `gorm:"column:favorite_count;default:0"`
	CommentCount  int64     `gorm:"column:comment_count;default:0"`
	Title         string    `gorm:"column:title;NOT NULL"`
	IsFavorite bool  `gorm:"column:is_favorite;default:0"`
	UpdatedAt   time.Time   `gorm:"column:updated_at;default:null " json:"updated_at"`
}

func (v *Video) TableName() string {
	return consts.VideoTableName
}

// QueryVideoByLatestTime query video info by latest create Time
func QueryVideoByLatestTime(ctx context.Context, latestTime int64) ([]*Video, error) {
	var videos []*Video
	time := time.UnixMilli(latestTime)
	err := DB.WithContext(ctx).Limit(30).Order("update_time desc").Where("update_time < ?", time).Find(&videos).Error
	if err != nil {
		klog.Error("QueryVideoByLatestTime find video error " + err.Error())
		return videos, err
	}
	return videos, nil
}

// QueryVideoByVideoIds query video info by video ids
func QueryVideoByVideoIds(ctx context.Context, videoId []int64) ([]*Video, error) {
	var videos []*Video
	err := DB.WithContext(ctx).Where("id in (?)", videoId).Find(&videos).Error
	if err != nil {
		klog.Error("QueryVideoByVideoIds error " + err.Error())
		return nil, err
	}
	return videos, nil
}
