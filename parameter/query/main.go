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

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to url matching: /welcome?firstname=Jane&lastname=Doe&food=apple&food=fish
	h.GET("/welcome", func(ctx context.Context, c *app.RequestContext) {
		firstname := c.DefaultQuery("firstname", "Guest")
		// shortcut for c.Request.URL.Query().Get("lastname")
		lastname := c.Query("lastname")

		// Iterate all queries and store the one with meeting the conditions in favoriteFood
		var favoriteFood []string
		c.QueryArgs().VisitAll(func(key, value []byte) {
			if string(key) == "food" {
				favoriteFood = append(favoriteFood, string(value))
			}
		})

		c.String(consts.StatusOK, "Hello %s %s, favorite food: %s", firstname, lastname, favoriteFood)
	})

	h.Spin()
}
