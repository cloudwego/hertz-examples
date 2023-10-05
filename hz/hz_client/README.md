# hz client generate

## install
```
go install github.com/cloudwego/hertz/cmd/hz@v0.5.0
```

## generate(Ignorable, can run code directly)
- server<br>
```
cd server
hz new --idl=../idl/psm.thrift --handler_by_method -t=template=slim
```
- client<br>
```
cd client
hz client --idl=../idl/psm.thrift --model_dir=hertz_gen -t=template=slim --client_dir=hz_client
```

## Create the `main.go` file in the client directory
```
// client/main.go
package main

import (
	"context"
	"fmt"
	"hertz-examples/hz/hz_client/client/hertz_gen/toutiao/middleware/hzClient/hertz"
	"time"

	hzClient "hertz-examples/hz/hz_client/client/hertz_gen/toutiao/middleware/hzClient"
)

func main() {
	idlCli, err := hertz.NewHertzClient("http://127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}

	{
		// query method
		queryReq := hzClient.QueryReq{
			QueryValue: "hello,query",
		}
		resp, rawResp, err := idlCli.QueryMethod(context.Background(), &queryReq)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp)
		fmt.Println(string(rawResp.Body()))
	}

	time.Sleep(500 * time.Millisecond)

	{
		// form method
		formReq := hzClient.FormReq{
			FormValue: "hello, form",
			FileValue: "./main.go",
		}
		resp, rawResp, err := idlCli.FormMethod(context.Background(), &formReq)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp)
		fmt.Println(string(rawResp.Body()))
	}

	time.Sleep(500 * time.Millisecond)

	{
		// path method
		pathReq := hzClient.PathReq{
			PathValue: "helloPath",
		}
		resp, rawResp, err := idlCli.PathMethod(context.Background(), &pathReq)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp)
		fmt.Println(string(rawResp.Body()))
	}

	time.Sleep(500 * time.Millisecond)

	{
		// body method
		bodyReq := hzClient.BodyReq{
			BodyValue:  "hello, body",
			QueryValue: "hello, query & body",
		}
		resp, rawResp, err := idlCli.BodyMethod(context.Background(), &bodyReq)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp)
		fmt.Println(string(rawResp.Body()))
	}
}
```
## run
- server<br>
```
cd server
go run .
```

- client<br>
```
cd client
go run .
```
