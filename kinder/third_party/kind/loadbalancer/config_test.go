/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package loadbalancer

import (
	"testing"

	"github.com/pkg/errors"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name   string
		data   *ConfigData
		expect error
	}{
		{
			name: "port >= 0 && IPv6 == true",
			data: &ConfigData{
				ControlPlanePort: 8888,
				IPv6:             true,
			},
			expect: nil,
		},
		{
			name: "port < 0 && IPv6 == true",
			data: &ConfigData{
				ControlPlanePort: -1,
				IPv6:             true,
			},
			expect: errors.New("isn't nil"),
		},
		{
			name: "port >= 0 && IPv6 == false",
			data: &ConfigData{
				ControlPlanePort: 8888,
				IPv6:             false,
			},
			expect: nil,
		},
		{
			name: "port < 0 && IPv6 == false",
			data: &ConfigData{
				ControlPlanePort: -1,
				IPv6:             false,
			},
			expect: errors.New("isn't nil"),
		},
		{
			name:   "data == nil",
			data:   nil,
			expect: errors.New("isn't nil"),
		},
	}
	for _, test := range tests {
		t.Run((test.name), func(t *testing.T) {
			if _, err := Config(test.data); (nil == err) != (nil == test.expect) {
				t.Errorf("expect = %v, got = %v", test.expect, err)
			}
		})
	}
}
