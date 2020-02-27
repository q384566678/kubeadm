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
