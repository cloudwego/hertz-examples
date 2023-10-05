# hertz_jwt

## Introduce

A demo with `Hertz` and `JWT`

- Use `hz` to generate code
- Use `JWT` to complete login and authentication
- Use `Gorm` and `MySQL`


## JWT

Use Hertz's JWT extension, refer to [jwt](https://github.com/hertz-contrib/jwt)

## Code generation tool

This demo use `hz` to generate code. The use of `hz` refers
to [hz](https://www.cloudwego.io/docs/hertz/tutorials/toolkit/toolkit/)

The `hz` commands used can be found in [Makefile](Makefile)

## Binding and Validate

The use of binding and Validate refers
to [Binding and Validate](https://www.cloudwego.io/docs/hertz/tutorials/basic-feature/binding-and-validate/)

## Gorm

This demo use `Gorm` to operate `MySQL` and refers to [Gorm](https://gorm.io/)

## How to run

### Run mysql docker

```bash
cd bizdemo/hertz_jwt && docker-compose up
```

### Run demo

```go
cd bizdemo/hertz_jwt && go run main.go
```