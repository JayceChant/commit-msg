package main

import (
	"os"

	"github.com/JayceChant/commit-msg/message"

)

func main() {
	msgFile := ""
	if len(os.Args) > 1 {
		msgFile = os.Args[1]
	}

	msg := getMsg(msgFile)

	defer func() {
		err := recover()
		state, ok := err.(message.State)
		if !ok {
			panic(err)
		}

		if state.IsNormal() {
			os.Exit(0)
		} else {
			os.Exit(int(state))
		}
	}()
	validateMsg(msg)
}
