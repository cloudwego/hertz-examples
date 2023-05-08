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
	"os"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	h := server.Default()

	// SetLevel sets the level of logs below which logs will not be output.
	hlog.SetLevel(hlog.LevelDebug)

	f, err := os.Create("hertz.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// SetOutput sets the output of default logger. By default, it is stderr.
	hlog.SetOutput(f)
	// if you want to output the log to the file and the stdout at the same time, you can use the following codes

	// fileWriter := io.MultiWriter(f, os.Stdout)
	// hlog.SetOutput(fileWriter)
	h.GET("/hello", func(ctx context.Context, c *app.RequestContext) {
		// it will be output
		hlog.Info("Hello, hertz")
		// it will not be output
		hlog.Trace("Hello, hertz")
		c.String(consts.StatusOK, "Hello hertz!")
	})

	h.Spin()
}
