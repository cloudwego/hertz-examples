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
	"crypto/tls"
	"fmt"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/hertz-contrib/reverseproxy"
)

var (
	wg        sync.WaitGroup
	tlsConfig *tls.Config
)

func main() {
	wg.Add(2)
	cert, err := tls.LoadX509KeyPair("tls/server.crt", "tls/server.key")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	tlsConfig = &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		},
	}

	tlsConfig.Certificates = append(tlsConfig.Certificates, cert)

	go func() {
		defer wg.Done()
		h := server.New(
			server.WithHostPorts(":8004"),
			server.WithTLS(tlsConfig),
			// if you want to use tls server. pls use go net
			server.WithTransport(standard.NewTransporter),
		)

		h.GET("/backend", func(cc context.Context, c *app.RequestContext) {
			c.JSON(200, utils.H{"msg": "pong"})
		})
		h.Spin()
	}()

	go func() {
		defer wg.Done()
		h := server.New(
			server.WithHostPorts(":8005"),
			server.WithTransport(standard.NewTransporter),
			server.WithTLS(tlsConfig),
		)
		proxy, err := reverseproxy.NewSingleHostReverseProxy("https://127.0.0.1:8004",
			client.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
			client.WithDialer(standard.NewDialer()),
		)
		if err != nil {
			panic(err)
		}
		h.GET("/backend", proxy.ServeHTTP)
		h.Spin()
	}()
	wg.Wait()
}
