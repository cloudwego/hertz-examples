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

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
)

func MyMiddleware(next client.Endpoint) client.Endpoint {
	return func(ctx context.Context, req *protocol.Request, resp *protocol.Response) (err error) {
		// pre-handle
		fmt.Println("pre-clientHandle")
		err = next(ctx, req, resp)
		if err != nil {
			return
		}
		// post-handle
		fmt.Println("post-clientHandle")
		return nil
	}
}

func main() {
	client, err := client.NewClient()
	if err != nil {
		return
	}
	client.Use(MyMiddleware)
	req := &protocol.Request{}
	res := &protocol.Response{}
	req.SetRequestURI("http://127.0.0.1:8080/middleware")
	err = client.Do(context.Background(), req, res)
	return
}
