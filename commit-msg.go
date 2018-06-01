package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func logAndExit(state MsgState, v ...interface{}) {
	a := append([]interface{}{state.Name()}, v...)
	if state <= Merge {
		log.Printf(state.Hint(), a...)
		os.Exit(0)
	} else if state <= FileMissing {
		log.Fatalf(state.Hint(), a...)
	} else {
		log.Printf(state.Hint(), a...)
		log.Fatalf(ruleHint, strings.Join(typeList[:], ", "))
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

func checkType(type_ string) {
	for _, t := range typeList {
		if type_ == t {
			return
		}
	}
	logAndExit(WrongType, type_, strings.Join(typeList[:], ", "))
}

func checkHeader(header string) {
	if checkEmpty(header) {
		logAndExit(EmptyHeader)
	}

	re := regexp.MustCompile(HEADER_PATTERN)
	groups := re.FindStringSubmatch(header)

	if groups == nil || checkEmpty(groups[5]) {
		logAndExit(BadHeaderFormat, header)
	}

	type_ := groups[3]
	checkType(type_)

	isFixupOrSquash := (groups[2] != "")
	// scope := groups[4] // TODO: 根据配置对scope检查
	// subject := groups[5] // TODO: 根据规则对subject检查

	length := len(header)
	if length > LINE_LIMIT &&
		!(isFixupOrSquash || type_ == "revert" || type_ == "Revert") {
		logAndExit(LineOverLong, length, LINE_LIMIT, header)
	}
}

func checkBody(body string) {
	if checkEmpty(body) {
		if BODY_REQUIRED {
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
		if length > LINE_LIMIT {
			logAndExit(LineOverLong, length, LINE_LIMIT, line)
		}
	}
}

func validateMsg(msg string) {
	if checkEmpty(msg) {
		logAndExit(EmptyMessage)
	}

	isMerge, err := regexp.MatchString(MERGE_PATTERN, msg)
	if err != nil {
		log.Println(err)
		logAndExit(UndefindedError)
	}

	if isMerge {
		logAndExit(Merge)
	}

	sections := strings.SplitN(msg, "\n", 2)
	checkHeader(sections[0])

	if len(sections) == 2 {
		checkBody(sections[1])
	} else if BODY_REQUIRED {
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
