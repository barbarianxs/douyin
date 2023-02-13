namespace go api

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




service ApiService {
    LoginUserResponse LoginUser(1: LoginUserRequest req) (api.post="/douyin/user/login")
    RegisterUserResponse RegisterUser(1: RegisterUserRequest req) (api.post="/douyin/user/register")

}