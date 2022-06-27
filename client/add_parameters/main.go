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
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func realURL(u string) string {
	encode := url.Values{}
	encode.Add("query", "query")
	encode.Add("q", "q1")
	encode.Add("q", "q2")
	encode.Add("vd", "1")

	return fmt.Sprintf("%s?%s", u, encode.Encode())
}

func main() {
	client, err := client.NewClient()
	if err != nil {
		return
	}
	req := &protocol.Request{}
	res := &protocol.Response{}

	// Construct your own request URL using the "net/url" library
	req.Header.SetMethod(consts.MethodPost)
	req.SetRequestURI(realURL("http://127.0.0.1:8080/v1/bind"))
	err = client.Do(context.Background(), req, res)
	if err != nil {
		return
	}

	// Use SetQueryString to set query parameters
	req.Reset()
	req.Header.SetMethod(consts.MethodPost)
	req.SetRequestURI("http://127.0.0.1:8080/v1/bind")
	req.SetQueryString("query=query&q=q1&q=q2&vd=1")
	err = client.Do(context.Background(), req, res)
	if err != nil {
		return
	}

	// Send "www-url-encoded" request
	req.Reset()
	req.Header.SetMethod(consts.MethodPost)
	req.SetRequestURI("http://127.0.0.1:8080/v1/bind?query=query&q=q1&q=q2&vd=1")
	req.SetFormData(map[string]string{
		"form": "test form",
	})
	err = client.Do(context.Background(), req, res)
	if err != nil {
		return
	}

	// Send "multipart/form-data" request
	req.Reset()
	req.Header.SetMethod(consts.MethodPost)
	req.SetRequestURI("http://127.0.0.1:8080/v1/bind?query=query&q=q1&q=q2&vd=1")
	req.SetMultipartFormData(map[string]string{
		"form": "test form",
	})
	err = client.Do(context.Background(), req, res)
	if err != nil {
		return
	}

	// Send "Json" request
	req.Reset()
	req.Header.SetMethod(consts.MethodPost)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	req.SetRequestURI("http://127.0.0.1:8080/v1/bind?query=query&q=q1&q=q2&vd=1")
	data := struct {
		Json string `json:"json"`
	}{
		"test json",
	}
	jsonByte, _ := json.Marshal(data)
	req.SetBody(jsonByte)
	err = client.Do(context.Background(), req, res)
	if err != nil {
		return
	}
}
