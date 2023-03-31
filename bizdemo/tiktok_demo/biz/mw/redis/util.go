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

package redis

import (
	"strconv"

	"github.com/go-redis/redis/v7"
)

// add k & v to redis
func add(c *redis.Client, k string, v int64) {
	tx := c.TxPipeline()
	tx.SAdd(k, v)
	tx.Expire(k, expireTime)
	tx.Exec()
}

// del k & v
func del(c *redis.Client, k string, v int64) {
	tx := c.TxPipeline()
	tx.SRem(k, v)
	tx.Expire(k, expireTime)
	tx.Exec()
}

// check the set of k if exist
func check(c *redis.Client, k string) bool {
	if e, _ := c.Exists(k).Result(); e > 0 {
		return true
	}
	return false
}

// exist check the relation k and v if exist
func exist(c *redis.Client, k string, v int64) bool {
	if e, _ := c.SIsMember(k, v).Result(); e {
		c.Expire(k, expireTime)
		return true
	}
	return false
}

// count get the size of the set of key
func count(c *redis.Client, k string) (sum int64, err error) {
	if sum, err = c.SCard(k).Result(); err == nil {
		c.Expire(k, expireTime)
		return sum, err
	}
	return sum, err
}

func get(c *redis.Client, k string) (vt []int64) {
	v, _ := c.SMembers(k).Result()
	c.Expire(k, expireTime)
	for _, vs := range v {
		v_i64, _ := strconv.ParseInt(vs, 10, 64)
		vt = append(vt, v_i64)
	}
	return vt
}
