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
