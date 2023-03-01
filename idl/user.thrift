namespace go user

enum ErrCode {
    SuccessCode                = 0
    ServiceErrCode             = 10001
    ParamErrCode               = 10002
    UserAlreadyExistErrCode    = 10003
    AuthorizationFailedErrCode = 10004
}

struct BaseResp {
    1: i32 status_code
    2: string status_msg
    3: i64 service_time
}

struct UserLogin {
    1: string username // 注册用户名，最长32个字符
    2: string password // 密码，最长32个字符
}

struct User {
    1: i64 id
    2: string name
    3: i64 follow_count // 关注总数
    4: i64 follower_count  // 粉丝总数
    5: bool is_follow  // true-已关注，false-未关注

}

struct Video {
    1: i64 id;
    2: User author;
    3: string play_url;
    4: string cover_url;
    5: i64 favorite_count;
    6: i64 comment_count;
    7: bool is_favorite;
    8: string title;
}

struct LoginUserRequest {
    1: string username (vt.min_size = "1")
    2: string password (vt.min_size = "1")
}

struct LoginUserResponse {
    1: i32 status_code
    2: string status_msg
    3: i64 user_id
    4: string token

}

struct LogoutUserRequest {
    1: string username
    2: string password
}

struct LogoutUserResponse {
    1: i32 status_code
    2: string status_msg
    3: i64 user_id
    4: string token

}

struct RegisterUserRequest {
    1: string username
    2: string password
}

struct RegisterUserResponse {
    1: i32 status_code
    2: string status_msg
    3: i64 user_id
    4: string token

}
struct UserInfoRequest {
    1: i64 user_id
    2: string token
}

struct UserInfoResponse {
    1: i32 status_code
    2: string status_msg
    3: User user

}

struct PublishActionRequest {
    1: i64 user_id;
    2: string file_url;
    3: string cover_url;
    4: string title;
}

struct PublishActionResponse {
    1: i32 status_code;
    2: string status_msg;
}

struct PublishListRequest {
    1: i64 user_id;
    2: string token;
}

struct PublishListResponse {
    1: i32 status_code;
    2: string status_msg;
    3: list<Video> video_list;
}



struct FeedRequest {
    1: i64 latest_time; // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: string token; // 可选参数，登录用户设置
    3: i64 user_id
}

// 例如当前请求的latest_time为9:00，那么返回的视频列表时间戳为[8:55,7:40, 6:30, 6:00]
// 所有这些视频中，最早发布的是 6:00的视频，那么6:00作为下一次请求时的latest_time
// 那么下次请求返回的视频时间戳就会小于6:00

struct FeedResponse {
    1: i32 status_code; // 状态码，0-成功，其他值-失败
    2: string status_msg; // 返回状态描述
    3: list<Video> video_list; // 视频列表
    4: i64 next_time; // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

struct VideoIdRequest{
    1:i64 video_id ;
    2:i64 search_id ;
}





service UserService {
    LoginUserResponse LoginUser(1: LoginUserRequest req)
    LogoutUserResponse LogoutUser(1: LogoutUserRequest req)
    RegisterUserResponse RegisterUser(1: RegisterUserRequest req)
    UserInfoResponse UserInfo(1: UserInfoRequest req)
    PublishActionResponse PublishAction(1: PublishActionRequest req)
    PublishListResponse PublishList(1: PublishListRequest req)
    FeedResponse GetUserFeed (1:FeedRequest req)
}