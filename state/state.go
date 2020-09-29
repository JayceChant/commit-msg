package state

import (
	"log"
	"os"
)

var (
	lang  *LangPack
	types string
	// type check
	_ error = State(0)
)

// Init ...
func Init(l string, t string) {
	var ok bool
	if lang, ok = langs[l]; !ok {
		lang = langEn
	}
	types = t
}

// State indicate the state of a commit message
type State int

// message states
const (
	// normal state
	Validated State = iota
	Merge
	// non format error
	ArgumentMissing
	FileMissing
	ReadError
	// format error
	EmptyMessage
	EmptyHeader
	BadHeaderFormat
	WrongType
	ScopeMissing
	WrongScope
	BodyMissing
	NoBlankLineBeforeBody
	LineOverLong
	UndefindedError
)

// Error ...
func (state State) Error() string {
	return lang.Hints[state]
}

// LogAndExit ...
func (state State) LogAndExit(v ...interface{}) {
	log.Printf(state.Error(), v...)

	if state.IsNormal() {
		os.Exit(0)
	}

	if state.IsFormatError() {
		log.Printf(lang.Rule, types)
	}

	os.Exit(int(state))
}

// IsNormal return if the state a normal state
func (state State) IsNormal() bool {
	return state <= Merge
}

// IsFormatError return if the state a format error
func (state State) IsFormatError() bool {
	return state >= EmptyMessage
}
