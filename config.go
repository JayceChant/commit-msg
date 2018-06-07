package main

import (
	"encoding/json"
	"os"
	"strings"
)

const (
	MERGE_PREFIX    = `Merge `
	HEADER_PATTERN   = `^((fixup! |squash! )?(\w+)(?:\(([^\)\s]+)\))?: (.+))(?:\n|$)`
	CONFIG_FILE_NAME = "commit-msg.cfg.json"
	HOOK_DIR         = "./.git/hooks/"
)

type GlobalConfig struct {
	Lang         string
	BodyRequired bool
	LineLimit    int
}

type LangPack struct {
	HintList []string
	Rule     string
}

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
	Config   *GlobalConfig
	Lang     *LangPack
	TypeList = [...]string{
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
	Types = strings.Join(TypeList[:], ", ")
)

func (state MsgState) Hint() string {
	return Lang.HintList[state]
}

func locateConfig() string {
	f, err := os.Stat(HOOK_DIR)
	if err != nil || !f.IsDir() {
		return CONFIG_FILE_NAME
	}
	return HOOK_DIR + CONFIG_FILE_NAME
}

func loadConfig(path string) *GlobalConfig {
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	var cfg GlobalConfig
	if err := dec.Decode(&cfg); err != nil {
		return nil
	}
	return &cfg
}

func initConfig(path string) *GlobalConfig {
	cfg := &GlobalConfig{"en", false, 80}
	f, err := os.Create(path)
	if err != nil {
		return cfg
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "    ")
	enc.Encode(cfg)
	return cfg
}

func init() {
	path := locateConfig()
	Config = loadConfig(path)
	if Config == nil {
		Config = initConfig(path)
	}

	switch Config.Lang {
	case "zh", "zh-CN":
		Lang = initLangZhCn()
	default:
		Lang = initLangEn()
	}
}

func initLangEn() *LangPack {
	hint := []string{
		"Validated: commit message meet the rule.\n",
		"Merge: merge commit detected，skip check.\n",
		"Error ArgumentMissing: commit message file argument missing.\n",
		"Error FileMissing: file %s not exists.\n",
		"Error ReadError: read file %s error.\n",
		"Error EmptyMessage: commit message has no content except whitespaces.\n",
		"Error EmptyHeader: header (first line) has no content except whitespaces.\n",
		`Error BadHeaderFormat: header (first line) not following the rule:
%s
if you can not find any error after check, maybe you use Chinese colon, or lack of whitespace after the colon.`,
		"Error WrongType: %s, should be one of the keywords:\n%s\n",
		"Error BodyMissing: body has no content except whitespaces.\n",
		"Error NoBlankLineBeforeBody: no empty line between header and body.\n",
		"Error LineOverLong: the length of line is %d, exceed %d:\n%s\n",
		"Error UndefindedError: unexpected error occurs, please raise an issue.\n",
	}
	rule := `Commit message rule as follow:
<type>(<scope>): <subject>
// empty line
<body>
// empty line
<footer>

(<scope>), <body> and <footer> are optional
<type>  must be one of %s
more specific instructions, please refer to: https://github.com/JayceChant/commit-msg.go`
	return &LangPack{hint, rule}
}

func initLangZhCn() *LangPack {
	hint := []string{
		"Validated: 提交信息符合规范。\n",
		"Merge: 合并提交，跳过规范检查。\n",
		"Error ArgumentMissing: 缺少文件参数。\n",
		"Error FileMissing: 文件 %s 不存在。\n",
		"Error ReadError: 读取 %s 错误。\n",
		"Error EmptyMessage: 提交信息没有内容（不包括空白字符）。\n",
		"Error EmptyHeader: 标题（第一行）没有内容（不包括空白字符）。\n",
		`Error BadHeaderFormat: 标题（第一行）不符合规范:
%s
如果您无法发现错误，请注意是否使用了中文冒号，或者冒号后面缺少空格。`,
		"Error WrongType: %s, 类型关键字应为以下选项中的一个:\n%s\n",
		"Error BodyMissing: 消息体没有内容（不包括空白字符）。\n",
		"Error NoBlankLineBeforeBody: 标题和消息体之间缺少空行。\n",
		"Error LineOverLong: 该行长度为 %d, 超出了 %d 的限制:\n%s\n",
		"Error UndefindedError: 没有预料到的错误，请提交一个错误报告。\n",
	}
	rule := `提交信息规范如下:
<type>(<scope>): <subject>
// 空行
<body>
// 空行
<footer>

(<scope>), <body> 和 <footer> 可选
<type> 必须是关键字 %s 之一
更多信息，请参考项目主页: https://github.com/JayceChant/commit-msg.go`
	return &LangPack{hint, rule}
}
