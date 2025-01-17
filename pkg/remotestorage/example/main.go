// Copyright 2018 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"io/ioutil"
	"log"

	"github.com/etcd-io/dbtester/pkg/remotestorage"

	"go.uber.org/zap"
)

func main() {
	kbs, err := ioutil.ReadFile("key.json")
	if err != nil {
		log.Fatal(err)
	}
	u, err := remotestorage.NewGoogleCloudStorage(zap.NewExample(), kbs, "etcd-development")
	if err != nil {
		log.Fatal(err)
	}
	if err := u.UploadFile("dbtester-results", "agent.log", "dir/agent.log", remotestorage.WithContentType("text/plain")); err != nil {
		log.Fatal(err)
	}

	// upload directories
	// kbs, err := ioutil.ReadFile("key.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// u, err := remotestorage.NewGoogleCloudStorage(kbs, "my-project")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := u.UploadDir("test-bucket", "articles", "articles"); err != nil {
	// 	log.Fatal(err)
	// }
}
