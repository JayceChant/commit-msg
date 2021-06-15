package lang

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/JayceChant/commit-msg/dir"
	. "github.com/JayceChant/commit-msg/state"
)

const (
	langFile = "commit-msg.%s.json"
)

func LoadLanguage(language string) LangPack {
	if l, ok := langs[language]; ok {
		return l
	}

	file := fmt.Sprintf(langFile, language)
	path := dir.FindFirstExist(file)
	if path == "" {
		return langEn
	}

	f, err := os.Open(path)
	if err != nil && !os.IsExist(err) {
		log.Printf("open language %v error: %v\n", language, err)
		return langEn
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	l := &langPack{}
	err = dec.Decode(l)
	if err != nil {
		log.Printf("decode language %v error: %v\n", language, err)
		return langEn
	}
	return l
}

// langPack ...
type langPack struct {
	Hints map[State]string `json:"hints"`
	Rule  string           `json:"rule"`
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
