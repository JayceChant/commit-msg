package state

import "fmt"

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
	langEn = &langPack{
		Hints: map[State]string{
			Validated:       "Validated: commit message meet the rule.",
			Merge:           "Merge: merge commit detected，skip check.",
			ArgumentMissing: "Error ArgumentMissing: commit message file argument missing.",
			FileMissing:     "Error FileMissing: file %s not exists.",
			ReadError:       "Error ReadError: read file %s error.",
			EmptyMessage:    "Error EmptyMessage: commit message has no content except whitespaces.",
			EmptyHeader:     "Error EmptyHeader: header (first line) has no content except whitespaces.",
			BadHeaderFormat: `Error BadHeaderFormat: header (first line) not following the rule:
	%s
	if you can not find any error after check, maybe you use Chinese colon, or lack of whitespace after the colon.`,
			WrongType:             "Error WrongType: %s, type should be one of the keywords:\n%s",
			ScopeMissing:          "Error ScopeMissing: (scope) is required right after type.",
			WrongScope:            "Error WrongScope: %s, scope should be one of the keywords:\n%s",
			BodyMissing:           "Error BodyMissing: body has no content except whitespaces.",
			NoBlankLineBeforeBody: "Error NoBlankLineBeforeBody: no empty line between header and body.",
			LineOverLong:          "Error LineOverLong: the length of line is %d, exceed %d:\n%s",
			UndefindedError:       "Error UndefindedError: unexpected error occurs, please raise an issue.",
		},
		Rule: `Commit message rule as follow:
		<type>(<scope>): <subject>
		// empty line
		<body>
		// empty line
		<footer>
		
		(<scope>), <body> and <footer> are optional by default
		<type>  must be one of %s
		more specific instructions, please refer to: https://github.com/JayceChant/commit-msg`,
	}

	langZhCn = &langPack{
		Hints: map[State]string{
			Validated:       "Validated: 提交信息符合规范。",
			Merge:           "Merge: 合并提交，跳过规范检查。",
			ArgumentMissing: "Error ArgumentMissing: 缺少文件参数。",
			FileMissing:     "Error FileMissing: 文件 %s 不存在。",
			ReadError:       "Error ReadError: 读取 %s 错误。",
			EmptyMessage:    "Error EmptyMessage: 提交信息没有内容（不包括空白字符）。",
			EmptyHeader:     "Error EmptyHeader: 标题（第一行）没有内容（不包括空白字符）。",
			BadHeaderFormat: `Error BadHeaderFormat: 标题（第一行）不符合规范:
	%s
	如果您无法发现错误，请注意是否使用了中文冒号，或者冒号后面缺少空格。`,
			WrongType:             "Error WrongType: %s, 类型关键字应为以下选项中的一个:\n%s",
			ScopeMissing:          "Error ScopeMissing: 类型后面缺少'(scope)'。",
			WrongScope:            "Error WrongScope: %s, 范围关键字应为以下选项中的一个:\n%s",
			BodyMissing:           "Error BodyMissing: 消息体没有内容（不包括空白字符）。",
			NoBlankLineBeforeBody: "Error NoBlankLineBeforeBody: 标题和消息体之间缺少空行。",
			LineOverLong:          "Error LineOverLong: 该行长度为 %d, 超出了 %d 的限制:\n%s",
			UndefindedError:       "Error UndefindedError: 没有预料到的错误，请提交一个错误报告。",
		},
		Rule: `提交信息规范如下:
		<type>(<scope>): <subject>
		// 空行
		<body>
		// 空行
		<footer>
		
		(<scope>), <body> 和 <footer> 默认可选，也可以在配置设置必选
		<type> 必须是关键字 %s 之一
		更多信息，请参考项目主页: https://github.com/JayceChant/commit-msg`,
	}

	langs = map[string]*langPack{
		"en":    langEn,
		"zh":    langZhCn,
		"zh-CN": langZhCn,
	}
)
