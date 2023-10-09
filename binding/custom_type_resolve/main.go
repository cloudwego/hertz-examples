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
	"reflect"
	"time"

	"github.com/cloudwego/hertz/pkg/route/param"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Nested struct {
	B string
	C string
}

type TestBind struct {
	A Nested `query:"a"`
}

func main() {
	bindConfig := binding.NewBindConfig()
	bindConfig.MustRegTypeUnmarshal(reflect.TypeOf(Nested{}), func(req *protocol.Request, params param.Params, text string) (reflect.Value, error) {
		if text == "" {
			return reflect.ValueOf(Nested{}), nil
		}
		val := Nested{
			B: text[:5],
			C: text[5:],
		}
		// 此外，也可以利用 req, params 来获取其他参数进行参数绑定
		return reflect.ValueOf(val), nil
	})
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	h.GET("customType", func(ctx context.Context, c *app.RequestContext) {
		var req TestBind
		c.Bind(&req)
		fmt.Printf("req: %v\n", req)
	})

	go h.Spin()

	time.Sleep(1000 * time.Millisecond)
	c, _ := client.NewClient()
	req := protocol.Request{}
	resp := protocol.Response{}
	req.SetMethod(consts.MethodGet)
	req.SetRequestURI("http://127.0.0.1:8080/customType?a=hellohertz")
	c.Do(context.Background(), &req, &resp)
}
