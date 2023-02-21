
package pack

import (
<<<<<<< HEAD
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/interact/dal/db"
=======
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/user/dal/db"
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/interact"
	
)

<<<<<<< HEAD
// User pack interact info
func User(u *db.User) *interact.User {
=======
// User pack user info
func User(u *db.User) *user.User {
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	if u == nil {
		return nil
	}

<<<<<<< HEAD
	return &interact.User{
		Id: int64(u.ID), 
		// Name: u.Username,
=======
	return &user.User{
		Id: int64(u.ID), 
		Name: u.Username,
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
		FollowCount: int64(u.FollowCount),
		FollowerCount: int64(u.FollowerCount),
		IsFollow: bool(u.IsFollow),
		// Avatar: string(u.Avatar),
		// BackgroundImage: string(u.BackgroundImage),
		// Signature: string(u.Signature),
		// TotalFavorited: string(u.TotalFavorited),
		// WorkCount: int64(u.WorkCount),
		// FavoriteCount: int64(u.FavoriteCount),
		}
}

<<<<<<< HEAD
// Users pack list of interact info
func Users(us []*db.User) []*interact.User {
	interacts := make([]*interact.User, 0)
	for _, u := range us {
		if temp := User(u); temp != nil {
			interacts = append(interacts, temp)
		}
	}
	return interacts
}

// User pack interact info
func Video(v *db.Video, author *db.User) *interact.Video {
=======
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

// User pack user info
func Video(v *db.Video, author *db.User) *user.Video {
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	if v == nil {
		return nil
	}

<<<<<<< HEAD
	return &interact.Video{
=======
	return &user.Video{
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
		Id: int64(v.ID), 
		Author: User(author),
		PlayUrl: string(v.PlayUrl),
		CoverUrl: string(v.CoverUrl),
		FavoriteCount: int64(v.FavoriteCount),
		CommentCount: int64(v.CommentCount),
		}
}

<<<<<<< HEAD
// Users pack list of interact info
func Videos(vs []*db.Video, author *db.User) []*interact.Video {
	videos := make([]*interact.Video, 0)
=======
// Users pack list of user info
func Videos(vs []*db.Video, author *db.User) []*user.Video {
	videos := make([]*user.Video, 0)
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	for _, v := range vs {
		if temp := Video(v, author); temp != nil {
			videos = append(videos, temp)
		}
	}
	return videos
}