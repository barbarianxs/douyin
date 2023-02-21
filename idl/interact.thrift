namespace go interact

struct FavoriteActionRequest {
    1: i64 user_id
    2: string token // 用户鉴权token
    3: i64 video_id  // 视频id
    4: i32 action_type   // 1-点赞，2-取消点赞
}

struct FavoriteActionResponse {
    1: i32 status_code   // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
}

struct FavoriteListRequest {
     1: i64 user_id //用户id
     2: string token //用户鉴权token
 }

struct FavoriteListResponse {
    1: i32 status_code //状态码，0-成功，其他值失败
    2: optional string status_msg //返回状态描述
    3: list<Video> video_list //用户点赞视频列表
}

struct User {
    1: i64 id // 用户id
    2: string name // 用户名称
    3: i64 follow_count // 关注总数
    4: i64 follower_count // 粉丝总数
    5: bool is_follow // true-已关注，false-未关注
}

struct Video{
   1: i64 id;
   2: User author;
   3: string play_url;
   4: string cover_url;
   5: i64 favorite_count;
   6: i64 comment_count;
   7: bool is_favorite;
   8: string title;
}

service FavoriteService {
    // 用户点赞
    FavoriteActionResponse FavoriteAction(1:FavoriteActionRequest req)
    // 用户点赞列表
    FavoriteListResponse FavoriteList(1:FavoriteListRequest req)
}