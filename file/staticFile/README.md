## How to run

* navigate to current directory  
  `cd file/staticFile`
* start staticFile server  
  `go run main.go`

## FS

```go
type FS struct {
    noCopy nocopy.NoCopy
    Root string
    IndexNames []string
    GenerateIndexPages bool
    Compress bool
    AcceptByteRange bool
    PathRewrite PathRewriteFunc
    PathNotFound HandlerFunc
    CacheDuration time.Duration
    CompressedFileSuffix string
    once sync.Once
	h    HandlerFunc
}
```

* `Root`: Path to the root directory to serve files from.
  * use your browser to access "http://127.0.0.1:8080/static/1.txt", you will get `./static/1.txt`
* `IndexNames`: set indexNames, when you access the directory, you will get one of the file in the slice.
  * use your browser to access "http://127.0.0.1:8080/static1/txt/"
* `GenerateIndexPages`: set GenerateIndexPages true, when you access the directory, you will get the index(without `IndexNames`).
  * use your browser to access "http://127.0.0.1:8080/static2/txts"
* `Compress`: set the compress true, server adds a `CompressedFileSuffix` suffix to the original file name.
* `CompressedFileSuffix`: server attempts to save the resulting compressed file under the new file name.
* `AcceptByteRange`: enables clients to request a specific range of bytes from a file on the server.
* `PathRewrite`: use `NewPathSlashesStripper` to set this filed, it returns path rewriter, which strips slashesCount.
  * use your browser to access "http://127.0.0.1:8080/static1/txt/2.txt"
* `PathNotFound`: `PathNotFound` fires when file is not found in filesystem
  * use your browser to access "http://127.0.0.1:8080/static1/hertz"
* `CacheDuration`: set the time interval for automatically closing inactive file handlers.
