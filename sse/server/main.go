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

		// 写入事件
		err := w.WriteEvent(id, "message", []byte(data))
		if err != nil {
			fmt.Printf("Error writing event: %v\n", err)
			fmt.Println("写入事件时发生错误，可能是客户端已断开连接")
			break
		}

		// 刷新缓冲区，确保事件立即发送
		if err := c.Flush(); err != nil {
			fmt.Printf("Error flushing buffer: %v\n", err)
			fmt.Println("刷新缓冲区时发生错误，可能是客户端已断开连接")
			break
		}

		// 等待 1 秒
		time.Sleep(1 * time.Second)
	}

	// 关闭 SSE 写入器
	w.Close()
	fmt.Println("SSE 事件发送完成")
}
