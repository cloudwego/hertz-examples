## How to run
* navigate to current directory  
`cd file/staticFile`
* start staticFile server  
`go run main.go`
* use your browser to access "http://127.0.0.1:8080/main"
* use your browser to access "http://127.0.0.1:8080/static/1.txt", you will get `./static/1.txt`
* use your browser to access "http://127.0.0.1:8080/static1/txt/2.txt", you will get `./static/2.txt`
* use your browser to access "http://127.0.0.1:8080/static1/hertz", there is no resource,so you will receive a reminder
* the `root` is the path to the root directory to serve files from. 
* the `CacheDuration` field Set the time interval for automatically closing inactive file handlers, the default value is `consts.FSHandlerCacheDuration`(10s)
* use your browser to access "http://127.0.0.1:8080/static1/txt/", you will get `./txt/1.txt` or `./txt/2.txt`,because we set the `indexNames`,when we access the directory,we will get one of the file in the slice
* if the `Compress` is true, when the server needs to return a compressed static file, it adds a `CompressedFileSuffix` suffix to the original file name and attempts to save the resulting compressed file under the new file name.
* if the `AcceptByteRange` is true, enables clients to request a specific range of bytes from a file on the server.
* use your browser to access "http://127.0.0.1:8080/static2/txts", you will get the index pages for directories(without `IndexNames`)
