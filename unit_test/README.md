# Unit Test

You can learn how to write unit tests without network transfer.   

We provide two interfaces `ResponseRecord` and `PerformRequest`
* ResponseRecord: It is used to record the response results, you can write serialized response data via `Write(buf []byte)`, and get the `hertz.Response` object via `Result()`
* PerformRequest: It is used to make a request to the specified engine with the specified url. You can provide a test engine and pass in the method, url, and optionally the Body and Header

For more information about unit test, please click [unit test](https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/unit-test/)
## How to run
  `go test unit_test/main_test.go`