//go:generate stringer -type=State
package state

import (
	"encoding"
	"fmt"
	"log"
	"os"
)

var (
	lang  LangPack
	types string
)

// Init ...
func Init(language LangPack, typeStr string) {
	lang = language
	types = typeStr
}

type LangPack interface {
	GetHint(state State, v ...interface{}) string
	GetRule(types string) string
}

func _() {
	// type check
	var _ encoding.TextMarshaler = State(0)
	var _ encoding.TextUnmarshaler = (*State)(nil)
}

// State indicate the state of a commit message
type State int8

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

// LogAndExit ...
func (state State) LogAndExit(v ...interface{}) {
	log.Println(lang.GetHint(state, v...))

	if state.IsNormal() {
		os.Exit(0)
	}

	if state.IsFormatError() {
		log.Println(lang.GetRule(types))
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

func (state State) MarshalText() (text []byte, err error) {
	return []byte(state.String()), nil
}

func (state *State) UnmarshalText(text []byte) error {
	str := string(text)
	for s := Validated; s <= UndefindedError; s++ {
		if s.String() == str {
			*state = s
			return nil
		}
	}
	return fmt.Errorf("unknown state : %v", str)
}
