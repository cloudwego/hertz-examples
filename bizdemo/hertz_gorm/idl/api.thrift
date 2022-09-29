// refer to  https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/toolkit/

namespace go user_gorm
namespace py user_gorm
namespace java user_gorm

enum Code {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}

enum Gender {
    Unknown = 0
    Male    = 1
    Female  = 2
}

struct User {
    1: i64 user_id
    2: string name
    3: Gender gender
    4: i64    age
    5: string introduce
}

struct CreateUserRequest{
    1: string name      (api.body="name", api.form="name",api.vd="(len($) > 0 && len($) < 100)")
    2: Gender gender    (api.body="gender", api.form="gender",api.vd="($ == 1||$ == 2)")
    3: i64    age       (api.body="age", api.form="age",api.vd="$>0")
    4: string introduce (api.body="introduce", api.form="introduce",api.vd="(len($) > 0 && len($) < 1000)")
}

struct CreateUserResponse{
   1: Code code
   2: string msg
}

struct QueryUserRequest{
   1: optional string Keyword (api.body="keyword", api.form="keyword",api.query="keyword")
   2: i64 page (api.body="page", api.form="page",api.query="page",api.vd="$ > 0")
   3: i64 page_size (api.body="page_size", api.form="page_size",api.query="page_size",api.vd="($ > 0 || $ <= 100)")
}

struct QueryUserResponse{
   1: Code code
   2: string msg
   3: list<User> users
   4: i64 totoal
}

struct DeleteUserRequest{
   1: i64    user_id   (api.path="user_id",api.vd="$>0")
}

struct DeleteUserResponse{
   1: Code code
   2: string msg
}

struct UpdateUserRequest{
    1: i64    user_id   (api.path="user_id",api.vd="$>0")
    2: string name      (api.body="name", api.form="name",api.vd="(len($) > 0 && len($) < 100)")
    3: Gender gender    (api.body="gender", api.form="gender",api.vd="($ == 1||$ == 2)")
    4: i64    age       (api.body="age", api.form="age",api.vd="$>0")
    5: string introduce (api.body="introduce", api.form="introduce",api.vd="(len($) > 0 && len($) < 1000)")
}

struct UpdateUserResponse{
   1: Code code
   2: string msg
}


service UserService {
   UpdateUserResponse UpdateUser(1:UpdateUserRequest req)(api.post="/v1/user/update/:user_id")
   DeleteUserResponse DeleteUser(1:DeleteUserRequest req)(api.post="/v1/user/delete/:user_id")
   QueryUserResponse  QueryUser(1: QueryUserRequest req)(api.post="/v1/user/query/")
   CreateUserResponse CreateUser(1:CreateUserRequest req)(api.post="/v1/user/create/")
}