package db

import (
	"context"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user"
	
	//"github.com/YANGJUNYAN0715/douyin/tree/zhao/cmd/user/dal/db"
	//"gorm.io/gorm"
)

// Video pack user info
func BuildVideo(ctx context.Context, v *Video, fromID int64) (*user.Video, error) {
	if v == nil {
		return nil, nil
	}
	// user, err := db.GetUserByID(ctx, int64(v.AuthorID))
	// if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return nil, err
	// }

	// author, err := User(ctx, user, fromID)
	// if err != nil {
	// 	return nil, err
	// }
	author :=user.User{};
	
	
	favorite_count := int64(v.FavoriteCount)
	comment_count := int64(v.CommentCount)

	return &user.Video{
		Id:            int64(v.ID),
		Author:        &author,
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: favorite_count,
		CommentCount:  comment_count,
		Title:         v.Title,
	}, nil
}

// Videos pack list of video info
func BuildVideos(ctx context.Context, vs [] *Video, fromID *int64) ([]*user.Video, error) {
	videos := make([]*user.Video, 0)
	for _, v := range vs {
		video2, err := BuildVideo(ctx, v, *fromID)
		if err != nil {
			return nil, err
		}

		if video2 != nil {
			flag := false
			// if *fromID != 0 {
			// 	results, err := db.GetFavoriteRelation(ctx, *fromID, int64(v.ID))
			// 	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			// 		return nil, err
			// 	} else if errors.Is(err, gorm.ErrRecordNotFound) {
			// 		flag = false
			// 	} else if results != nil && results.AuthorID != 0 {
			// 		flag = true
			// 	}
			// }
			video2.IsFavorite = flag
			videos = append(videos, video2)
		}
	}
	return videos, nil
}