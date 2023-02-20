package db

import(
	"context"
	// "log"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/consts"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/relation"
	"gorm.io/gorm"
)

// Relation Gorm data structure
type User struct {
	gorm.Model
	Username       string  `gorm:"index:idx_username,unique;type:varchar(40);not null" json:"username"`
	Password       string  `gorm:"type:varchar(256);not null" json:"password"`
	// FavoriteVideos []Video `gorm:"many2many:user_favorite_videos" json:"favorite_videos"`
	FollowingCount int     `gorm:"default:0" json:"following_count"`
	FollowerCount  int     `gorm:"default:0" json:"follower_count"`
}

// Relation表 记录关注关系
// 不设置外键 提高效率 通过程序保证参照完整性
type Relation struct {
	gorm.Model
	UserID   int  `gorm:"index:idx_userid;not null"`
	ToUserID int  `gorm:"index:index:idx_userid_to;not null"`
}



func (u *User) TableName() string {
	return consts.RelationTableName
}


// GetRelation get relation info
func GetRelation(ctx context.Context, uid int64, tid int64) (*Relation, error) {
	relations := new(Relation)

	if err := DB.WithContext(ctx).First(&relations, "user_id = ? and to_user_id = ?", uid, tid).Error; err != nil {
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
		err := tx.Create(&Relation{UserID: int(uid), ToUserID: int(tid)}).Error
		if err != nil {
			return err
		}

		// 2.改变 user 表中的 following count
		res := tx.Table(consts.UserTableName).Where("ID = ?", uid).Update("following_count", gorm.Expr("following_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.RelationActionErr
		}

		// 3.改变 user 表中的 follower count
		res = tx.Table(consts.UserTableName).Where("ID = ?", tid).Update("follower_count", gorm.Expr("follower_count + ?", 1))
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
		if err := tx.Where("user_id = ? AND to_user_id=?", uid, tid).First(&relations).Error; err != nil {
			return err
		}

		// 1. 删除关注数据
		err := tx.Unscoped().Delete(&relations).Error
		if err != nil {
			return err
		}
		// 2.改变 user 表中的 following count
		res := tx.Table(consts.UserTableName).Where("ID = ?", uid).Update("following_count", gorm.Expr("following_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.RelationActionErr
		}

		// 3.改变 user 表中的 follower count
		res = tx.Table(consts.UserTableName).Where("ID = ?", tid).Update("follower_count", gorm.Expr("follower_count - ?", 1))
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
	err := DB.WithContext(ctx).Where("user_id = ?", uid).Find(&RelationList).Error
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
		userIDs= append(userIDs,int64(u.UserID))
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
	err := DB.WithContext(ctx).Where("user_id = ?", id).Find(&LRelationList).Error
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
	}

	RuserIDs :=make([]int64,0)
	for _,u := range RRelationList{
		RuserIDs= append(RuserIDs,int64(u.UserID))
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
	// log.Println(userIDs)
	users, err := MGetUsers(ctx,userIDs)
	if err != nil {
		return nil, err
	}
	return BuildFriendUsers(ctx,id,users)
}