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

package test

const userServiceAddr = "http://127.0.0.1:8888"

var m1 = map[string]string{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NSwidXNlcm5hbWUiOiJhZG1pbiIsInJpZHMiOlt7ImlkIjoxLCJuYW1lIjoiYWRtaW4ifSx7ImlkIjoyLCJuYW1lIjoicm9sZSJ9XSwiZXhwIjoxNjc3OTgzMTg5fQ.l9dJnbXyAIW-KqOXZ1jQzWwCWiyi84zLI8akWX2bfwM",
}

var rolem1 = map[string]string{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwidXNlcm5hbWUiOiJyb2xlX3VzZXIiLCJyaWRzIjpbeyJpZCI6MiwibmFtZSI6InJvbGUifV0sImV4cCI6MTY3Nzk4NzQ2MH0.9CDKrUldcxg4dX2vqNWCBfuy44f3G7uwUqah6FhZb5c",
}

var permissionm1 = map[string]string{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NywidXNlcm5hbWUiOiJwZXJtaXNzaW9uX3VzZXIiLCJyaWRzIjpbeyJpZCI6MywibmFtZSI6InBlcm1pc3Npb24ifV0sImV4cCI6MTY3Nzk4ODQzN30.86z5EECrvRZuIiLDHYjwccbu81LoP6zxIuwjykA0fSk",
}
var header []byte
