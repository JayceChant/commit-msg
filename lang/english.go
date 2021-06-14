package lang

import (
	. "github.com/JayceChant/commit-msg/state"
)

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
	if you can not find any error after check, maybe you use full-width colon, or lack of whitespace after the colon.`,
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
)
