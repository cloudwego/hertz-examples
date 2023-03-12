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

package minio

import (
	"context"
	"fmt"
	"offer_tiktok/pkg/constants"
	"testing"

	"github.com/minio/minio-go/v7"
)

func TestBucketExist(t *testing.T) {
	ctx := context.Background()
	exists, err := Client.BucketExists(ctx, constants.MinioVideoBucketName)
	if err != nil {
		fmt.Println(err)
		return
	}
	if exists {
		fmt.Printf("%v found!\n", constants.MinioVideoBucketName)
	} else {
		fmt.Println("not found!")
	}
}

func TestBuckMake(t *testing.T) {
	ctx := context.Background()
	exists, err := Client.BucketExists(ctx, constants.MinioVideoBucketName)
	if err != nil {
		fmt.Println(err)
		return
	}
	if exists {
		fmt.Printf("%v found!\n", constants.MinioVideoBucketName)
	} else {
		err = Client.MakeBucket(ctx, constants.MinioVideoBucketName, minio.MakeBucketOptions{})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Successfully created mybucket %v\n", constants.MinioVideoBucketName)
	}
}

func TestGetObjURL(t *testing.T) {
	Init()
	ctx := context.Background()
	url, _ := GetObjURL(ctx, constants.MinioVideoBucketName, "1000.1676403991.mp4")
	fmt.Println(url.String())
}
