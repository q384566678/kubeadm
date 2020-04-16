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

package commands

import (
	"testing"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	operatorv1 "k8s.io/kubeadm/operator/api/v1alpha1"
)

func TestRunCommand(t *testing.T) {
	tests := []struct {
		name   string
		c      *operatorv1.CommandDescriptor
		log    logr.Logger
		expect error
	}{
		{
			name:   "c = nil",
			c:      nil,
			log:    nil,
			expect: errors.New("invalid Task.Spec.[]CommandDescriptor. There are no command implementations matching this spec"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := RunCommand(test.c, test.log); err.Error() != test.expect.Error() {
				t.Errorf("expect = %v, got = %v", test.expect, err)
			}
		})
	}
}
