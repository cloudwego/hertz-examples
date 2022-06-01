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
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	client, err := client.NewClient()
	if err != nil {
		return
	}
	req := &protocol.Request{}
	res := &protocol.Response{}
	req.SetMethod(consts.MethodPost)
	req.SetRequestURI("http://127.0.0.1:8080/singleFile")
	req.SetFile("file", "./client/upload_file/file1.txt")

	err = client.Do(context.Background(), req, res)
	fmt.Println(err, string(res.Body()))

	req.SetRequestURI("http://127.0.0.1:8080/multiFile")
	req.SetFile("file", "./client/upload_file/file2.txt")
	err = client.Do(context.Background(), req, res)
	fmt.Println(err, string(res.Body()))
	return
}
