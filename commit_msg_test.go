// commit-msg_test.go
package main

import (
	"testing"

)

func TestGetMsg(t *testing.T) {
	if msg := getMsg("testcase/sample.txt"); msg == "" {
		t.Error("Failed!")
	}
}

func handleState(t *testing.T, s interface{}, expected MessageState) {
	state, ok := s.(MessageState)
	if !ok || state != expected {
		t.Errorf("Failed! %v", s)
	}
}

func TestValidatedSample(t *testing.T) {
	defer func() {
		handleState(t, recover(), Validated)
	}()

	validateMsg(getMsg("testcase/sample.txt"))
}

func TestMerge(t *testing.T) {
	defer func() {
		handleState(t, recover(), Merge)
	}()

	validateMsg(getMsg("testcase/tortoiseGitMerge.txt"))
}

func TestRevert(t *testing.T) {
	defer func() {
		handleState(t, recover(), Validated)
	}()

	validateMsg(getMsg("testcase/tortoiseGitRevert.txt"))
}
