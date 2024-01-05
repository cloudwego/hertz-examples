/*
 * Copyright 2023 CloudWeGo Authors
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

package service

import (
	"context"
	"fmt"

	example "cwgo/example/hex/kitex_gen/hello/example"

	"github.com/cloudwego/kitex/pkg/klog"
)

type HelloMethodService struct {
	ctx context.Context
} // NewHelloMethodService new HelloMethodService
func NewHelloMethodService(ctx context.Context) *HelloMethodService {
	return &HelloMethodService{ctx: ctx}
}

// Run create note info
func (s *HelloMethodService) Run(request *example.HelloReq) (resp *example.HelloResp, err error) {
	// Finish your business logic.
	resp = new(example.HelloResp)
	resp.RespBody = fmt.Sprintf("[KITEX] hello, %s", request.Name)
	klog.Infof("[KITEX] hello, %s", request.Name)
	return
}
