package main

import (
	"context"
	"unsafe"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/hertz-contrib/reverseproxy"
)

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
func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:9998"))
	proxy, err := reverseproxy.NewSingleHostReverseProxy("http://127.0.0.1:9998/proxy")
	if err != nil {
		panic(err)
	}
	headers := map[string]string{"Key1": "value1"}

	proxy.SetModifyResponse(func(response *protocol.Response) error {
		response.Header.Set("Key2", "value2")
		return nil
	})

	h.GET("/backend", func(c context.Context, ctx *app.RequestContext) {
		// post ctx to reserveproxy
		ctx.Request.SetHeaders(headers)
		proxy.ServeHTTP(c, ctx)
		// get ctx from reserveproxy
		_ = ctx.Response.Header.Get("Key2")
	})
	h.Spin()

}

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
