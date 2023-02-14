
package db

import (
	"context"
	"fmt"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/consts"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ToUserId   int64  `gorm:"type:varchar(32);not null" json:"to_user_id"`
	FromUserId int64  `gorm:"type:varchar(32);not null" json:"from_user_id"`
	Content    string `gorm:"type:varchar(256);not null" json:"content"`
	// CreatedAt   time.Time             `json:"createAt"`
	message_time
	
}

func (u *Message) TableName() string {
	return consts.MessageTableName
}

// MGetMessages multiple get list of message info
func MGetMessages(ctx context.Context, uid string, toUI string, offset int) ([]*Message, error) {
	res := make([]*Message, 0)
	
	if err := DB.WithContext(ctx).
	Model(&Message{}).
	Select("messages.*, users.username").
	Joins("INNER Join users on users.id = messages.user_id").
	Where("(" +
		"(" + "messages.from_user_id = " + uid + " and messages.from_user_id=" + toUId + ")" +
		" or " +
		"(" + "messages.user_id = " + toUId + " and messages.to_user_id=" + uid + ")" +
		")").
	Order("messages.id desc").
	// Offset(offset).
	Limit(100).
	Scan(&res).Error; err != nil{
		return nil, err
	}

	// if offset == 0{
	// 	sort.Slice(res, func(i, j int) bool {
	// 		return res[i]["id"].(uint32) < res[j]["id"].(uint32)
	// 	})
	// }
	return res, nil
}

// CreateMessage create message info
func CreateMessage(ctx context.Context, messages []*Message) error {
	
	if err := DB.WithContext(ctx).Create(messages).Error; err != nil {
		return err
	}
	return nil
}

// // QueryMessage query list of message info
// func QueryMessage(ctx context.Context, to_user_id string) ([]*Message, error) {
// 	res := make([]*Message, 0)
// 	if err := DB.WithContext(ctx).Where("ToUserId = ?", to_user_id).Find(&res).Error; err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }
