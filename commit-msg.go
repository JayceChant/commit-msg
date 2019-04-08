package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func logAndExit(state msgState, v ...interface{}) {
	if state <= Merge {
		log.Printf(state.Hint(), v...)
		os.Exit(0)
	} else if state <= FileMissing {
		log.Fatalf(state.Hint(), v...)
	} else {
		log.Printf(state.Hint(), v...)
		log.Fatalf(Lang.Rule, Types)
	}
}

func getMsg(path string) string {
	if path == "" {
		logAndExit(ArgumentMissing)
	}

	f, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		logAndExit(FileMissing, path)
	}

	if f.IsDir() {
		log.Println(path, "is not a file.")
		logAndExit(FileMissing, path)
	}

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
		logAndExit(ReadError, path)
	}

	return string(buf)
}

func checkEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

func checkType(typ string) {
	for _, t := range TypeList {
		if typ == t {
			return
		}
	}
	logAndExit(WrongType, typ, Types)
}

func checkHeader(header string) {
	if checkEmpty(header) {
		logAndExit(EmptyHeader)
	}

	re := regexp.MustCompile(headerPattern)
	groups := re.FindStringSubmatch(header)

	if groups == nil || checkEmpty(groups[5]) {
		logAndExit(BadHeaderFormat, header)
	}

	typ := groups[3]
	checkType(typ)

	isFixupOrSquash := (groups[2] != "")
	// scope := groups[4] // TODO: 根据配置对scope检查
	// subject := groups[5] // TODO: 根据规则对subject检查

	length := len(header)
	if length > Config.LineLimit &&
		!(isFixupOrSquash || typ == "revert" || typ == "Revert") {
		logAndExit(LineOverLong, length, Config.LineLimit, header)
	}
}

func checkBody(body string) {
	if checkEmpty(body) {
		if Config.BodyRequired {
			logAndExit(BodyMissing)
		} else {
			logAndExit(Validated)
		}
	}

	if !checkEmpty(strings.SplitN(body, "\n", 2)[0]) {
		logAndExit(NoBlankLineBeforeBody)
	}

	for _, line := range strings.Split(body, "\n") {
		length := len(line)
		if length > Config.LineLimit {
			logAndExit(LineOverLong, length, Config.LineLimit, line)
		}
	}
}

func validateMsg(msg string) {
	if checkEmpty(msg) {
		logAndExit(EmptyMessage)
	}

	if strings.HasPrefix(msg, mergePrefix) {
		logAndExit(Merge)
	}

	sections := strings.SplitN(msg, "\n", 2)
	checkHeader(sections[0])

	if len(sections) == 2 {
		checkBody(sections[1])
	} else if Config.BodyRequired {
		logAndExit(BodyMissing)
	}

	logAndExit(Validated)
}

func main() {
	msgFile := ""
	if len(os.Args) > 1 {
		msgFile = os.Args[1]
	}

	msg := getMsg(msgFile)

	validateMsg(msg)
}
