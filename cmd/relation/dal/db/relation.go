
package db

import (
	"context"
	// "fmt"
	"time"
<<<<<<< HEAD
<<<<<<< HEAD
	"log"
=======
>>>>>>> origin/guo
=======
	"log"
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/consts"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/user/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/relation"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
	"gorm.io/gorm"
)
// type User struct {
// 	gorm.Model
// 	ID            int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// 	Name string `json:"name"`
// 	FollowCount   int64  `json:"follow_count"`
// 	FollowerCount int64  `json:"follower_count"`
// 	IsFollow      bool   `json:"is_follow"`
// 	// Avatar string  `json:"avatar"`
// 	// BackgroundImage string  `json:"background_image"`
// 	// Signature string  `json:"signature"`
// 	// TotalFavorited string  `json:"total_favorited"`
// 	// WorkCount int64  `json:"work_count"`
// 	// FavoriteCount int64  `json:"favorite_count"`
// }

type Message struct {
	gorm.Model
	ToUserId   int64  `gorm:"type:varchar(32);not null" json:"to_user_id"`
	FromUserId int64  `gorm:"type:varchar(32);not null" json:"from_user_id"`
	Content    string `gorm:"type:varchar(256);not null" json:"content"`
<<<<<<< HEAD
<<<<<<< HEAD
	CreateTime   time.Time   `gorm:"column:create_time;default:null " json:"create_time"`
=======
	// CreatedAt   time.Time             `json:"createAt"`
>>>>>>> origin/guo
=======
	CreateTime   time.Time   `gorm:"column:create_time;default:null " json:"create_time"`
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	
	
}

// Relation Gorm data structure
type User struct {
	gorm.Model
	ID            int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Username       string  `gorm:"index:idx_username,unique;type:varchar(40);not null" json:"username"`
	Password       string  `gorm:"type:varchar(256);not null" json:"password"`
	// FavoriteVideos []Video `gorm:"many2many:user_favorite_videos" json:"favorite_videos"`
	FollowingCount int64     `gorm:"default:0" json:"following_count"`
	FollowerCount  int64     `gorm:"default:0" json:"follower_count"`
}

<<<<<<< HEAD
<<<<<<< HEAD
type Relation struct {
	gorm.Model
	ID         int64     `gorm:"column:id;primary_key;AUTO_INCERMENT"`
	// FollowTime time.Time `gorm:"column:follow_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	FromUserID int64     `gorm:"column:from_user_id;NOT NULL"`
	ToUserID   int64     `gorm:"column:to_user_id;NOT NULL"`
	// CreateTime int64  `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	// UpdateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"`
}

// // Relation表 记录关注关系
// // 不设置外键 提高效率 通过程序保证参照完整性
// type Relation struct {
// 	gorm.Model
// 	UserID   int64  `gorm:"index:idx_userid;not null"`
// 	ToUserID int64  `gorm:"index:index:idx_userid_to;not null"`
// }
=======
type Follow struct {
=======
type Relation struct {
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	gorm.Model
	ID         int64     `gorm:"column:id;primary_key;AUTO_INCERMENT"`
	// FollowTime time.Time `gorm:"column:follow_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	FromUserID int64     `gorm:"column:from_user_id;NOT NULL"`
	ToUserID   int64     `gorm:"column:to_user_id;NOT NULL"`
	// CreateTime int64  `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	// UpdateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"`
}

<<<<<<< HEAD
// Relation表 记录关注关系
// 不设置外键 提高效率 通过程序保证参照完整性
type Relation struct {
	gorm.Model
	UserID   int64  `gorm:"index:idx_userid;not null"`
	ToUserID int64  `gorm:"index:index:idx_userid_to;not null"`
}
>>>>>>> origin/guo
=======
// // Relation表 记录关注关系
// // 不设置外键 提高效率 通过程序保证参照完整性
// type Relation struct {
// 	gorm.Model
// 	UserID   int64  `gorm:"index:idx_userid;not null"`
// 	ToUserID int64  `gorm:"index:index:idx_userid_to;not null"`
// }
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a


func (u *Relation) TableName() string {
	return consts.RelationTableName
}

func (u *Message) TableName() string {
	return consts.MessageTableName

}

// GetRelation get relation info
func GetRelation(ctx context.Context, uid int64, tid int64) (*Relation, error) {
	relations := new(Relation)

<<<<<<< HEAD
<<<<<<< HEAD
	if err := DB.WithContext(ctx).First(&relations, "from_user_id = ? and to_user_id = ?", uid, tid).Error; err != nil {
=======
	if err := DB.WithContext(ctx).First(&relations, "user_id = ? and to_user_id = ?", uid, tid).Error; err != nil {
>>>>>>> origin/guo
=======
	if err := DB.WithContext(ctx).First(&relations, "from_user_id = ? and to_user_id = ?", uid, tid).Error; err != nil {
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
		return nil, err
	}
	return relations, nil
}
//根据id获取user
// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}
	// 从usr表中根据id查找到users的信息
	if err := DB.Table(consts.UserTableName).WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
// NewAction creates a new Relation
// uid关注tid，所以uid的关注人数加一，tid的粉丝数加一
func NewAction(ctx context.Context, uid int64, tid int64) error {
	relations,_ :=GetRelation(ctx,uid,tid)
		if relations != nil{
			return errno.RelationExistErr
		}

	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作
		// 1. 新增关注数据
<<<<<<< HEAD
<<<<<<< HEAD
		err := tx.Create(&Relation{FromUserID: uid, ToUserID: tid}).Error
=======
		err := tx.Create(&Relation{UserID: uid, ToUserID: tid}).Error
>>>>>>> origin/guo
=======
		err := tx.Create(&Relation{FromUserID: uid, ToUserID: tid}).Error
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
		if err != nil {
			return err
		}

		// 2.改变 user 表中的 following count
<<<<<<< HEAD
<<<<<<< HEAD
		res := tx.Table(consts.UserTableName).Where("id = ?", uid).Update("follow_count", gorm.Expr("follow_count + ?", 1))
=======
		res := tx.Table(consts.UserTableName).Where("ID = ?", uid).Update("following_count", gorm.Expr("following_count + ?", 1))
>>>>>>> origin/guo
=======
		res := tx.Table(consts.UserTableName).Where("id = ?", uid).Update("follow_count", gorm.Expr("follow_count + ?", 1))
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.RelationActionErr
		}

		// 3.改变 user 表中的 follower count
<<<<<<< HEAD
<<<<<<< HEAD
		res = tx.Table(consts.UserTableName).Where("id = ?", tid).Update("follower_count", gorm.Expr("follower_count + ?", 1))
=======
		res = tx.Table(consts.UserTableName).Where("ID = ?", tid).Update("follower_count", gorm.Expr("follower_count + ?", 1))
>>>>>>> origin/guo
=======
		res = tx.Table(consts.UserTableName).Where("id = ?", tid).Update("follower_count", gorm.Expr("follower_count + ?", 1))
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.RelationActionErr
		}

		return nil
	})
	return err
}

// DelAction deletes a relation from the database.
func DelAction(ctx context.Context, uid int64, tid int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作
		relations := new(Relation)
<<<<<<< HEAD
<<<<<<< HEAD
		if err := tx.Where("from_user_id = ? AND to_user_id=?", uid, tid).First(&relations).Error; err != nil {
=======
		if err := tx.Where("user_id = ? AND to_user_id=?", uid, tid).First(&relations).Error; err != nil {
>>>>>>> origin/guo
=======
		if err := tx.Where("from_user_id = ? AND to_user_id=?", uid, tid).First(&relations).Error; err != nil {
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
			return err
		}

		// 1. 删除关注数据
		err := tx.Unscoped().Delete(&relations).Error
		if err != nil {
			return err
		}
		// 2.改变 user 表中的 following count
<<<<<<< HEAD
<<<<<<< HEAD
		res := tx.Table(consts.UserTableName).Where("id = ?", uid).Update("follow_count", gorm.Expr("follow_count - ?", 1))
=======
		res := tx.Table(consts.UserTableName).Where("ID = ?", uid).Update("following_count", gorm.Expr("following_count - ?", 1))
>>>>>>> origin/guo
=======
		res := tx.Table(consts.UserTableName).Where("id = ?", uid).Update("follow_count", gorm.Expr("follow_count - ?", 1))
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.RelationActionErr
		}

		// 3.改变 user 表中的 follower count
<<<<<<< HEAD
<<<<<<< HEAD
		res = tx.Table(consts.UserTableName).Where("id = ?", tid).Update("follower_count", gorm.Expr("follower_count - ?", 1))
=======
		res = tx.Table(consts.UserTableName).Where("ID = ?", tid).Update("follower_count", gorm.Expr("follower_count - ?", 1))
>>>>>>> origin/guo
=======
		res = tx.Table(consts.UserTableName).Where("id = ?", tid).Update("follower_count", gorm.Expr("follower_count - ?", 1))
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.RelationActionErr
		}

		return nil
	})
	return err
}

// RelationFollowList returns the Following List.
func RelationFollowList(ctx context.Context, uid int64) ([]*relation.User, error) {
	var RelationList []*Relation
<<<<<<< HEAD
<<<<<<< HEAD
	err := DB.WithContext(ctx).Where("from_user_id = ?", uid).Find(&RelationList).Error
=======
	err := DB.WithContext(ctx).Where("user_id = ?", uid).Find(&RelationList).Error
>>>>>>> origin/guo
=======
	err := DB.WithContext(ctx).Where("from_user_id = ?", uid).Find(&RelationList).Error
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	if err != nil {
		return nil, err
	}
	userIDs :=make([]int64,0)
	for _,u := range RelationList{
		userIDs= append(userIDs,int64(u.ToUserID))
	}
	users, err := MGetUsers(ctx,userIDs)
	if err != nil {
		return nil, err
	}
	// log.Println(users)
	return BuildUsers(ctx,uid,users)
}

// RelationFollowerList returns the Follower List.
func RelationFollowerList(ctx context.Context, tid int64) ([]*relation.User, error) {
	var RelationList []*Relation
	err := DB.WithContext(ctx).Where("to_user_id = ?", tid).Find(&RelationList).Error
	if err != nil {
		return nil, err
	}
	userIDs :=make([]int64,0)
	for _,u := range RelationList{
<<<<<<< HEAD
<<<<<<< HEAD
		userIDs= append(userIDs,int64(u.FromUserID))
=======
		userIDs= append(userIDs,int64(u.UserID))
>>>>>>> origin/guo
=======
		userIDs= append(userIDs,int64(u.FromUserID))
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	}
	users, err := MGetUsers(ctx,userIDs)
	if err != nil {
		return nil, err
	}
	return BuildUsers(ctx,tid,users)
}

// 朋友：相互关注者->粉丝和关注者的交集
// RelationFriendList returns the Follower List.
func RelationFriendList(ctx context.Context, id int64) ([]*relation.FriendUser, error) {
	var LRelationList []*Relation //关注者
	var RRelationList []*Relation //粉丝
<<<<<<< HEAD
<<<<<<< HEAD
	err := DB.WithContext(ctx).Where("from_user_id = ?", id).Find(&LRelationList).Error
=======
	err := DB.WithContext(ctx).Where("user_id = ?", id).Find(&LRelationList).Error
>>>>>>> origin/guo
=======
	err := DB.WithContext(ctx).Where("from_user_id = ?", id).Find(&LRelationList).Error
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	if err != nil {
		return nil, err
	}
	err = DB.WithContext(ctx).Where("to_user_id = ?", id).Find(&RRelationList).Error
	if err != nil {
		return nil, err
	}

	LuserIDs :=make([]int64,0)
	for _,u := range LRelationList{
		LuserIDs= append(LuserIDs,int64(u.ToUserID))
<<<<<<< HEAD
<<<<<<< HEAD
		log.Println(LuserIDs)
=======
>>>>>>> origin/guo
=======
		log.Println(LuserIDs)
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	}

	RuserIDs :=make([]int64,0)
	for _,u := range RRelationList{
<<<<<<< HEAD
<<<<<<< HEAD
		RuserIDs= append(RuserIDs,int64(u.FromUserID))
		log.Println(RuserIDs)
=======
		RuserIDs= append(RuserIDs,int64(u.UserID))
>>>>>>> origin/guo
=======
		RuserIDs= append(RuserIDs,int64(u.FromUserID))
		log.Println(RuserIDs)
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	}
	userIDs :=make([]int64,0)

	m := make(map[int64]int)
	for _,v :=range LuserIDs{
		m[v]++
	}
	for _,v :=range RuserIDs{
		if m[v]==1{
			userIDs = append(userIDs,v)
		}
	}
<<<<<<< HEAD
<<<<<<< HEAD
	log.Println(userIDs)
=======
	// log.Println(userIDs)
>>>>>>> origin/guo
=======
	log.Println(userIDs)
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	users, err := MGetUsers(ctx,userIDs)
	if err != nil {
		return nil, err
	}
	return BuildFriendUsers(ctx,id,users)
}


// MGetMessages multiple get list of message info
func MGetMessages(ctx context.Context, uid int64, toUId int64) ([]*Message, error) {
	res := make([]*Message, 0)
	
<<<<<<< HEAD
<<<<<<< HEAD
	if err := DB.WithContext(ctx).Model(&Message{}).Where("from_user_id = ? AND to_user_id = ? Or from_user_id = ? AND to_user_id = ?", uid, toUId, toUId, uid).Order("id desc").Scan(&res).Error; err != nil{
		return nil, err
	}
	
	// log.Println(":::::::::::::::::::::::::::::::", res)
=======
	if err := DB.WithContext(ctx).Model(&Message{}).Where("from_user_id = ?", uid).Where("to_user_id = ?", toUId).Order("id desc").Scan(&res).Error; err != nil{
		return nil, err
	}

>>>>>>> origin/guo
=======
	if err := DB.WithContext(ctx).Model(&Message{}).Where("from_user_id = ? AND to_user_id = ? Or from_user_id = ? AND to_user_id = ?", uid, toUId, toUId, uid).Order("id desc").Scan(&res).Error; err != nil{
		return nil, err
	}
	
	// log.Println(":::::::::::::::::::::::::::::::", res)
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	// if offset == 0{
	// 	sort.Slice(res, func(i, j int) bool {
	// 		return res[i]["id"].(uint32) < res[j]["id"].(uint32)
	// 	})
	// }
	return res, nil
}

// CreateMessage create message info
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
func CreateMessage(ctx context.Context, message *Message) error {
	log.Println(message)
	if err := DB.WithContext(ctx).Create(message).Error; err != nil {
		log.Println(err)
<<<<<<< HEAD
		return err
	}
	log.Println("++++++++++++++++++++++++++++++",message)
=======
func CreateMessage(ctx context.Context, messages []*Message) error {
	
	if err := DB.WithContext(ctx).Create(messages).Error; err != nil {
		return err
	}
>>>>>>> origin/guo
=======
		return err
	}
	log.Println("++++++++++++++++++++++++++++++",message)
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
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
