namespace go relation
enum ErrCode {
    SuccessCode                = 0
    ServiceErrCode             = 90001
    ParamErrCode               = 90002
    MessageIsNullErrCode    = 90003
    AuthorizationFailedErrCode = 90004
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


struct RelationActionRequest {
    1: i64 user_id // 用户id
    2: string token // 用户鉴权token
    3: i64 to_user_id // 对方用户id
    4: i32 action_type // 1-关注，2-取消关注
}

struct RelationActionResponse {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
}
struct RelationFollowListRequest {
    1: i64 user_id // 用户id
    2: string token // 用户鉴权token
}
struct RelationFollowListResponse {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<User> user_list // 用户信息列表
}
struct RelationFollowerListRequest {
    1: i64 user_id // 用户id
    2: string token // 用户鉴权token
}
struct RelationFollowerListResponse {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<User> user_list // 用户列表
}
struct RelationFriendListRequest {
    1: i64 user_id // 用户id
    2: string token // 用户鉴权token
}
struct RelationFriendListResponse {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<FriendUser> user_list // 用户列表
}





struct Message {
    1: i64 id                  // 消息id
    2: i64 to_user_id          // 该消息接收者的id
    3: i64 from_user_id        // 该消息发送者的id
    4: string content         // 消息内容
    5:optional i64 create_time      // 消息创建时间
}


struct MessageChatRequest {
    1: i64 from_user_id          // 用户id
    2: string token       
    3: i64 to_user_id        // 对方用户id
}

struct MessageChatResponse {
    1: i32 status_code
    2: string status_msg
    3: list<Message> messages
    
}

struct MessageActionRequest {
    1: i64 from_user_id           // 用户鉴权token
    2: string token
    3: i64 to_user_id         // 对方用户id
    4: i64 action_type       // 1-发送消息
    5: string content                // 消息内容
}

struct MessageActionResponse {
    1: i32 status_code
    2: string status_msg
}

service RelationService {
    RelationActionResponse RelationAction (1: RelationActionRequest req)
    RelationFollowListResponse RelationFollowList (1: RelationFollowListRequest req)
    RelationFollowerListResponse RelationFollowerList (1: RelationFollowerListRequest req)
    RelationFriendListResponse RelationFriendList (1: RelationFriendListRequest req)
    MessageChatResponse MessageChat(1: MessageChatRequest req)               // 消息记录
    MessageActionResponse MessageAction(1: MessageActionRequest req)         // 发送消息
}

