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
	"time"

	"github.com/cloudwego/hertz/pkg/protocol"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	h.GET("/cookie", func(ctx context.Context, c *app.RequestContext) {
		mc := "myCookie"
		// get specific key
		val := c.Cookie(mc)
		if val == nil {
			// set a cookie
			fmt.Printf("There is no cookie named: %s, and make one...\n", mc)
			cookie := protocol.AcquireCookie()
			cookie.SetKey("myCookie")
			cookie.SetValue("a nice cookie!")
			cookie.SetExpire(time.Now().Add(3600 * time.Second))
			cookie.SetPath("/")
			cookie.SetHTTPOnly(true)
			cookie.SetSecure(false)
			c.Response.Header.SetCookie(cookie)
			protocol.ReleaseCookie(cookie)
			c.WriteString("A cookie is ready.")
			return
		}

		fmt.Printf("Got a cookie: %s\nAnd eat it!", val)
		// instruct upload_file to delete a cookie
		// DelClientCookie instructs the upload_file to remove the given cookie.
		// This doesn't work for a cookie with specific domain or path,
		// you should delete it manually like:
		//
		//      c := AcquireCookie()
		//      c.SetKey(mc)
		//      c.SetDomain("example.com")
		//      c.SetPath("/path")
		//      c.SetExpire(CookieExpireDelete)
		//      h.SetCookie(c)
		//      ReleaseCookie(c)
		//
		c.Response.Header.DelClientCookie(mc)

		// construct the full struct of a cookie in response's header
		respCookie := protocol.AcquireCookie()
		respCookie.SetKey(mc)
		c.Response.Header.Cookie(respCookie)
		fmt.Printf("(The expire time of cookie is set to: %s)\n", respCookie.Expire())
		protocol.ReleaseCookie(respCookie)
		c.WriteString("The cookie has been eaten.")
	})

	h.Spin()
}
