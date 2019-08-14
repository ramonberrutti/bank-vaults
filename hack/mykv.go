// Copyright © 2019 Banzai Cloud
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
	"fmt"

	"github.com/banzaicloud/bank-vaults/pkg/kv"
)

type myKV struct {
}

func New(config map[string]string) (kv.Service, error) {
	fmt.Printf("myKV New(%+v)\n", config)
	return &myKV{}, nil
}

func (h *myKV) Get(key string) ([]byte, error) {
	fmt.Println("myKV Get()")
	return nil, nil
}

func (h *myKV) Set(key string, value []byte) error {
	fmt.Println("myKV Set()")
	return nil
}
