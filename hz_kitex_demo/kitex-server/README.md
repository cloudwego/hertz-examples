# Kitex Server
How to use kitex to generate kitex server scaffolding and develop services.
* Generate kitex server code scaffolding with the kitex command: `kitex -service student ../idl/student_management.thrift`
* Complete the business logic in `handler.go`.
* Start server: `go run .`