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
	"time"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/client/loadbalance"
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	loadbalanceEx "github.com/hertz-contrib/loadbalance/round_robin"
	"github.com/hertz-contrib/registry/nacos"
)

func main() {
	cli, err := client.NewClient()
	if err != nil {
		hlog.Fatal(err)
	}
	lb := loadbalanceEx.NewRoundRobinBalancer()
	opt := loadbalance.Options{
		RefreshInterval: 10 * time.Second,
		ExpireInterval:  25 * time.Second,
	}
	r, err := nacos.NewDefaultNacosResolver()
	if err != nil {
		hlog.Fatal(err)
	}
	cli.Use(sd.Discovery(r, sd.WithLoadBalanceOptions(lb, opt)))
	for i := 0; i < 10; i++ {
		status, body, err := cli.Get(context.Background(), nil, "http://hertz.test.demo/ping")
		if err != nil {
			hlog.Fatal(err)
		}
		hlog.Infof("code=%d,body=%s\n", status, string(body))
	}
}
