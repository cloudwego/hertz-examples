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
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/sse"
)

func main() {
	// 创建 Hertz 服务器实例
	h := server.Default(
		server.WithHostPorts(":8080"),
		server.WithSenseClientDisconnection(true),
	)

	// 设置路由
	h.GET("/sse", sseHandler)

	// 启动服务器
	h.Spin()
}

// sseHandler 处理 SSE 请求
func sseHandler(ctx context.Context, c *app.RequestContext) {
	// 获取上次事件 ID（如果有）
	lastEventID := sse.GetLastEventID(&c.Request)
	fmt.Printf("Server Got LastEventID: %s\n", lastEventID)

	// 创建 SSE 写入器
	w := sse.NewWriter(c)

	// 创建一个通道用于检测客户端断开连接
	connClosed := c.Finished()

	// 使用 goroutine 监听连接状态
	go func() {
		<-connClosed
		fmt.Println("客户端连接已断开")
	}()

	// 发送 10 个事件，每个事件间隔 1 秒
	for i := 0; i < 10; i++ {
		// 检查连接是否已关闭
		select {
		case <-connClosed:
			fmt.Println("检测到客户端已断开连接，停止发送事件")
			return
		default:
			// 继续发送事件
		}

		// 创建事件 ID
		id := fmt.Sprintf("id-%d", i)

		// 创建事件数据
		data := fmt.Sprintf("Event data: %d at %s", i, time.Now().Format(time.RFC3339))

		fmt.Println("准备发送事件:", id, data)
		// 写入事件
		err := w.WriteEvent(id, "message", []byte(data))
		if err != nil {
			fmt.Printf("Error writing event: %v\n", err)
			break
		}

		// 等待 1 秒
		time.Sleep(5 * time.Second)
	}

	// 关闭 SSE 写入器
	w.Close()
	fmt.Println("SSE 事件发送完成")
}
