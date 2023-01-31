namespace go user

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

struct LoginUserRequest {
    1: string username (vt.min_size = "1")
    2: string password (vt.min_size = "1")
}

struct LoginUserResponse {
    1: BaseResp base_resp
}

struct LogoutUserRequest {
    1: string username (vt.min_size = "1")
    2: string password (vt.min_size = "1")
}

struct LogoutUserResponse {
    1: BaseResp base_resp
}
service UserService {
    LoginUserResponse LoginUser(1: LoginUserRequest req)
    LogoutUserResponse LogoutUser(1: LogoutUserRequest req)

    CheckUserResponse CheckUser(1: CheckUserRequest req)
}