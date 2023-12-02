// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/reverseproxy"
	"github.com/hertz-contrib/websocket"
)

var (
	proxyURL    = "ws://127.0.0.1:8080/ws"
	backendURL  = "ws://127.0.0.1:9090/backend"
	proxyAddr   = "127.0.0.1:8080"
	backendAddr = "127.0.0.1:9090"
)

func main() {
	// websocket reverse proxy
	wsrp := reverseproxy.NewWSReverseProxy(backendURL)
	ps := server.Default(server.WithHostPorts(proxyAddr))
	ps.GET("/ws", wsrp.ServeHTTP)
	go ps.Spin()

	time.Sleep(time.Second)

	go func() {
		// backend server
		bs := server.Default(server.WithHostPorts(backendAddr))
		bs.GET("/backend", func(ctx context.Context, c *app.RequestContext) {
			upgrader := &websocket.HertzUpgrader{}
			if err := upgrader.Upgrade(c, func(conn *websocket.Conn) {
				for {
					msgType, msg, err := conn.ReadMessage()
					if err != nil {
						hlog.Errorf("backend read message err: %v", err)
					}
					err = conn.WriteMessage(msgType, msg)
					if err != nil {
						hlog.Errorf("backend write message err: %v", err)
					}
				}
			}); err != nil {
				hlog.Errorf("upgrade error: %v", err)
				return
			}
		})
		bs.Spin()
	}()

	time.Sleep(time.Second)

	// client
	conn, _, err := reverseproxy.DefaultOptions.Dialer.Dial(proxyURL, make(http.Header))
	if err != nil {
		hlog.Errorf("client dial err: %v", err)
		return
	}

	time.Sleep(time.Second)
	var echoInput string
	for {
		fmt.Print("send: ")
		_, _ = fmt.Scanln(&echoInput)
		err = conn.WriteMessage(websocket.TextMessage, []byte(echoInput))
		if err != nil {
			hlog.Errorf("client write message err: %v", err)
		}
		_, echoOutput, err := conn.ReadMessage()
		if err != nil {
			hlog.Errorf("client read message err: %v", err)
		}
		fmt.Println("receive: " + string(echoOutput))
	}
}
