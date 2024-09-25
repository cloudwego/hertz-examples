# hertz_swagger_gen

## Introduction

An example using `Hertz` and `thrift-gen-http-swagger`.

- Defines `HTTP` interfaces using `thrift` IDL
- Generates code using `hz`
- Uses `Gorm` and `MySQL`
- Generates `swagger` files and `swagger ui` service using `thrift-gen-http-swagger` plugin

- `/swagger` provides `swagger` files and a `swagger ui` server
- `/handler.go` contains the basic business logic for updating, adding, deleting, and querying users

## IDL

This example defines `HTTP` interfaces using `thrift` IDL. The specific interface definitions can be found in [user.thrift](idl/user.thrift).

## Code Generation Tool

This example uses `hz` to generate code. For more details on how to use `hz`, refer to the official [hz documentation](https://www.cloudwego.io/docs/hertz/tutorials/toolkit/toolkit/).

The `hz` commands used can be found in the [Makefile](Makefile).

## Binding and Validation

For details on binding and validation, refer to [Binding and Validate](https://www.cloudwego.io/docs/hertz/tutorials/basic-feature/binding-and-validate/).

## Plugin

`thrift-gen-http-swagger` generates `swagger` documentation and `swagger ui` service through code generation.

For more information on using the plugin, refer to [swagger-generate](https://github.com/hertz-contrib/swagger-generate).

## How to Run

### Run MySQL Docker

```bash
cd bizdemo/hertz_swagger_gen && docker-compose up
```

### Run Example

```bash
cd bizdemo/hertz_swagger_gen
go run .
```