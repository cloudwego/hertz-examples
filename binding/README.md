# Bind and Validate
## How to use
[binding usage](https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/binding-and-validate/)

## How to run
1. `go run binding/main.go`
2. You can use the hertz client to send requests    
   `go run client/add_parameters/main.go`   
        
    Or you can use curl:
    ```shell
    curl --location --request POST '127.0.0.1:8080/v2/bind?query=hello&q=q1&q=q2&vd=1' \
    --header 'header: header' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "json":"hello,hertz"
    }'
    ```