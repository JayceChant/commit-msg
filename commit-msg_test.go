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

func TestValidateSample(t *testing.T) {
	defer func() {
		err := recover()
		state, ok := err.(msgState)
		if !ok || state > Merge {
			t.Errorf("Failed! %v", err)
		}
	}()

	validateMsg(getMsg("testcase/sample.txt"))
}
