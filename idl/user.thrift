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
    2: string token;
    3: string file_url;
    4: string cover_url;
    5: string title;
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

service UserService {
    LoginUserResponse LoginUser(1: LoginUserRequest req)
    LogoutUserResponse LogoutUser(1: LogoutUserRequest req)
    RegisterUserResponse RegisterUser(1: RegisterUserRequest req)
    UserInfoResponse UserInfo(1: UserInfoRequest req)
    PublishActionResponse PublishAction(1: PublishActionRequest req)
    PublishListResponse PublishList(1: PublishListRequest req)
}