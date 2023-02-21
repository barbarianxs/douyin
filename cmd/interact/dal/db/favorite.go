package db

import (
	"context"
<<<<<<< HEAD
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/consts"
=======
	"github.com/YANGJUNYAN0715/douyin/tree/li/pkg/consts"
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

type Favorite struct {
	gorm.Model
	UserId  int64 `gorm:"column:user_id;not null;index:idx_userid"`
	VideoId int64 `gorm:"column:video_id;not null;index:idx_videoid"`
}

func (Favorite) TableName() string {
	return "favorite"
}

// 根据当前用户id和视频id获取点赞信息
func QueryFavoriteByIds(ctx context.Context, currentId int64, videoIds []int64) (map[int64]*Favorite, error) {
	var favorites []*Favorite
	err := DB.WithContext(ctx).Where("user_id = ? AND video_id IN ?", currentId, videoIds).Find(&favorites).Error
	if err != nil {
		klog.Error("quert favorite record fail " + err.Error())
		return nil, err
	}
	favoriteMap := make(map[int64]*Favorite)
	for _, favorite := range favorites {
		favoriteMap[favorite.VideoId] = favorite
	}
	return favoriteMap, nil
}

<<<<<<< HEAD
// CreateFavorite add a record to the favorite table through a transaction, and add the number of video guokes
=======
// CreateFavorite add a record to the favorite table through a transaction, and add the number of video likes
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
func CreateFavorite(ctx context.Context, favorite *Favorite, videoId int64) error {
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table("video").Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
		if err != nil {
			klog.Error("AddFavoriteCount error " + err.Error())
			return err
		}

		err = tx.Table("favorite").Create(favorite).Error
		if err != nil {
			klog.Error("create favorite record fail " + err.Error())
			return err
		}

		return nil
	})
	return nil
}

<<<<<<< HEAD
// DeleteFavorite Delete a record in the favorite table and reduce the number of video guokes
=======
// DeleteFavorite Delete a record in the favorite table and reduce the number of video likes
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
func DeleteFavorite(ctx context.Context, currentId int64, videoId int64) error {
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var favorite *Favorite
		err := tx.Table("favorite").Where("user_id = ? AND video_id = ?", currentId, videoId).Delete(&favorite).Error
		if err != nil {
			klog.Error("delete favorite record fail " + err.Error())
			return err
		}

		err = tx.Table("video").Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
		if err != nil {
			klog.Error("SubFavoriteCount error " + err.Error())
			return err
		}
		return nil
	})
	return nil
}

// QueryFavoriteById 通过一个用户id查询出该用户点赞的所有视频id号
func QueryFavoriteById(ctx context.Context, userId int64) ([]int64, error) {
	var favorites []*Favorite
	err := DB.WithContext(ctx).Table("favorite").Where("user_id = ?", userId).Find(&favorites).Error
	if err != nil {
		klog.Error("query favorite record fail " + err.Error())
		return nil, err
	}
	videoIds := make([]int64, 0)
	for _, favorite := range favorites {
		videoIds = append(videoIds, favorite.VideoId)
	}
	return videoIds, nil
}

// Video Gorm Data Structures
type Video struct {
	gorm.Model
	UserId        int64     `gorm:"column:user_id;not null;index:idx_userid"`
	Title         string    `gorm:"column:title;type:varchar(128);not null"`
	PlayUrl       string    `gorm:"column:play_url;varchar(128);not null"`
	CoverUrl      string    `gorm:"column:cover_url;varchar(128);not null"`
	FavoriteCount int64     `gorm:"column:favorite_count;default:0"`
	CommentCount  int64     `gorm:"column:comment_count;default:0"`
	UpdatedAt     time.Time `gorm:"column:update_time;not null;index:idx_update"`
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
func QueryVideoByVideoIds(ctx context.Context, videoIds []int64) ([]*Video, error) {
	var videos []*Video
	err := DB.WithContext(ctx).Where("id in (?)", videoIds).Find(&videos).Error
	if err != nil {
		klog.Error("QueryVideoByVideoIds error " + err.Error())
		return nil, err
	}
	return videos, nil
}

type User struct {
	gorm.Model
<<<<<<< HEAD
	Username string `json:"username"`
=======
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	Name          string `gorm:"column:name;index:idx_username,unique;type:varchar(32);not null"`
	Password      string `gorm:"column:password;type:varchar(32);not null"`
	FollowCount   int64  `gorm:"column:follow_count;default:0"`
	FollowerCount int64  `gorm:"column:follower_count;default:0"`
<<<<<<< HEAD
	IsFollow      bool   `json:"is_follow"`
=======
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
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
	err := DB.WithContext(ctx).Where("user_id = ? AND to_user_id IN ?", currentId, userIds).Find(&relations).Error
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
<<<<<<< HEAD
			klog.Error("delete relation record faguo " + err.Error())
=======
			klog.Error("delete relation record fali " + err.Error())
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
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