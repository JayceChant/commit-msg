package main

import (
	"log"

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
	if state <= Merge {
		log.Printf(state.Hint(), v...)
	} else if state <= FileMissing {
		log.Printf(state.Hint(), v...)
	} else {
		log.Printf(state.Hint(), v...)
		log.Printf(Lang.Rule, Types)
	}
	panic(state)
}
