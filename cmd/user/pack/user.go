
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
		Id: int64(u.ID), 
		Name: u.Name,
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
	if v == nil {
		return nil
	}

	return &user.Video{
		Id: int64(v.ID), 
		Author: User(author),
		PlayUrl: string(v.PlayUrl),
		CoverUrl: string(v.CoverUrl),
		FavoriteCount: int64(v.FavoriteCount),
		CommentCount: int64(v.CommentCount),
		}
}

// Users pack list of user info
func Videos(vs []*db.Video, author *db.User) []*user.Video {
	videos := make([]*user.Video, 0)
	for _, v := range vs {
		if temp := Video(v, author); temp != nil {
			videos = append(videos, temp)
		}
	}
	return videos
}