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
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/hertz-contrib/reverseproxy"
)

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8000"))
	// modify response
	proxy, _ := reverseproxy.NewSingleHostReverseProxy("http://127.0.0.1:8000/proxy")
	proxy.SetModifyResponse(func(resp *protocol.Response) error {
		resp.SetStatusCode(200)
		resp.SetBodyRaw([]byte("change response success"))
		return nil
	})
	h.GET("/modifyResponse", proxy.ServeHTTP)
	h.Spin()
}
