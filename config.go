package main

const (
	MERGE_PATTERN  = `^Merge `
	HEADER_PATTERN = `^((fixup! |squash! )?(\w+)(?:\(([^\)\s]+)\))?: (.+))(?:\n|$)`
	LINE_LIMIT     = 80
	BODY_REQUIRED  = false
)

type MsgState int

const (
	// normal state
	Validated MsgState = iota
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

var (
	msgStates = [...]string{
		"Validated",
		"Merge",
		"ArgumentMissing",
		"FileMissing",
		"ReadError",
		"EmptyMessage",
		"EmptyHeader",
		"BadHeaderFormat",
		"WrongType",
		"BodyMissing",
		"NoBlankLineBeforeBody",
		"LineOverLong",
		"UndefindedError",
	}
	stateHint = [...]string{
		"%s: commit message meet the rule.\n",
		"%s: merge commit detected，skip check.\n",
		"Error %s: commit message file argument missing.\n",
		"Error %s: file %s not exists.\n",
		"Error %s: read file %s error.\n",
		"Error %s: commit message has no content except whitespaces.\n",
		"Error %s: header (first line) has no content except whitespaces.\n",
		`Error %s: header (first line) not following the rule:
%s
if you can not find any error after check, maybe you use Chinese colon, or lack of whitespace after the colon.`,
		"Error %s: %s, should be one of the keywords:\n%s\n",
		"Error %s: body has no content except whitespaces.\n",
		"Error %s: no empty line between header and body.\n",
		"Error %s: the length of line is %d, exceed %d:\n%s\n",
		"Error %s: unexpected error occurs, please raise an issue.",
	}
	typeList = [...]string{
		"feat",     // new feature 新功能
		"fix",      // fix bug 修复
		"docs",     // documentation 文档
		"style",    // changes not affect logic 格式（不影响代码运行的变动）
		"refactor", // 重构（既不是新增功能，也不是修改bug的代码变动）
		"perf",     // performance 提升性能
		"test",     // 增加测试
		"chore",    // 构建过程或辅助工具的变动'
		"revert",   // 撤销以前的 commit
		"Revert",   // 有些工具生成的 revert 首字母大写
	}
	ruleHint = `Commit message rule as follow:
<type>(<scope>): <subject>
// empty line
<body>
// empty line
<footer>

(<scope>), <body> and <footer> are optional
<type>  must be one of %s
more specific instructions, please refer to https://github.com/JayceChant/commit-msg.go`
)

func (state MsgState) Name() string {
	return msgStates[state]
}

func (state MsgState) Hint() string {
	return stateHint[state]
}
