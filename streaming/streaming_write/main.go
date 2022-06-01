/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"bytes"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"), server.WithStreamBody(true))

	h.GET("/streamWrite", handler1)

	// Demo: synchronized reading and writing
	h.GET("/syncWrite", handler2)

	h.Spin()
}

func handler1(ctx context.Context, c *app.RequestContext) {
	bs := []byte("hello, hertz!")
	wb := bytes.NewBuffer(bs)
	c.SetBodyStream(wb, len(bs))
}

func handler2(ctx context.Context, c *app.RequestContext) {
	rw := newChunkReader()
	// Content-Length may be negative:
	// -1 means Transfer-Encoding: chunked.
	// -2 means Transfer-Encoding: identity.
	c.SetBodyStream(rw, -1)

	go func() {
		for i := 1; i < 1000; i++ {
			// For each streaming_write, the upload_file prints
			rw.Write([]byte(fmt.Sprintf("===%d===\n", i)))
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}
		rw.Close()
	}()

	go func() {
		<-c.Finished()
		fmt.Println("request process end")
	}()
}

type ChunkReader struct {
	rw  bytes.Buffer
	w2r chan struct{}
	r2w chan struct{}
}

func newChunkReader() *ChunkReader {
	var rw bytes.Buffer
	w2r := make(chan struct{})
	r2w := make(chan struct{})
	cr := &ChunkReader{rw, w2r, r2w}
	return cr
}

var closeOnce = new(sync.Once)

func (cr *ChunkReader) Read(p []byte) (n int, err error) {
	for {
		_, ok := <-cr.w2r
		if !ok {
			closeOnce.Do(func() {
				close(cr.r2w)
			})
			n, err = cr.rw.Read(p)
			return
		}

		n, err = cr.rw.Read(p)

		cr.r2w <- struct{}{}

		if n == 0 {
			continue
		}
		return
	}
}

func (cr *ChunkReader) Write(p []byte) (n int, err error) {
	n, err = cr.rw.Write(p)
	cr.w2r <- struct{}{}
	<-cr.r2w
	return
}

func (cr *ChunkReader) Close() {
	close(cr.w2r)
}
