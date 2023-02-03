# hertz_casbin

## Introduce

A demo with `Hertz` and `Casbin`, this demo aims to understand the application of rbac.

Casbin is a powerful and efficient open-source access control library for Golang projects. It provides support for enforcing authorization based on various access control models.

- Use `thrift` IDL to define `HTTP` interface
- Use `hz` to generate code
- Use `casbin` to judgment authority
- Use `Gorm` and `MySQL`
- Use `JWT` to complete login and authentication
 

## casbin

Simplistic Example of role-based HTTP Authorization with [casbin](https://github.com/casbin/casbin) using [scs](https://github.com/alexedwards/scs) for session handling.

## IDL

This demo use `thrift` IDL to define `HTTP` interface. The specific interface define in [user.thrift](idl/casbin.thrift)

## Code generation tool

This demo use `hz` to generate code. The use of `hz` refers to [hz](https://www.cloudwego.io/docs/hertz/tutorials/toolkit/toolkit/)

The `hz` commands used can be found in [Makefile](Makefile)

## Binding and Validate

The use of binding and Validate refers
to [Binding and Validate](https://www.cloudwego.io/docs/hertz/tutorials/basic-feature/binding-and-validate/)

## Gorm

This demo use `Gorm` to operate `MySQL` and refers to [Gorm](https://gorm.io/)

 
## How to run

### Run MySQL  

```bash
cd bizdemo/hertz_casbin && docker-compose up
```

### Generate MySQL table

Connect MySQL and execute [casbin.sql](biz/model/sql/casbin.sql)

### Run demo

```bash
cd bizdemo/hertz_casbin
go run .

```


 