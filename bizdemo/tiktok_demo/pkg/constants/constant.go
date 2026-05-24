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

package constants

import (
	"fmt"
	"os"
)

// GetMySQLDSN retrieves the MySQL DSN from env, with fatal exit if unset.
func GetMySQLDSN() string {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		fmt.Fprintf(os.Stderr, "fatal: DB_DSN is not set\n")
		os.Exit(1)
	}
	return dsn
}

// GetMinioEndpoint retrieves the MinIO endpoint from env, with fatal exit if unset.
func GetMinioEndpoint() string {
	ep := os.Getenv("MINIO_ENDPOINT")
	if ep == "" {
		fmt.Fprintf(os.Stderr, "fatal: MINIO_ENDPOINT is not set\n")
		os.Exit(1)
	}
	return ep
}

// GetMinioAccessKeyID retrieves the MinIO access key from env, with fatal exit if unset.
func GetMinioAccessKeyID() string {
	key := os.Getenv("MINIO_ACCESS_KEY_ID")
	if key == "" {
		fmt.Fprintf(os.Stderr, "fatal: MINIO_ACCESS_KEY_ID is not set\n")
		os.Exit(1)
	}
	return key
}

// GetMinioSecretAccessKey retrieves the MinIO secret key from env, with fatal exit if unset.
func GetMinioSecretAccessKey() string {
	key := os.Getenv("MINIO_SECRET_ACCESS_KEY")
	if key == "" {
		fmt.Fprintf(os.Stderr, "fatal: MINIO_SECRET_ACCESS_KEY is not set\n")
		os.Exit(1)
	}
	return key
}

// GetRedisAddr retrieves the Redis address from env, with fatal exit if unset.
func GetRedisAddr() string {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		fmt.Fprintf(os.Stderr, "fatal: REDIS_ADDR is not set\n")
		os.Exit(1)
	}
	return addr
}

// GetRedisPassword retrieves the Redis password from env, with fatal exit if unset.
func GetRedisPassword() string {
	return os.Getenv("REDIS_PASSWORD")
}

// connection information
const (
	MiniouseSSL = false
)

// constants in the project
const (
	UserTableName      = "users"
	FollowsTableName   = "follows"
	VideosTableName    = "videos"
	MessageTableName   = "messages"
	FavoritesTableName = "likes"
	CommentTableName   = "comments"

	VideoFeedCount       = 30
	FavoriteActionType   = 1
	UnFavoriteActionType = 2

	MinioVideoBucketName = "videobucket"
	MinioImgBucketName   = "imagebucket"

	TestSign       = "测试账号！ offer"
	TestAva        = "avatar/test1.jpg"
	TestBackground = "background/test1.png"
)
