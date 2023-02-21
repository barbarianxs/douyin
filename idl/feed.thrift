
namespace go feed

struct douyin_feed_request {
    1: i64 latest_time; // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: string token; // 可选参数，登录用户设置
}

// 例如当前请求的latest_time为9:00，那么返回的视频列表时间戳为[8:55,7:40, 6:30, 6:00]
// 所有这些视频中，最早发布的是 6:00的视频，那么6:00作为下一次请求时的latest_time
// 那么下次请求返回的视频时间戳就会小于6:00
struct douyin_feed_response {
    1: i32 status_code; // 状态码，0-成功，其他值-失败
    2: string status_msg; // 返回状态描述
    3: list<Video> video_list; // 视频列表
    4: i64 next_time; // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

struct video_id_request{
    1:i64 video_id ;
    2:i64 search_id ;
}

struct Video {
    1:i64 id = 1; // 视频唯一标识
    2:User author; // 视频作者信息
    3:string play_url; // 视频播放地址
    4:string cover_url; // 视频封面地址
    5:i64 favorite_count; // 视频的点赞总数
    6:i64 comment_count; // 视频的评论总数
    7:bool is_favorite; // true-已点赞，false-未点赞
    8:string title; // 视频标题
}

struct User {
    1: i64 id // 用户id
    2: string name // 用户名称
    3: i64 follow_count // 关注总数
    4: i64 follower_count // 粉丝总数
    5: bool is_follow // true-已关注，false-未关注
}
service FeedService {
    douyin_feed_response GetUserFeed (1:douyin_feed_request req)
}

