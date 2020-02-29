package state

import (
	"log"
	"os"

)

var (
	// Lang ...
	Lang *LangPack
	// Types ...
	Types string
)

// Config ...
func Config(l *LangPack, t string) {
	Lang = l
	Types = t
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
	BodyMissing
	NoBlankLineBeforeBody
	LineOverLong
	UndefindedError
)

// Hint ...
func (state State) Hint() string {
	return Lang.Hints[state]
}

// LogAndExit ...
func (state State) LogAndExit(v ...interface{}) {
	if state.IsNormal() {
		log.Printf(state.Hint(), v...)
		os.Exit(0)
	}

	if state.IsFormatError() {
		log.Printf(state.Hint(), v...)
		log.Printf(Lang.Rule, Types)
		os.Exit(int(state))
	}

	// non-format error
	log.Printf(state.Hint(), v...)
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
