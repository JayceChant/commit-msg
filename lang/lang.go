package lang

import (
	"fmt"

	. "github.com/JayceChant/commit-msg/state"
)

func LoadLanguage(language string) LangPack {
	if l, ok := langs[language]; ok {
		return l
	}
	return langEn
}

// langPack ...
type langPack struct {
	Hints map[State]string
	Rule  string
}

func (l *langPack) GetHint(state State, v ...interface{}) string {
	return fmt.Sprintf(l.Hints[state], v...)
}

func (l *langPack) GetRule(types string) string {
	return fmt.Sprintf(l.Rule, types)
}

var (
	langs = map[string]*langPack{
		"en":    langEn,
		"zh":    langZhCn,
		"zh-CN": langZhCn,
	}
)
