package state

// LangPack ...
type LangPack struct {
	Hints [13]string
	Rule  string
}

var (
	langEn = &LangPack{
		Hints: [...]string{
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
		},
		Rule: `Commit message rule as follow:
		<type>(<scope>): <subject>
		// empty line
		<body>
		// empty line
		<footer>
		
		(<scope>), <body> and <footer> are optional
		<type>  must be one of %s
		more specific instructions, please refer to: https://github.com/JayceChant/commit-msg`,
	}

	langZhCn = &LangPack{
		Hints: [...]string{
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
		},
		Rule: `提交信息规范如下:
		<type>(<scope>): <subject>
		// 空行
		<body>
		// 空行
		<footer>
		
		(<scope>), <body> 和 <footer> 可选
		<type> 必须是关键字 %s 之一
		更多信息，请参考项目主页: https://github.com/JayceChant/commit-msg`,
	}

	langs = map[string]*LangPack{
		"en":    langEn,
		"zh":    langZhCn,
		"zh-CN": langZhCn,
	}
)
