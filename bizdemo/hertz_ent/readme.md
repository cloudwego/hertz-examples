# hertz_ent

## Introduce

A demo with `Hertz` and `Ent`

- Use `proto` IDL to define `HTTP` interface
- Use `hz` to generate code
- Use `Hertz` binding and validate
- Use `Ent` and `MySQL`

## IDL

This demo use `proto` IDL to define `HTTP` interface. The specific user interface define in [user.proto](idl/user/user.proto).

## Code generation tool

This demo use `hz` to generate code. The use of `hz` refers
to [hz](https://www.cloudwego.io/docs/hertz/tutorials/toolkit/toolkit/).

The `hz` commands used can be found in [Makefile](Makefile).

## Binding and Validate

The use of binding and Validate refers
to [Binding and Validate](https://www.cloudwego.io/docs/hertz/tutorials/basic-feature/binding-and-validate/).

## Ent

ent - An Entity Framework For Go.

This demo use `Ent` to operate `MySQL` and refers to [Ent](https://github.com/ent/ent).

#### Quick Start

- Update the Database DSN to your own in [Database init file](biz/dal/mysql/init.go).
- Go to the root directory of your project, and run following command, will generate the schema for User under biz/model/ent/schema/ directory:
```bash
  cd biz/model
  go run -mod=mod entgo.io/ent/cmd/ent init User
  ```
- Add fields to the User schema, run go generate to produce the ent operation files
```bash
  go generate ./ent
  ```
- For more Ent usage, please refer to [Ent Guides](https://entgo.io/).

## How to run

### Run mysql docker

```bash
cd bizdemo/hertz_ent && docker-compose up
```

### Run demo

```
cd bizdemo/hertz_ent
go build -o hertz_ent &&./hertz_ent
```