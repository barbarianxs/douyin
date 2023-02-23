package db

import (
	"context"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	// "time"
)
type User struct {
	gorm.Model
	Username string `json:"username"`
	Name          string `gorm:"column:name;index:idx_username,unique;type:varchar(32);not null"`
	Password      string `gorm:"column:password;type:varchar(32);not null"`
	FollowCount   int64  `gorm:"column:follow_count;default:0"`
	FollowerCount int64  `gorm:"column:follower_count;default:0"`
	IsFollow      bool   `json:"is_follow"`
}

func (User) TableName() string {
	return "user"
}

// 根据用户id获取用户信息
func QueryUserByIds(ctx context.Context, userIds []int64) ([]*User, error) {
	var users []*User
	err := DB.WithContext(ctx).Where("id in (?)", userIds).Find(&users).Error
	if err != nil {
		klog.Error("query user by ids fail " + err.Error())
		return nil, err
	}
	return users, nil
}

// 根据用户名获取用户信息
func QueryUserByName(ctx context.Context, userName string) ([]*User, error) {
	var users []*User
	err := DB.WithContext(ctx).Where("name = ?", userName).Find(&users).Error
	if err != nil {
		klog.Error("query user by name fail " + err.Error())
		return nil, err
	}
	return users, nil
}

// 上传用户信息到数据库
func UploadUserData(ctx context.Context, username string, password string) (int64, error) {
	user := &User{
		Name:          username,
		Password:      password,
		FollowCount:   0,
		FollowerCount: 0,
	}
	err := DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		klog.Error("upload user data fail " + err.Error())
		return 0, err
	}
	return int64(user.ID), nil
}

// Relation Gorm Data Structures
type Relation struct {
	gorm.Model
	UserId   int64 `gorm:"column:user_id;not null;index:idx_userid"`
	ToUserId int64 `gorm:"column:to_user_id;not null;index:idx_touserid"`
}

func (Relation) TableName() string {
	return "relation"
}

// 根据当前用户id和目标用户id获取关注信息
func QueryRelationByIds(ctx context.Context, currentId int64, userIds []int64) (map[int64]*Relation, error) {
	var relations []*Relation
	err := DB.WithContext(ctx).Where("from_user_id = ? AND to_user_id IN ?", currentId, userIds).Find(&relations).Error
	if err != nil {
		klog.Error("query relation by ids " + err.Error())
		return nil, err
	}
	relationMap := make(map[int64]*Relation)
	for _, relation := range relations {
		relationMap[relation.ToUserId] = relation
	}
	return relationMap, nil
}

// 增加当前用户的关注总数，增加其他用户的粉丝总数，创建关注记录
func Create(ctx context.Context, currentId int64, toUserId int64) error {
	relationRaw := &Relation{
		UserId:   currentId,
		ToUserId: toUserId,
	}
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table("user").Where("id = ?", currentId).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error
		if err != nil {
			klog.Error("add user follow_count fail " + err.Error())
			return err
		}

		err = tx.Table("user").Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error
		if err != nil {
			klog.Error("add user follower_count fail " + err.Error())
			return err
		}

		err = tx.Table("relation").Create(&relationRaw).Error
		if err != nil {
			klog.Error("create relation record fail " + err.Error())
			return err
		}

		return nil
	})
	return nil
}

// 减少当前用户的关注总数，减少其他用户的粉丝总数，删除关注记录
func Delete(ctx context.Context, currentId int64, toUserId int64) error {
	var relationRaw *Relation
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table("user").Where("id = ?", currentId).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error
		if err != nil {
			klog.Error("sub user follow_count fail " + err.Error())
			return err
		}

		err = tx.Table("user").Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error
		if err != nil {
			klog.Error("sub user follower_count fail " + err.Error())
			return err
		}

		err = tx.Table("relation").Where("user_id = ? AND to_user_id = ?", currentId, toUserId).Delete(&relationRaw).Error
		if err != nil {
			klog.Error("delete relation record faguo " + err.Error())
			return err
		}
		return nil
	})
	return nil
}

// 通过用户id，查询该用户关注的用户，返回两者之间的关注记录
func QueryFollowById(ctx context.Context, userId int64) ([]*Relation, error) {
	var relations []*Relation
	err := DB.WithContext(ctx).Table("relation").Where("user_id = ?", userId).Find(&relations).Error
	if err != nil {
		klog.Error("query follow by id fail " + err.Error())
		return nil, err
	}
	return relations, nil
}

// 通过用户id，查询该用户的粉丝， 返回两者之间的关注记录
func QueryFollowerById(ctx context.Context, userId int64) ([]*Relation, error) {
	var relations []*Relation
	err := DB.WithContext(ctx).Table("relation").Where("to_user_id = ?", userId).Find(&relations).Error
	if err != nil {
		klog.Error("query follower by id fail " + err.Error())
		return nil, err
	}
	return relations, nil
}

// QueryUserInfo query list of user info
func QueryUserInfo(ctx context.Context, uid int64) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("id = ?", uid).Find(&res).Error; err != nil {
		return nil, err
	}
	// if res != 1{
	// 	return 
	// }
	return res, nil
}
