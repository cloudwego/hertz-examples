# hertz_gen

## Introduce

A demo with `Hertz` and `GORM/Gen`

- Use `proto` IDL to define `HTTP` interface
- Use `hz` to generate code
- Use `Hertz` binding and validate
- Use `GORM/Gen` and `MySQL`

## IDL

This demo use `proto` IDL to define `HTTP` interface. The specific interface define in [user.proto](idl/user/user.proto)

## Code generation tool

This demo use `hz` to generate code. The use of `hz` refers
to [hz](https://www.cloudwego.io/docs/hertz/tutorials/toolkit/toolkit/)

The `hz` commands used can be found in [Makefile](Makefile)

## Binding and Validate

The use of binding and Validate refers
to [Binding and Validate](https://www.cloudwego.io/docs/hertz/tutorials/basic-feature/binding-and-validate/)

## GORM/Gen

GEN: Friendly & Safer GORM powered by Code Generation.

This demo use `GORM/Gen` to operate `MySQL` and refers to [Gen](https://gorm.io/gen/)

## How to run

### Run mysql docker

```bash
cd bizdemo/hertz_gen && docker-compose up
```

### Generate MySQL table

Connect MySQL and execute [user.sql](biz/model/sql/user.sql)

### Run demo

```go
cd bizdemo/hertz_gen
go build -o hertz_gen &&./hertz_gen
```