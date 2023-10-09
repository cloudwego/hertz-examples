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

type BindError struct {
	ErrType, FailField, Msg string
}

// Error implements error interface.
func (e *BindError) Error() string {
	if e.Msg != "" {
		return e.ErrType + ": expr_path=" + e.FailField + ", cause=" + e.Msg
	}
	return e.ErrType + ": expr_path=" + e.FailField + ", cause=invalid"
}

type ValidateError struct {
	ErrType, FailField, Msg string
}

// Error implements error interface.
func (e *ValidateError) Error() string {
	if e.Msg != "" {
		return e.ErrType + ": expr_path=" + e.FailField + ", cause=" + e.Msg
	}
	return e.ErrType + ": expr_path=" + e.FailField + ", cause=invalid"
}

func main() {
	CustomValidateErrFunc := func(failField, msg string) error {
		err := ValidateError{
			ErrType:   "validateErr",
			FailField: "[validateFailField]: " + failField,
			Msg:       "[validateErrMsg]: " + msg,
		}

		return &err
	}
	validateConfig := binding.NewValidateConfig()
	validateConfig.SetValidatorErrorFactory(CustomValidateErrFunc)
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"),
		server.WithValidateConfig(validateConfig))

	h.GET("bindErr", func(ctx context.Context, c *app.RequestContext) {
		type TestBind struct {
			A string `query:"a,required"`
		}
		var req TestBind
		err := c.Bind(&req)
		fmt.Printf("error: %v\n", err)
	})

	h.GET("validateErr", func(ctx context.Context, c *app.RequestContext) {
		type TestValidate struct {
			B int `query:"b" vd:"$>100; msg:'C must greater than 100'"`
		}
		var req TestValidate
		err := c.Bind(&req)
		if err != nil {
			panic(err)
		}
		err = c.Validate(&req)
		fmt.Printf("error: %v\n", err)
	})

	go h.Spin()

	time.Sleep(1000 * time.Millisecond)
	c, _ := client.NewClient()
	req := protocol.Request{}
	resp := protocol.Response{}
	req.SetMethod(consts.MethodGet)
	req.SetRequestURI("http://127.0.0.1:8080/bindErr")
	c.Do(context.Background(), &req, &resp)

	req.SetRequestURI("http://127.0.0.1:8080/validateErr?b=1")
	c.Do(context.Background(), &req, &resp)
}
