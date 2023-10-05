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