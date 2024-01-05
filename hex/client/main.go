/*
 * Copyright 2023 CloudWeGo Authors
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
	"cwgo/example/hex/kitex_gen/hello/example"
	"cwgo/example/hex/kitex_gen/hello/example/helloservice"
	"fmt"

	"github.com/cloudwego/kitex/client"
)

func main() {
	kc, err := helloservice.NewClient("p.s.m", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		panic(err)
	}
	req := &example.HelloReq{Name: "hex"}
	resp, err := kc.HelloMethod(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
