package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	MERGE_PATTERN  = `^Merge `
	HEADER_PATTERN = `^((fixup! |squash! )?(\w+)(?:\(([^\)\s]+)\))?: (.+))(?:\n|$)`
	LINE_LIMIT     = 100
	BODY_REQUIRED  = false
)

var (
	typeList = [...]string{
		"feat",     // 新功能（feature）
		"fix",      // 修补bug
		"docs",     // 文档（documentation）
		"style",    // 格式（不影响代码运行的变动）
		"refactor", // 重构（既不是新增功能，也不是修改bug的代码变动）
		"perf",     // 提升性能（performance）
		"test",     // 增加测试
		"chore",    // 构建过程或辅助工具的变动'
		"revert",   // 撤销以前的 commit
		"Revert"}   // 有些版本的工具生成的revert message首字母大写
)

func getMsg(path string) string {
	if path == "" {
		log.Fatalln("Arg missing.")
	}

	f, err := os.Stat(path)
	if err != nil {
		log.Fatalln(err)
	}

	if f.IsDir() {
		log.Fatalln(path, "is not a file.")
	}

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
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
	log.Fatalln("Wrong type.")
}

func checkHeader(header string) {
	if checkEmpty(header) {
		log.Fatalln("Empty header.")
	}

	re := regexp.MustCompile(HEADER_PATTERN)
	groups := re.FindStringSubmatch(header)

	isFixupOrSquash := (groups[2] != "")

	type_ := groups[3]
	// scope = groups[4] // TODO: 根据配置对scope检查
	// subject = (groups[5] != "") // TODO: 根据规则对subject检查
	checkType(type_)

	length := len(header)
	if length > LINE_LIMIT && !(isFixupOrSquash || type_ == "revert") {
		log.Fatalln("Line overlong.")
	}
}

func checkBody(body string) {
	if checkEmpty(body) {
		if BODY_REQUIRED {
			log.Fatalln("Empty body.")
		} else {
			return
		}
	}

	if !checkEmpty(strings.SplitN(body, "\n", 2)[0]) {
		log.Fatalln("Blank line lacking between header and body.")
	}

	for _, line := range strings.Split(body, "\n") {
		if len(line) > LINE_LIMIT {
			log.Fatalln("Line overlong.")
		}
	}
}

func validateMsg(msg string) {
	msg = strings.TrimSpace(msg)
	if msg == "" {
		log.Fatalln("Empty Message.")
	}

	isMerge, err := regexp.MatchString(MERGE_PATTERN, msg)
	if err != nil {
		log.Fatalln(err)
	}

	if isMerge {
		log.Println("Merge message, skip check.")
		os.Exit(0)
	}

	sections := strings.SplitN(msg, "\n", 2)
	checkHeader(sections[0])

	if len(sections) == 2 {
		checkBody(sections[1])
	} else if BODY_REQUIRED {
		log.Fatalln("Body missing. Maybe a new line is lacking.")
	}
}

func main() {
	msgFile := ""
	if len(os.Args) > 1 {
		msgFile = os.Args[1]
	}

	msg := getMsg(msgFile)

	validateMsg(msg)
}
