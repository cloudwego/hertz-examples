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

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	Do()
	DoDeadline()
	DoRedirects()
	DoTimeout()
	Get()
	GetDeadline()
	GetTimeout()
	Post()
}

func Do() {
	c, err := client.NewClient()
	if err != nil {
		return
	}
	req := &protocol.Request{}
	res := &protocol.Response{}
	req.SetMethod(consts.MethodGet)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	req.SetRequestURI("http://www.example.com")
	err = c.Do(context.Background(), req, res)
	if err != nil {
		return
	}
	fmt.Printf("%v", string(res.Body()))
}

func DoDeadline() {
	c, err := client.NewClient()
	if err != nil {
		return
	}
	req := &protocol.Request{}
	res := &protocol.Response{}
	req.SetMethod(consts.MethodGet)
	req.SetRequestURI("http://www.example.com")
	c.DoDeadline(context.Background(), req, res, time.Now().Add(1*time.Second))
	fmt.Printf("%v\n", string(res.Body()))
}

func DoRedirects() {
	c, err := client.NewClient()
	if err != nil {
		return
	}
	req := &protocol.Request{}
	res := &protocol.Response{}
	req.SetMethod(consts.MethodGet)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	req.SetRequestURI("http://www.example.com")
	err = c.DoRedirects(context.Background(), req, res, 1)
	if err != nil {
		return
	}
	fmt.Printf("%v\n", string(res.Body()))
}

func DoTimeout() {
	c, err := client.NewClient()
	if err != nil {
		return
	}
	req := &protocol.Request{}
	res := &protocol.Response{}
	req.SetMethod(consts.MethodGet)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	req.SetRequestURI("http://www.example.com")
	err = c.DoTimeout(context.Background(), req, res, 1*time.Second)
	if err != nil {
		return
	}
	fmt.Printf("%v\n", string(res.Body()))
}

func Get() {
	c, err := client.NewClient()
	if err != nil {
		return
	}
	status, body, _ := c.Get(context.Background(), nil, "http://www.example.com")
	fmt.Printf("status=%v body=%v\n", status, string(body))
}

func GetDeadline() {
	c, err := client.NewClient()
	if err != nil {
		return
	}
	status, body, _ := c.GetDeadline(context.Background(), nil, "http://www.example.com", time.Now().Add(1*time.Second))
	fmt.Printf("status=%v body=%v\n", status, string(body))
}

func GetTimeout() {
	c, err := client.NewClient()
	if err != nil {
		return
	}
	status, body, _ := c.GetTimeout(context.Background(), nil, "http://www.example.com", 1*time.Second)
	fmt.Printf("status=%v body=%v\n", status, string(body))
}

func Post() {
	c, err := client.NewClient()
	if err != nil {
		return
	}

	var postArgs protocol.Args
	postArgs.Set("arg", "a") // Set post args
	status, body, _ := c.Post(context.Background(), nil, "http://www.example.com", &postArgs)
	fmt.Printf("status=%v body=%v\n", status, string(body))
}
