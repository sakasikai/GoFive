namespace go GoFive

struct BaseResp {
    1:i64 status_code
    2:string status_msg
}

struct User {
    1:i64 id
    2:string name
    3:i64 follow_count
    4:i64 follower_count
    5:bool is_follow
}

//
struct CreateUserRequest {
    1:string user_name
    2:string password
}

struct CreateUserResponse {
    1:BaseResp base_resp
}

//
struct QueryUserByIDRequest {
    1:i64 user_id
}

struct QueryUserByNameRequest {
    1:string user_name
}

struct QueryUserResponse {
    1:User user
    2:BaseResp base_resp
}

//
struct CheckUserResponse {
    1:string user_name
    2:string password
}
struct CheckUserRequest {
    1:i64 user_id
    2:BaseResp base_resp
}

service UserService {
    CreateUserResponse CreateUser(1:CreateUserRequest req)
    QueryUserResponse QueryUserByID(1:QueryUserByIDRequest req)
    QueryUserResponse QueryUserByName(1:QueryUserByNameRequest req)
    CheckUserResponse CheckUser(1:CheckUserRequest req)
//    UpdateUserResponse UpdateUser(1:UpdateUserRequest req)
}


