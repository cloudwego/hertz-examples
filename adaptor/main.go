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
	"context"
	"fmt"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	h := server.Default()

	h.GET("/hello", func(ctx context.Context, c *app.RequestContext) {
		req, err := adaptor.GetCompatRequest(&c.Request)
		if err != nil {
			fmt.Println(err)
			return
		}
		// You may build more logic on req
		fmt.Println(req.URL.String())

		// caution: don't pass in c.GetResponse() as it return a copy of response
		rw := adaptor.GetCompatResponseWriter(&c.Response)

		handler(rw, req)
	})

	h.Spin()
}
