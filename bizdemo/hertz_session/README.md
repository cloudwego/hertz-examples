# hertz_session

## Introduce

A demo with `Hertz` and `Session`

- Use `thrift` IDL to define `HTTP` interface
- Use `hz` to generate code
- Use `hertz-contrib/sessions` to store sessions
- Use `Gorm` and `MySQL`

## hertz-contrib/sessions

Use Hertz's sessions extension, refer to [hertz-contrib/sessions](https://github.com/hertz-contrib/sessions)

## IDL

This demo use `thrift` IDL to define `HTTP` interface. The specific interface define in [user.thrift](idl/user.thrift)

## Code generation tool

This demo use `hz` to generate code. The use of `hz` refers to [hz](https://www.cloudwego.io/docs/hertz/tutorials/toolkit/toolkit/)

The `hz` commands used can be found in [Makefile](Makefile)

## Binding and Validate

The use of binding and Validate refers
to [Binding and Validate](https://www.cloudwego.io/docs/hertz/tutorials/basic-feature/binding-and-validate/)

## Gorm

This demo use `Gorm` to operate `MySQL` and refers to [Gorm](https://gorm.io/)

## How to run

### Run MySQL and Redis docker

```bash
cd bizdemo/hertz_session && docker-compose up
```

### Run demo

```bash
cd bizdemo/hertz_session
go build -o hertz_session && ./hertz_session
```