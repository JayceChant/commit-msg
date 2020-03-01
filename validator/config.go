package validator

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/JayceChant/commit-msg/state"
	homedir "github.com/mitchellh/go-homedir"

)

const (
	configFileName = ".commit-msg.json"
	hookDir        = "./.git/hooks/"
)

type globalConfig struct {
	Lang         string
	BodyRequired bool
	LineLimit    int
}

var (
	// Config ...
	Config *globalConfig = &globalConfig{Lang: "en", BodyRequired: false, LineLimit: 80}
	// TypeList ...
	TypeList = []string{
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
	// Types ...
	Types = strings.Join(TypeList, ", ")
)

func locateConfigs() []string {
	cfgs := make([]string, 0)
	if home, err := homedir.Dir(); err == nil {
		cfgs = append(cfgs, path.Join(home, configFileName))
	}

	f, err := os.Stat(hookDir)
	if err == nil && f.IsDir() {
		// working dir is on project root
		cfgs = append(cfgs, path.Join(hookDir, configFileName))
	} else {
		// work around for test
		cfgs = append(cfgs, configFileName)
	}
	return cfgs
}

func loadConfig(path string, cfg *globalConfig) *globalConfig {
	f, err := os.Open(path)
	if err != nil {
		return cfg
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	if err := dec.Decode(cfg); err != nil {
		fmt.Println(err)
	}
	return cfg
}

// func initConfig(path string) *globalConfig {
// 	cfg := &globalConfig{"en", false, 80}
// 	f, err := os.Create(path)
// 	if err != nil {
// 		return cfg
// 	}
// 	defer f.Close()
// 	enc := json.NewEncoder(f)
// 	enc.SetIndent("", "    ")
// 	enc.Encode(cfg)
// 	return cfg
// }

func init() {
	paths := locateConfigs()
	for _, p := range paths {
		// TODO: fix json value overlaping
		Config = loadConfig(p, Config)
	}
	// if Config == nil {
	// 	Config = initConfig(path)
	// }

	var l *state.LangPack
	switch Config.Lang {
	case "zh", "zh-CN":
		l = initLangZhCn()
	default:
		l = initLangEn()
	}

	state.Config(l, Types)
}

func initLangEn() *state.LangPack {
	hints := []string{
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
	return &state.LangPack{
		Hints: hints,
		Rule:  rule,
	}
}

func initLangZhCn() *state.LangPack {
	hints := []string{
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
	return &state.LangPack{
		Hints: hints,
		Rule:  rule,
	}
}
