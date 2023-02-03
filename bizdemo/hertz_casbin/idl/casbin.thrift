
namespace go casbin

enum Code {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}



struct BasicResponse{
   1: Code code
   2: string msg
}



/***********Permission***************/
struct Permission {
    1: i64 id
    2: string v1
    3: string v2
}

struct PermissionRole {
    1: i64 id
    2: i64 rid
    3: i64 pid
}


struct CreatePermissionRequest{
    1: string v1      (api.body="v1", api.form="v1",api.vd="(len($) > 0 && len($) < 100)")
    2: string v2     (api.body="v2", api.form="v2",api.vd="(len($) > 0 && len($) < 100)")
}


struct CreatePermissionResponse{
   1: Code code
   2: string msg
   3: Permission permission
}



struct BindPermissionRoleRequest{
    1: string pid      (api.body="pid", api.form="pid",api.vd="($ > 0 && $ < 100)")
    2: string rid     (api.body="rid", api.form="rid",api.vd="($ > 0 && $ < 100)")
}


struct BindPermissionRoleResponse{
   1: Code code
   2: string msg
   3: PermissionRole permissionRole
}


service PermissionService {
   CreatePermissionResponse CreatePermission(1:CreatePermissionRequest req)(api.post="/v1/permission/create/")
   BindPermissionRoleResponse BindPermissionRole(1:BindPermissionRoleRequest req)(api.post="/v1/permissionrole/bind/")
}



/***********Role***************/
struct Role {
    1: i64 id
    2: string name
}

struct CreateRoleRequest{
    1: string name      (api.body="name", api.form="name",api.vd="(len($) > 0 && len($) < 100)")
}


struct CreateRoleResponse{
   1: Code code
   2: string msg
   3: Role role
}


struct BindRoleRequest{
    1: i64 uid      (api.body="uid", api.form="uid",api.vd="($ > 0)")
    2: i64 rid      (api.body="rid", api.form="rid",api.vd="($ > 0)")
}

struct BindRoleResponse{
   1: Code code
   2: string msg
}




service RoleService {
   CreateRoleResponse CreateRole(1:CreateRoleRequest req)(api.post="/v1/role/create/")
   BindRoleResponse BindRole(1:BindRoleRequest req)(api.post="/v1/role/bind/")

}



/***********user***************/
struct User {
    1: i64 id
    2: string username
    3: string password
}

struct UserRole {
    1: i64 id
    2: i64 rid
    3: i64 uid
}



struct LoginRequest{
    1: string username      (api.body="username", api.form="username",api.vd="(len($) > 0 && len($) < 100)")
    2: string password    (api.body="password", api.form="password",api.vd="(len($) > 0 && len($) < 100)")
}

struct QueryUserResponse{
   1: Code code
   2: string msg
   3: string token
}

service UserService {
   QueryUserResponse  Login(1: LoginRequest req)(api.post="/v1/login")
}

