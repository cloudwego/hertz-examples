/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"gopkg.in/yaml.v3"
)

func main() {
	h := server.Default(server.WithHostPorts(":8080"))

	h.GET("/someYAML", func(ctx context.Context, c *app.RequestContext) {
		c.Render(consts.StatusOK, YAML{Data: "hello,world"})
	})

	h.Spin()
}

type YAML struct {
	Data interface{}
}

var yamlContentType = "application/yaml; charset=utf-8"

func (r YAML) Render(resp *protocol.Response) error {
	writeContentType(resp, yamlContentType)
	yamlBytes, err := yaml.Marshal(r.Data)
	if err != nil {
		return err
	}

	resp.AppendBody(yamlBytes)

	return nil
}

func (r YAML) WriteContentType(w *protocol.Response) {
	writeContentType(w, yamlContentType)
}

func writeContentType(resp *protocol.Response, value string) {
	resp.Header.SetContentType(value)
}
