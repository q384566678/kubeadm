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

	"github.com/pkg/errors"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name   string
		c      *cmd
		expect error
	}{
		{
			name:   "commands: pwd",
			c:      newCmd("pwd"),
			expect: nil,
		},
		{
			name:   "commands: ls -l .",
			c:      newCmd("ls", "-l", "."),
			expect: nil,
		},
		{
			name:   "commands: ''",
			c:      newCmd(""),
			expect: errors.New("commands nil"),
		},
	}
	for _, test := range tests {
		t.Run((test.name), func(t *testing.T) {
			if err := test.c.Run(); (nil == err) != (nil == test.expect) {
				t.Errorf("expect = %v, got = %v", test.expect, err)
			}
		})
	}
}

func TestRunWithEcho(t *testing.T) {
	tests := []struct {
		name   string
		c      *cmd
		expect error
	}{
		{
			name:   "commands: pwd",
			c:      newCmd("pwd"),
			expect: nil,
		},
		{
			name:   "commands: ls -l .",
			c:      newCmd("ls", "-l", "."),
			expect: nil,
		},
		{
			name:   "commands: ''",
			c:      newCmd(""),
			expect: errors.New("commands nil"),
		},
	}
	for _, test := range tests {
		t.Run((test.name), func(t *testing.T) {
			if err := test.c.RunWithEcho(); (nil == err) != (nil == test.expect) {
				t.Errorf("expect = %v, got = %v", test.expect, err)
			}
		})
	}

}

func TestRunAndCapture(t *testing.T) {
	tests := []struct {
		name   string
		c      *cmd
		expect error
	}{
		{
			name:   "commands: pwd",
			c:      newCmd("pwd"),
			expect: nil,
		},
		{
			name:   "commands: ls -l .",
			c:      newCmd("ls", "-l", "."),
			expect: nil,
		},
		{
			name:   "commands: ''",
			c:      newCmd(""),
			expect: errors.New("commands nil"),
		},
	}
	for _, test := range tests {
		t.Run((test.name), func(t *testing.T) {
			if _, err := test.c.RunAndCapture(); (nil == err) != (nil == test.expect) {
				t.Errorf("expect = %v, got = %v", test.expect, err)
			}
		})
	}
}
