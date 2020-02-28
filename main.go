package main

import (
	"os"

)

func main() {
	msgFile := ""
	if len(os.Args) > 1 {
		msgFile = os.Args[1]
	}

	msg := getMsg(msgFile)

	defer func() {
		err := recover()
		state, ok := err.(MessageState)
		if !ok {
			panic(err)
		}

		if state <= Merge {
			os.Exit(0)
		} else {
			os.Exit(int(state))
		}
	}()
	validateMsg(msg)
}
