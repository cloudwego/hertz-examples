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

package consts

import (
	"fmt"
	"os"
)

// getSessionSecret retrieves the session secret from env, with fatal exit if unset.
func GetSessionSecret() []byte {
	key := os.Getenv("SESSION_SECRET")
	if key == "" {
		fmt.Fprintf(os.Stderr, "fatal: SESSION_SECRET is not set. Generate one with: openssl rand -base64 32\n")
		os.Exit(1)
	}
	return []byte(key)
}

// GetCSRFSecret retrieves the CSRF secret from env, with fatal exit if unset.
func GetCSRFSecret() string {
	key := os.Getenv("CSRF_SECRET")
	if key == "" {
		fmt.Fprintf(os.Stderr, "fatal: CSRF_SECRET is not set. Generate one with: openssl rand -base64 32\n")
		os.Exit(1)
	}
	return key
}

// GetMySQLDSN retrieves the MySQL DSN from env, with fatal exit if unset.
func GetMySQLDSN() string {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		fmt.Fprintf(os.Stderr, "fatal: DB_DSN is not set\n")
		os.Exit(1)
	}
	return dsn
}

// constants
const (
	TCP           = "tcp"
	UserTableName = "users"
	RedisAddr     = "127.0.0.1:6379"
	MaxIdleNum    = 10
	RedisPasswd   = ""
	CSRFKeyLookUp = "form:csrf"
	Username      = "username"
	HertzSession  = "HERTZ-SESSION"
)

// error msg
const (
	Success     = "success"
	RegisterErr = "user already exists"
	LoginErr    = "wrong username or password"
	PageErr     = "please login first"
	CSRFErr     = "csrf exception"
)
