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
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/sse"
)

func main() {
	// 创建 Hertz 客户端
	c, err := client.NewClient()
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		return
	}

	// 创建请求和响应对象
	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
	defer protocol.ReleaseRequest(req)
	defer protocol.ReleaseResponse(resp)

	// 设置请求 URI 和方法
	req.SetRequestURI("http://localhost:8080/sse")
	req.SetMethod("GET")

	// 设置请求头
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")

	// 添加 SSE Accept MIME 类型
	sse.AddAcceptMIME(req)

	// 可选：设置上次接收到的事件 ID
	// req.Header.Set(sse.LastEventIDHeader, "id-0")

	// 发送请求
	fmt.Println("Connecting to SSE server...")
	if err := c.Do(context.Background(), req, resp); err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}

	// 检查响应状态码
	if resp.StatusCode() != 200 {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode())
		return
	}

	// 创建 SSE 读取器
	r, err := sse.NewReader(resp)
	if err != nil {
		fmt.Printf("Error creating SSE reader: %v\n", err)
		return
	}
	defer r.Close()

	fmt.Println("Connected to SSE stream. Receiving events...")

	// 设置信号处理，以便可以优雅地退出
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// 创建上下文，用于取消 SSE 事件处理
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 使用 ForEach 方法迭代处理 SSE 事件
	err = r.ForEach(ctx, func(e *sse.Event) error {
		// 等待信号或错误
		select {
		case <-sigCh:
			fmt.Println("\nReceived interrupt signal. Exiting...")
			cancel()
			return nil
		default:
		}
		fmt.Printf("Event received:\n")
		fmt.Printf("  ID: %s\n", e.ID)
		fmt.Printf("  Type: %s\n", e.Type)
		fmt.Printf("  Data: %s\n\n", string(e.Data))
		return nil
	})

	if err != nil {
		fmt.Printf("error processing events: %v", err)
	} else {
		fmt.Println("All events processed successfully.")
	}

	time.Sleep(5 * time.Second)
}
