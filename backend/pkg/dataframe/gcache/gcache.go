// Copyright © 2022 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
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

package gcache

import (
	"time"

	"github.com/bluele/gcache"
)

type DataFrame struct {
	backend gcache.Cache
	ttl     time.Duration
}

func (df *DataFrame) Init(ttl time.Duration) error {
	df.ttl = ttl
	backend := gcache.New(20).
		LRU().
		Build()
	df.backend = backend

	return nil
}

func (df *DataFrame) Set(key string, value interface{}) bool {
	df.backend.SetWithExpire(key, value, df.ttl)

	return true
}

func (df *DataFrame) Get(key string) (interface{}, bool) {
	result, err := df.backend.Get(key)
	if err != nil {
		return nil, false
	}

	return result, true
}

func (df *DataFrame) Del(key string) {
	df.backend.Remove(key)
}
