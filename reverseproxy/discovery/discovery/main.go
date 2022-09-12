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
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/hertz-contrib/registry/nacos"
	"github.com/hertz-contrib/reverseproxy"
)

func main() {
	cli, err := client.NewClient()
	if err != nil {
		panic(err)
	}
	r, err := nacos.NewDefaultNacosResolver()
	if err != nil {
		panic(err)
	}
	cli.Use(sd.Discovery(r))
	h := server.New(server.WithHostPorts(":8741"))
	proxy, _ := reverseproxy.NewSingleHostReverseProxy("http://demo.hertz-contrib.reverseproxy")
	proxy.SetClient(cli)
	proxy.SetDirector(func(req *protocol.Request) {
		req.SetRequestURI(string(reverseproxy.JoinURLPath(req, proxy.Target)))
		req.Header.SetHostBytes(req.URI().Host())
		req.Options().Apply([]config.RequestOption{config.WithSD(true)})
	})
	h.GET("/backend", proxy.ServeHTTP)
	h.Spin()
}
