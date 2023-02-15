namespace go api

enum ErrCode {
    SuccessCode                = 0
    ServiceErrCode             = 10001
    ParamErrCode               = 10002
    UserAlreadyExistErrCode    = 10003
    AuthorizationFailedErrCode = 10004
}

struct BaseResp {
    1: i64 status_code
    2: string status_message
    3: i64 service_time
}

struct User {
    1: i64 user_id
    2: string username
    3: string avatar
}

struct Message {
    1:required i64 id                  // 消息id
    2:required i64 to_user_id          // 该消息接收者的id
    3:required i64 from_user_id        // 该消息发送者的id
    4:required string content         // 消息内容
    5:optional string create_time      // 消息创建时间
}

struct LoginUserRequest {
    1: string username (api.form="username", api.vd="len($) > 0")
    2: string password (api.form="password", api.vd="len($) > 0")
}

struct LoginUserResponse {
    1: BaseResp base_resp
}

struct RegisterUserRequest {
    1: string username (api.form="username", api.vd="len($) > 0")
    2: string password (api.form="password", api.vd="len($) > 0")
}

struct RegisterUserResponse {
    1: BaseResp base_resp
}


struct MessageChatRequest {
    1:required i64 from_user_id    (api.chat="from_user_id", api.vd="len($) > 0")      // 用户鉴权token
    2:required i64 to_user_id  (api.chat="to_user_id", api.vd="len($) > 0")      // 对方用户id
}

struct MessageChatResponse {
    1: list<Message> messages
    2: BaseResp base_resp
    
}

struct MessageActionRequest {
    1:required i64 from_user_id       (api.form="from_user_id", api.vd="len($) > 0")               // 用户鉴权token
    2:required i64 to_user_id    (api.form="to_user_id", api.vd="len($) > 0")      // 对方用户id
    3:required i32 action_type    (api.form="action_type", api.vd="len($) > 0")    // 1-发送消息
    4:required string content      (api.form="content", api.vd="len($) > 0")           // 消息内容
}

struct MessageActionResponse {
    1: BaseResp base_resp
}


service ApiService {
    LoginUserResponse LoginUser(1: LoginUserRequest req) (api.post="/douyin/user/login/")
    RegisterUserResponse RegisterUser(1: RegisterUserRequest req) (api.post="/douyin/user/register/")
    MessageChatResponse MessageChat(1: MessageChatRequest req) (api.get="/douyin/message/chat/")      // 消息记录
    MessageActionResponse MessageAction(1: MessageActionRequest req) (api.post="/douyin/message/action/")     // 发送消息
}