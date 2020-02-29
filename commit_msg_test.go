package main

import (
	"os"
	"os/exec"
	"runtime"
	"testing"

)

func TestGetMsg(t *testing.T) {
	if msg := getMsg("testcase/sample.txt"); msg == "" {
		t.Error("Failed!")
	}
}

func testExitFunc(t *testing.T, f func(), expected int) {
	if os.Getenv("TEST_RUNNER") == "1" {
		f()
		return
	}
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	ft := runtime.FuncForPC(pc[0])
	cmd := exec.Command(os.Args[0], "-test.run="+ft.Name())
	cmd.Env = append(os.Environ(), "TEST_RUNNER=1")
	err := cmd.Run()
	e, ok := err.(*exec.ExitError)

	if expected == 0 {
		if ok && e.ExitCode() != 0 {
			t.Errorf("exit code got %d, expected %d", e.ExitCode(), 0)
		}
		return
	}

	if !ok {
		t.Error(err)
		return
	}

	if e.ExitCode() != expected {
		t.Errorf("exit code got %d, expected %d", e.ExitCode(), expected)
	}
}

func TestValidatedSample(t *testing.T) {
	testExitFunc(t, func() {
		validateMsg(getMsg("testcase/sample.txt"))
	}, 0)
}

func TestMerge(t *testing.T) {
	testExitFunc(t, func() {
		validateMsg(getMsg("testcase/tortoiseGitMerge.txt"))
	}, 0)
}

func TestRevert(t *testing.T) {
	testExitFunc(t, func() {
		validateMsg(getMsg("testcase/tortoiseGitRevert.txt"))
	}, 0)
}
