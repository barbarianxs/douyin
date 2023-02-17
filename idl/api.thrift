namespace go api

enum ErrCode {
    SuccessCode                = 0
    ServiceErrCode             = 10001
    ParamErrCode               = 10002
    UserAlreadyExistErrCode    = 10003
    AuthorizationFailedErrCode = 10004
}

struct BaseResp {
    1: i32 status_code
    2: string status_message
    3: i64 service_time
}

struct User {
    1: i64 user_id
    2: string username
    3: string avatar
}

struct Message {
    1:i64 id                  // 消息id
    2:i64 to_user_id          // 该消息接收者的id
    3:i64 from_user_id        // 该消息发送者的id
    4:string content         // 消息内容
    5:optional string create_time      // 消息创建时间
}

struct LoginUserRequest {
    1: string username
    2: string password
}

struct LoginUserResponse {
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
struct MessageChatRequest {
    1:i64 from_user_id  
    2:i64 to_user_id 
}

struct MessageChatResponse {
    1: list<Message> messages
    2: BaseResp base_resp
    
}

struct MessageActionRequest {
    1:i64 from_user_id   
    2:i64 to_user_id
    3:i32 action_type
    4:string content
}

struct MessageActionResponse {
    1: BaseResp base_resp
}


service ApiService {
    LoginUserResponse LoginUser(1: LoginUserRequest req) (api.post="/douyin/user/login/")
    RegisterUserResponse RegisterUser(1: RegisterUserRequest req) (api.post="/douyin/user/register/")
    UserInfoResponse UserInfo(1: UserInfoRequest req) (api.get="/douyin/user/")
    MessageChatResponse MessageChat(1: MessageChatRequest req) (api.get="/douyin/message/chat/")      // 消息记录
    MessageActionResponse MessageAction(1: MessageActionRequest req) (api.post="/douyin/message/action/")     // 发送消息
}