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
	"log"
	"net/http"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	hertzSentinel "github.com/hertz-contrib/opensergo/sentinel"
)

func initSentinel() {
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatalf("Unexpected error: %+v", err)
	}
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "client_test",
			Threshold:              0.0,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
			StatIntervalInMs:       1000,
		},
	})
	if err != nil {
		log.Fatalf("Unexpected error: %+v", err)
		return
	}
}

func main() {
	initSentinel()
	c, err := client.NewClient()
	if err != nil {
		log.Fatalf("Unexpected error: %+v", err)
		return
	}
	c.Use(hertzSentinel.SentinelClientMiddleware(
		// customize resource extractor if required
		// method_path by default
		hertzSentinel.WithClientResourceExtractor(func(ctx context.Context,
			request *protocol.Request, response *protocol.Response,
		) string {
			return "client_test"
		}),
		// customize block fallback if required
		// abort with status 429 by default
		hertzSentinel.WithClientBlockFallback(func(ctx context.Context, req *protocol.Request,
			resp *protocol.Response, blockError error,
		) error {
			resp.SetStatusCode(http.StatusBadRequest)
			resp.SetBody([]byte("request failed"))
			return blockError
		}),
	))

	req := &protocol.Request{}
	res := &protocol.Response{}
	req.SetMethod(consts.MethodGet)
	req.SetRequestURI("http://127.0.0.1:8081/client_test")
	err = c.Do(context.Background(), req, res)
	fmt.Printf("response body: %v, code: %v\n", string(res.Body()), res.StatusCode())
	fmt.Printf("error: %v", err)
}
