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
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/registry/nacos"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		num := i
		go func() {
			addr := fmt.Sprintf("127.0.0.1:800%d", num)
			r, err := nacos.NewDefaultNacosRegistry()
			if err != nil {
				hlog.Fatal(err)
			}
			h := server.Default(
				server.WithHostPorts(addr),
				server.WithRegistry(r, &registry.Info{
					ServiceName: "hertz.test.demo",
					Addr:        utils.NewNetAddr("tcp", addr),
					Weight:      10,
					Tags:        nil,
				}),
			)
			h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
				ctx.JSON(consts.StatusOK, utils.H{"addr": addr})
			})
			h.Spin()
			wg.Done()
		}()
	}
	wg.Wait()
}
