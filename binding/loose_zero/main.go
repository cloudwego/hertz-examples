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
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	bindConfig := binding.NewBindConfig()
	bindConfig.LooseZeroMode = true
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"), server.WithBindConfig(bindConfig))

	h.GET("looseZero", func(ctx context.Context, c *app.RequestContext) {
		type Loose struct {
			A int `query:"a"`
		}
		var req Loose
		err := c.BindAndValidate(&req)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			panic(err)
		}
		fmt.Printf("req: %v\n", req)
	})

	go h.Spin()

	time.Sleep(1000 * time.Millisecond)
	c, _ := client.NewClient()
	req := protocol.Request{}
	resp := protocol.Response{}
	req.SetMethod(consts.MethodGet)
	req.SetRequestURI("http://127.0.0.1:8080/looseZero?a")
	c.Do(context.Background(), &req, &resp)
}
