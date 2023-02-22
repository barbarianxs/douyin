package pack

import (
	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/interact/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/interact"
)

// VideoList pack video list info
func VideoList(currentId int64, videoData []*db.Video, userMap map[int64]*db.User, favoriteMap map[int64]*db.Favorite, relationMap map[int64]*db.Relation) []*interact.Video {
	videoList := make([]*interact.Video, 0)
	for _, video := range videoData {
		videoUser, ok := userMap[video.UserId]
		if !ok {
			videoUser = &db.User{
				Name:          "未知用户",
				FollowCount:   0,
				FollowerCount: 0,
			}
			videoUser.ID = 0
		}

		var isFavorite bool = false
		var isFollow bool = false

		if currentId != -1 {
			_, ok := favoriteMap[int64(video.ID)]
			if ok {
				isFavorite = true
			}
			_, ok = relationMap[video.UserId]
			if ok {
				isFollow = true
			}
		}
		videoList = append(videoList, &interact.Video{
			Id: int64(video.ID),
			Author: &interact.User{
				Id:            int64(videoUser.ID),
				Name:          videoUser.Name,
				FollowCount:   videoUser.FollowCount,
				FollowerCount: videoUser.FollowerCount,
				IsFollow:      isFollow,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		})
	}

	return videoList
}