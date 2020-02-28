package main

import (
	"os"

	"github.com/JayceChant/commit-msg/validator"

)

func main() {
	msgFile := ""
	if len(os.Args) > 1 {
		msgFile = os.Args[1]
	}

	validator.Validate(msgFile)
}
