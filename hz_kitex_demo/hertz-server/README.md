# Hertz Server
How to use hz to generate hertz scaffolding and develop services.
* Generate hertz code scaffolding with the hz new command: `hz new --idl=../idl/student_api.thrift`
* Generate kitex client code scaffolding with the kitex command: `kitex ../idl/student_management.thrift`
* Initialization of business logic and registration of global middleware in `main.go`
* Complete the business logic in `biz/handler/api/student_api.go` and make the rpc calls in it using the kitex client.
* Start server: `go run .`
* Test the interface with packet sending tools such as 'postman'.
* Update idl, add interface.
* Use the hz update command to synchronize idl updates to the code.
* Complete new business logic.
* Start server: `go run .`
* Test the interface with packet sending tools such as 'postman'.