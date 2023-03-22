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

package utils

import (
	"errors"

	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/pkg/errno"
)

type BaseResp struct {
	StatusCode int32
	StatusMsg  string
}

// BuildBaseResp convert error and build BaseResp
func BuildBaseResp(err error) *BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

// baseResp build BaseResp from error
func baseResp(err errno.ErrNo) *BaseResp {
	return &BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
}
