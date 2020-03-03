package validator

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"

	"github.com/JayceChant/commit-msg/state"
)

var (
	config = &globalConfig{Lang: "en", BodyRequired: true, LineLimit: 80}
)

func assertExitCode(t *testing.T, f func(), flag string, expected int) {
	if env := os.Getenv("TEST_RUNNER"); env != "" {
		if env == flag {
			f()
		}
		return
	}

	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	ft := runtime.FuncForPC(pc[0])
	nn := strings.Split(ft.Name(), ".")

	cmd := exec.Command(os.Args[0], "-test.run=^("+nn[len(nn)-1]+")$")
	cmd.Env = append(os.Environ(), "TEST_RUNNER="+flag)
	err := cmd.Run()
	e, ok := err.(*exec.ExitError)

	if expected == 0 {
		if ok && e.ExitCode() != 0 {
			t.Errorf("exit code got %d, expected %d", e.ExitCode(), 0)
		}
		return
	}

	if !ok {
		t.Errorf("expect ExitError with code %d, got err %v", expected, err)
		return
	}

	if e.ExitCode() != expected {
		t.Errorf("exit code got %d, expected %d", e.ExitCode(), expected)
	}
}

func TestGetMsg(t *testing.T) {
	assertExitCode(t, func() {
		getMsg("")
	}, "ArgumentMissing", int(state.ArgumentMissing))

	assertExitCode(t, func() {
		getMsg("file_not_existed.txt")
	}, "FileMissing", int(state.FileMissing))

	assertExitCode(t, func() {
		if msg := getMsg("testcase/normal_sample.txt"); msg == "" {
			os.Exit(int(state.EmptyMessage))
		}
	}, "Normal", 0)
}

var (
	emptyCases = []struct {
		text string
		want bool
	}{
		{"", true},
		{"   ", true},
		{"  	 ", true},
		{"\n\r\t", true},
		{"  some words ", false},
	}
)

func TestCheckEmpty(t *testing.T) {
	for _, tt := range emptyCases {
		t.Run("checkEmpty", func(t *testing.T) {
			if got := checkEmpty(tt.text); got != tt.want {
				t.Errorf(`checkEmpty("%s")=%v, want %v`, tt.text, got, tt.want)
			}
		})
	}
}

func TestCheckType(t *testing.T) {
	assertExitCode(t, func() {
		checkType("feat")
	}, "feat", 0)

	assertExitCode(t, func() {
		checkType("")
	}, "no_type", int(state.WrongType))

	assertExitCode(t, func() {
		checkType("Feat")
	}, "wrong_type", int(state.WrongType))
}

var (
	headerCases = []struct {
		text string
		name string
		want int
	}{
		{"", "empty_header", int(state.EmptyHeader)},
		{"\r\r\t\n", "empty_header2", int(state.EmptyHeader)},
		{"something in wrong format", "bad_header_format", int(state.BadHeaderFormat)},
		{"test:header without space after colon", "bad_header_no_colon", int(state.BadHeaderFormat)},
		{"test：Chinese(full width) colon", "bad_header_full_width", int(state.BadHeaderFormat)},
		{"test: ", "bad_header_no_title", int(state.BadHeaderFormat)},
	}
)

func TestCheckHeader(t *testing.T) {
	for _, tt := range headerCases {
		assertExitCode(t, func() {
			checkHeader(tt.text, config)
		}, tt.name, tt.want)
	}
}

var (
	bodyCases = []struct {
		text string
		name string
		want int
	}{
		{"", "body_missing", int(state.BodyMissing)},
		{"body", "no_empty_line", int(state.NoBlankLineBeforeBody)},
		{"\r\na body with too looooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooong line", "line_over_long", int(state.LineOverLong)},
		{"\r\nnormal body", "normal", 0},
	}
)

func TestCheckBody(t *testing.T) {
	for _, tt := range bodyCases {
		assertExitCode(t, func() {
			checkBody(tt.text, config)
		}, tt.name, tt.want)
	}
}

func TestValidate(t *testing.T) {
	assertExitCode(t, func() {
		Validate("testcase/normal_sample.txt")
	}, "1", 0)
}

func TestTortoiseGit(t *testing.T) {
	assertExitCode(t, func() {
		validateMsg("Merge remote-tracking branch 'remotes/origin/feat_xyz'", config)
	}, "Merge", 0)

	assertExitCode(t, func() {
		validateMsg(`Revert "fix: abc issue & xyz problems, some words to make it lonnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnger"

This reverts commit 1234567890abcdef1234567890abcdef12345678.`, config)
	}, "Revert", 0)
}
