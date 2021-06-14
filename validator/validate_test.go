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
	zeroCfg         = &validateConfig{}
	defaultCfg      = &validateConfig{Lang: "en", BodyRequired: true, LineLimit: 80}
	scopeRequired   = &validateConfig{ScopeRequired: true}
	scopesSpecified = &validateConfig{Scopes: []string{"model", "view", "controller"}}
)

func assertExitCode(t *testing.T, f func(), name string, expected int) {
	if env := os.Getenv("TEST_RUNNER"); env != "" {
		if env == name {
			f()
		}
		return
	}

	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	ft := runtime.FuncForPC(pc[0])
	nn := strings.Split(ft.Name(), ".")

	cmd := exec.Command(os.Args[0], "-test.run=^("+nn[len(nn)-1]+")$")
	cmd.Env = append(os.Environ(), "TEST_RUNNER="+name)
	err := cmd.Run()
	e, ok := err.(*exec.ExitError)

	if expected == 0 {
		if ok && e.ExitCode() != 0 {
			t.Errorf("%s: exit code got %d, expected 0", name, e.ExitCode())
		}
		return
	}

	if !ok {
		t.Errorf("%s: expect ExitError with code %d, got err %v", name, expected, err)
		return
	}

	if e.ExitCode() != expected {
		t.Errorf("%s: exit code got %d, expected %d", name, e.ExitCode(), expected)
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

func TestValidateEmpty(t *testing.T) {
	var emptyCases = []struct {
		text string
		want bool
	}{
		{"", true},
		{"   ", true},
		{"  	 ", true},
		{"\n\r\t", true},
		{"  some words ", false},
	}
	for _, tt := range emptyCases {
		t.Run("checkEmpty", func(t *testing.T) {
			if got := isEmpty(tt.text); got != tt.want {
				t.Errorf(`checkEmpty("%s")=%v, want %v`, tt.text, got, tt.want)
			}
		})
	}
}

func TestValidateType(t *testing.T) {
	assertExitCode(t, func() {
		validateType("feat")
	}, "feat", 0)

	assertExitCode(t, func() {
		validateType("")
	}, "no_type", int(state.WrongType))

	assertExitCode(t, func() {
		validateType("Feat")
	}, "wrong_type", int(state.WrongType))
}

func TestValidateHeader(t *testing.T) {
	var headerCases = []struct {
		text   string
		name   string
		config *validateConfig
		want   int
	}{
		{"", "empty_header", defaultCfg, int(state.EmptyHeader)},
		{"\r\r\t\n", "empty_header2", defaultCfg, int(state.EmptyHeader)},
		{"something in wrong format", "bad_header_format", defaultCfg, int(state.BadHeaderFormat)},
		{"test:header without space after colon", "bad_header_no_colon", defaultCfg, int(state.BadHeaderFormat)},
		{"test：Chinese(full width) colon", "bad_header_full_width", defaultCfg, int(state.BadHeaderFormat)},
		{"test: ", "bad_header_no_title", defaultCfg, int(state.BadHeaderFormat)},
		{"feat: something changes", "scope_missing", scopeRequired, int(state.ScopeMissing)},
		{"feat( ): something changes", "empty_scope", scopeRequired, int(state.BadHeaderFormat)},
		{"feat: header that too lonnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnng", "header_too_long", defaultCfg, int(state.LineOverLong)},
		{"feat: header that too lonnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnng", "header_length_no_limit", zeroCfg, 0},
	}
	for _, tt := range headerCases {
		assertExitCode(t, func() {
			validateHeader(tt.text, tt.config)
		}, tt.name, tt.want)
	}
}

func TestValidateBody(t *testing.T) {
	var bodyCases = []struct {
		text   string
		name   string
		config *validateConfig
		want   int
	}{
		{"", "body_missing", defaultCfg, int(state.BodyMissing)},
		{"body", "no_blank_line", defaultCfg, int(state.NoBlankLineBeforeBody)},
		{"\r\na body with too looooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooong line", "body_line_over_long", defaultCfg, int(state.LineOverLong)},
		{"\r\na body with too looooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooong line", "body_line_no_limit", zeroCfg, 0},
		{"\r\nnormal body", "normal", defaultCfg, 0},
	}
	for _, tt := range bodyCases {
		assertExitCode(t, func() {
			validateBody(tt.text, tt.config)
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
		validateMsg("Merge remote-tracking branch 'remotes/origin/feat_xyz'", defaultCfg)
	}, "Merge", 0)

	assertExitCode(t, func() {
		validateMsg(`Revert "fix: abc issue & xyz problems, some words to make it lonnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnger"

This reverts commit 1234567890abcdef1234567890abcdef12345678.`, defaultCfg)
	}, "Revert", 0)
}

func TestValidateScope(t *testing.T) {
	var scopeCases = []struct {
		text   string
		name   string
		config *validateConfig
		want   int
	}{
		{"model", "normal", scopeRequired, 0},
		{"", "empty_but_not_required", defaultCfg, 0},
		{"", "empty_scope", scopeRequired, int(state.ScopeMissing)},
		{"model", "scope_in_range", scopesSpecified, 0},
		{"module", "wrong_scope", scopesSpecified, int(state.WrongScope)},
	}
	for _, tt := range scopeCases {
		assertExitCode(t, func() {
			validateScope(tt.text, tt.config)
		}, tt.name, tt.want)
	}
}
