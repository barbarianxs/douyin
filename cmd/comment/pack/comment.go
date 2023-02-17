package pack

import (
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/cmd/user/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/kitex_gen/user"
)

// User pack user info
//打包 user 还没用到， 后续需要改，followcount,followercount,isfollow都需要从数据库查询
func BuildUser(u *db.User) *user.User {
	if u == nil {
		return nil
	}
	return &user.User{
		Id: int64(u.ID),
		Name: u.Username,
		FollowCount: int64(u.FollowingCount),
		FollowerCount:int64(u.FollowerCount),
		IsFollow:true,
	}
}

// Users pack list of user info
// func Users(us []*db.User) []*user.User {
// 	users := make([]*user.User, 0)
// 	for _, u := range us {
// 		if temp := User(u); temp != nil {
// 			users = append(users, temp)
// 		}
// 	}
// 	return users
// }
