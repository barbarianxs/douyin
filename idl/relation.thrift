namespace go relation


struct User {
    1: i64 id // 用户id
    2: string name // 用户名称
    3: i64 follow_count // 关注总数
    4: i64 follower_count // 粉丝总数
    5: bool is_follow // true-已关注，false-未关注
}

struct douyin_relation_action_request {
    1: i64 user_id // 用户id
    2: string token // 用户鉴权token
    3: i64 to_user_id // 对方用户id
    4: i32 action_type // 1-关注，2-取消关注
}

struct douyin_relation_action_response {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
}
struct douyin_relation_follow_list_request {
    1: i64 user_id // 用户id
    2: string token // 用户鉴权token
}
struct douyin_relation_follow_list_response {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<User> user_list // 用户信息列表
}
struct douyin_relation_follower_list_request {
    1: i64 user_id // 用户id
    2: string token // 用户鉴权token
}
struct douyin_relation_follower_list_response {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<User> user_list // 用户列表
}
struct douyin_relation_friend_list_request {
    1: i64 user_id // 用户id
    2: string token // 用户鉴权token
}
struct douyin_relation_friend_list_response {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<FriendUser> user_list // 用户列表
}
struct FriendUser {
    1: i64 id // 用户id
    2: string name // 用户名称
    3: i64 follow_count // 关注总数
    4: i64 follower_count // 粉丝总数
    5: bool is_follow // true-已关注，false-未关注
    6: string avatar // 用户头像Url
    7: string message // 和该好友的最新聊天消息
    8: i64 msg_type // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

service RelationService {
    douyin_relation_action_response RelationAction (1: douyin_relation_action_request req)
    douyin_relation_follow_list_response RelationFollowList (1: douyin_relation_follow_list_request req)
    douyin_relation_follower_list_response RelationFollowerList (1: douyin_relation_follower_list_request req)
    douyin_relation_friend_list_response RelationFriendList (1: douyin_relation_friend_list_request req)
}