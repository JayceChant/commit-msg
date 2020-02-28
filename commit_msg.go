package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

)

func getMsg(path string) string {
	if path == "" {
		ArgumentMissing.LogAndExit()
	}

	f, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		FileMissing.LogAndExit(path)
	}

	if f.IsDir() {
		log.Println(path, "is not a file.")
		FileMissing.LogAndExit(path)
	}

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
		ReadError.LogAndExit(path)
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
	WrongType.LogAndExit(typ, Types)
}

func checkHeader(header string) {
	if checkEmpty(header) {
		EmptyHeader.LogAndExit()
	}

	re := regexp.MustCompile(headerPattern)
	groups := re.FindStringSubmatch(header)

	if groups == nil || checkEmpty(groups[5]) {
		BadHeaderFormat.LogAndExit(header)
	}

	typ := groups[3]
	checkType(typ)

	isFixupOrSquash := (groups[2] != "")
	// TODO: 根据配置对scope检查
	// scope := groups[4]
	// TODO: 根据规则对subject检查
	// subject := groups[5]

	length := len(header)
	if length > Config.LineLimit &&
		!(isFixupOrSquash || typ == "revert" || typ == "Revert") {
		LineOverLong.LogAndExit(length, Config.LineLimit, header)
	}
}

func checkBody(body string) {
	if checkEmpty(body) {
		if Config.BodyRequired {
			BodyMissing.LogAndExit()
		} else {
			Validated.LogAndExit()
		}
	}

	if !checkEmpty(strings.SplitN(body, "\n", 2)[0]) {
		NoBlankLineBeforeBody.LogAndExit()
	}

	for _, line := range strings.Split(body, "\n") {
		length := len(line)
		if length > Config.LineLimit {
			LineOverLong.LogAndExit(length, Config.LineLimit, line)
		}
	}
}

func validateMsg(msg string) {
	if checkEmpty(msg) {
		EmptyMessage.LogAndExit()
	}

	if strings.HasPrefix(msg, mergePrefix) {
		Merge.LogAndExit()
	}

	sections := strings.SplitN(msg, "\n", 2)

	if m, _ := regexp.MatchString(revertPattern, sections[0]); !m {
		checkHeader(sections[0])
	}

	if len(sections) == 2 {
		checkBody(sections[1])
	} else if Config.BodyRequired {
		BodyMissing.LogAndExit()
	}

	Validated.LogAndExit()
}
