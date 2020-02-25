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

	operatorv1 "k8s.io/kubeadm/operator/api/v1alpha1"
)

func TestplanCustom(t *testing.T) {
	tests := []struct {
		name      string
		operation *operatorv1.Operation
		spec      *operatorv1.CustomOperationSpec
		expect    int
	}{
		{
			name:      "WorkFlow = 0",
			operation: &operatorv1.Operation{},
			spec:      &operatorv1.CustomOperationSpec{},
			expect:    0,
		},
		{
			name:      "WorkFlow = 3",
			operation: &operatorv1.Operation{},
			spec: &operatorv1.CustomOperationSpec{
				Workflow: []operatorv1.RuntimeTaskGroup{
					operatorv1.RuntimeTaskGroup{},
					operatorv1.RuntimeTaskGroup{},
					operatorv1.RuntimeTaskGroup{},
				},
			},
			expect: 3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if res := planCustom(test.operation, test.spec); len(res.Items) != test.expect {
				t.Errorf("expect: %d, got: %d", test.expect, len(res.Items))
			}
		})
	}
}
