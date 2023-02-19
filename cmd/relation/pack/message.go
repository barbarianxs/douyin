package pack

import (
	// "time"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/relation/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/relation"
)

// Message pack message info
func Message(u *db.Message) *relation.Message {
	if u == nil {
		return nil
	}
	return &relation.Message{
		Id:         int64(u.ID),
		FromUserId: u.FromUserId,
		ToUserId:   u.ToUserId,
		Content:    u.Content,
		// CreateTime: u.CreatedAt.Unix(),
	}
}

// Messages pack list of message info
func Messages(msgs []*db.Message) []*relation.Message {
	messages := make([]*relation.Message, 0)
	for _, msg := range msgs {
		if temp := Message(msg); temp != nil {
			messages = append(messages, temp)
		}
	}
	return messages
}

// func FormatTime(t time.Time) *string {
// 	s := t.Format("2006-01-02 15:04:05")
// 	return &s
// }