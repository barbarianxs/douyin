package db

import (
	"context"
	"time"

	// "gorm.io/gorm"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user"
)

// Video Gorm Data Structures
// Video 属于 Author, AuthorID是外键(belongs to)
// type Video struct {
// 	gorm.Model
// 	UpdatedAt     time.Time `gorm:"column:update_time;not null;index:idx_update" `
// 	Author        user.User      `gorm:"foreignkey:AuthorID"`
// 	AuthorID      int       `gorm:"index:idx_authorid;not null"`
// 	PlayUrl       string    `gorm:"type:varchar(255);not null"`
// 	CoverUrl      string    `gorm:"type:varchar(255)"`
// 	FavoriteCount int       `gorm:"default:0"`
// 	CommentCount  int       `gorm:"default:0"`
// 	Title         string    `gorm:"type:varchar(50);not null"`
// }

// func (Video) TableName() string {
// 	return "video"
// }

// MGetVideoss multiple get list of videos info
func MGetVideos(ctx context.Context, limit int, latestTime int64) ([]*Video, error) {
	videos := make([]*Video, 0)

	if latestTime == 0 {
		cur_time := int64(time.Now().UnixMilli())
		latestTime = cur_time
	}
	conn := DB.WithContext(ctx)

	if err := conn.Limit(limit).Order("update_at desc").Find(&videos, "update_at < ?", time.UnixMilli(latestTime)).Error; err != nil {
		return nil, err
	}
	return videos, nil
}