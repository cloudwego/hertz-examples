## How to run
* navigate to current directory  
`cd file/staticFile`
* start staticFile server  
`go run main.go`
* use your browser to access "http://127.0.0.1:8080/main"
* use your browser to access "http://127.0.0.1:8080/static/1.txt", you will get `./static/1.txt`
* use your browser to access "http://127.0.0.1:8080/static1/txt/2.txt", you will get `./static/2.txt`
* use your browser to access "http://127.0.0.1:8080/static1/hertz", there is no resource,so you will receive a reminder
* use your browser to access "http://127.0.0.1:8080/static1/txt/", you will get `./txt/1.txt` or `./txt/2.txt`,because we set the `indexNames`,when we access the directory, we will get one of the file in the slice
* use your browser to access "http://127.0.0.1:8080/static2/txts", you will get the index pages for directories(without `IndexNames`)

* `PathRewrite`: use `NewPathSlashesStripper` to set this filed, it returns path rewriter, which strips slashesCount
* `CacheDuration`: set the time interval for automatically closing inactive file handlers
* `IndexNames`: set indexNames, when you access the directory, you will get one of the file in the slice
* `Compress`: set the compress true, server adds a `CompressedFileSuffix` suffix to the original file name
* `CompressedFileSuffix`: server attempts to save the resulting compressed file under the new file name
* `AcceptByteRange`: enables clients to request a specific range of bytes from a file on the server
* `GenerateIndexPages`: set GenerateIndexPages true, when you access the directory, you will get the index(without `IndexNames`)
