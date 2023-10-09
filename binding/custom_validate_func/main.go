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

type ValidateStruct struct {
	A string `query:"a" vd:"test($)"`
}

func main() {
	validateConfig := binding.NewValidateConfig()
	validateConfig.MustRegValidateFunc("test", func(args ...interface{}) error {
		if len(args) != 1 {
			return fmt.Errorf("the args must be one")
		}
		s, _ := args[0].(string)
		if s == "123" {
			return fmt.Errorf("the args can not be 123")
		}
		return nil
	})
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	h.GET("customValidate", func(ctx context.Context, c *app.RequestContext) {
		var req ValidateStruct
		err := c.Bind(&req)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		err = c.Validate(&req)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	})

	go h.Spin()

	time.Sleep(1000 * time.Millisecond)
	c, _ := client.NewClient()
	req := protocol.Request{}
	resp := protocol.Response{}
	req.SetMethod(consts.MethodGet)
	req.SetRequestURI("http://127.0.0.1:8080/customValidate?a=123")
	c.Do(context.Background(), &req, &resp)
}
