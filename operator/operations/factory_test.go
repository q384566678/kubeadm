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

package operations

import (
	"testing"

	"github.com/pkg/errors"
	operatorv1 "k8s.io/kubeadm/operator/api/v1alpha1"
)

func TestDaemonSetNodeSelectorLabels(t *testing.T) {
	tests := []struct {
		name      string
		operation *operatorv1.Operation
		expect    error
	}{
		{
			name:      "operation == nil",
			operation: nil,
			expect:    errors.New("Invalid Operation.Spec.OperatorDescriptor. There are no operation implementation matching this spec"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := DaemonSetNodeSelectorLabels(tt.operation)
			if err.Error() != tt.expect.Error() {
				t.Errorf("expect = %v, got = %v", tt.expect, err)
			}
		})
	}
}

func TestTaskGroupList(t *testing.T) {
	tests := []struct {
		name      string
		operation *operatorv1.Operation
		expect    error
	}{
		{
			name:      "operation == nil",
			operation: nil,
			expect:    errors.New("Invalid Operation.Spec.OperatorDescriptor. There are no operation implementation matching this spec"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := TaskGroupList(tt.operation)
			if err.Error() != tt.expect.Error() {
				t.Errorf("expect = %v, got = %v", tt.expect, err)
			}
		})
	}
}
