
package pack

import (
	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/interact/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/interact"
	
)

// User pack interact info
func User(u *db.User) *interact.User {
	if u == nil {
		return nil
	}

	return &interact.User{
		Id: int64(u.ID), 
		// Name: u.Username,
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

// Users pack list of interact info
func Users(us []*db.User) []*interact.User {
	users := make([]*interact.User, 0)
	for _, u := range us {
		if temp := User(u); temp != nil {
			users = append(users, temp)
		}
	}
	return users
}

// User pack interact info
func Video(v *db.Video, author *db.User) *interact.Video {
	if v == nil {
		return nil
	}

	return &interact.Video{
		Id: int64(v.ID), 
		Author: User(author),
		PlayUrl: string(v.PlayUrl),
		CoverUrl: string(v.CoverUrl),
		FavoriteCount: int64(v.FavoriteCount),
		CommentCount: int64(v.CommentCount),
		}
}

// Users pack list of interact info 一个作者下的所有视频发布
func Videos(vs []*db.Video, author *db.User) []*interact.Video {
	videos := make([]*interact.Video, 0)
	for _, v := range vs {
		if temp := Video(v, author); temp != nil {
			videos = append(videos, temp)
		}
	}
	return videos
}


// Comment pack interact info
func Comment(c *db.Comment, u *db.User) *interact.Comment {
	if u == nil {
		return nil
	}

	return &interact.Comment{
		Id: int64(c.ID), 
		User: User(u),
		// VideoId: int64(c.VideoId),
		Content: string(c.Content),
		CreateDate: string(c.CreateDate),
		
		}
}


// Users pack list of interact info
// func Comments(cs []*db.Comment) []*interact.Comment {
// 	comment_list := make([]*interact.Comment, 0)
// 	for _, c := range cs {
// 		if temp := Comment(c); temp != nil {
// 			comment_list = append(comment_list, temp)
// 		}
// 	}
// 	return comment_list
// }
