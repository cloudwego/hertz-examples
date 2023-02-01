# hertz_session

## Introduce

A demo with `Hertz` and `Casbin`, this demo aims to demonstrate a distributed session solution based on the `hertz-contrib/sessions`.

The distributed session solution based on redis is to store the sessions of different servers in redis or redis cluster, 
which aims to solve the problem that the sessions of multiple servers are not synchronized in the case of distributed system.

- Use `thrift` IDL to define `HTTP` interface
- Use `hz` to generate code
- Use `hertz-contrib/sessions` to store sessions
- Use `casbin` to judgment authority 
- Use `Gorm` and `MySQL`

## hertz-contrib/sessions

Use Hertz's sessions extension, refer to [hertz-contrib/sessions](https://github.com/hertz-contrib/sessions)

## casbin

Simplistic Example of role-based HTTP Authorization with [casbin](https://github.com/casbin/casbin) using [scs](https://github.com/alexedwards/scs) for session handling.

## IDL

This demo use `thrift` IDL to define `HTTP` interface. The specific interface define in [user.thrift](idl/user.thrift)

### Generate MySQL table

Connect MySQL and execute [user.sql](biz/model/sql/user.sql)

### Permission list

This is where the permissions are defined [policy.csv](conf/policy.csv),each http request is determined.



## Code generation tool

This demo use `hz` to generate code. The use of `hz` refers to [hz](https://www.cloudwego.io/docs/hertz/tutorials/toolkit/toolkit/)

The `hz` commands used can be found in [Makefile](Makefile)

## Binding and Validate

The use of binding and Validate refers
to [Binding and Validate](https://www.cloudwego.io/docs/hertz/tutorials/basic-feature/binding-and-validate/)

## Gorm

This demo use `Gorm` to operate `MySQL` and refers to [Gorm](https://gorm.io/)

## AdminLTE

This demo captures the [AdminLTE](https://github.com/ColorlibHQ/AdminLTE) login and registration page as the front-end page.

## How to run

### Run MySQL and Redis docker

```bash
cd bizdemo/hertz_casbin && docker-compose up
```

### Run demo

```bash
cd bizdemo/hertz_casbin
go run .
```

Which starts a server at `http://localhost:8888` with the following routes:

* `POST /login` - accessible if not logged in
    * takes `username` as a form-data parameter - password is 123
    * Valid Users:
        * `admin` ID: `1`, Role: `admin`
        * `darren` ID: `2`, Role: `member`
* `POST /logout` - accessible if logged in
* `GET /findUser` - accessible if logged in as a member
* `GET /member/list` - accessible if logged in as a member && admin
* `GET /admin/list` - accessible if logged in as an admin 
 