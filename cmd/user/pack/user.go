
package pack

import (
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/user/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user"
	
)

// User pack user info
func User(u *db.User) *user.User {
	if u == nil {
		return nil
	}

	return &user.User{
		UserId: int64(u.ID), 
		Username: u.Username,
		FollowCount: int64(u.FollowCount),
		FollowerCount: int64(u.FollowerCount),
		IsFollow: bool(u.IsFollow ),
		}
}

// Users pack list of user info
func Users(us []*db.User) []*user.User {
	users := make([]*user.User, 0)
	for _, u := range us {
		if temp := User(u); temp != nil {
			users = append(users, temp)
		}
	}
	return users
}
