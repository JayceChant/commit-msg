package main

import (
	"log"
	"os"

)

// MessageState indicate the state of a commit message
type MessageState int

// message states
const (
	// normal state
	Validated MessageState = iota
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
func (state MessageState) Hint() string {
	return Lang.HintList[state]
}

// LogAndExit ...
func (state MessageState) LogAndExit(v ...interface{}) {
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
func (state MessageState) IsNormal() bool {
	return state <= Merge
}

// IsFormatError return if the state a format error
func (state MessageState) IsFormatError() bool {
	return state >= EmptyMessage
}
