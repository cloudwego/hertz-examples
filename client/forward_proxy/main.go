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
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/cloudwego/hertz/pkg/protocol"
)

func main() {
	// Proxy address
	proxyURL := "http://<__user_name__>:<__password__>@<__proxy_addr__>:<__proxy_port__>"

	parsedProxyURL := protocol.ParseURI(proxyURL)
	client, err := client.NewClient(client.WithDialer(standard.NewDialer()))
	if err != nil {
		return
	}
	client.SetProxy(protocol.ProxyURI(parsedProxyURL))
	upstreamURL := "http://google.com"
	_, body, _ := client.Get(context.Background(), nil, upstreamURL)
	fmt.Println(string(body))
}
