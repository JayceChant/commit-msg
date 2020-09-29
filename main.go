package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/JayceChant/commit-msg/validator"
)

const (
	versionStr = "v0.2.0"
)

var (
	versionFlag = flag.Bool("version", false, "")
)

func main() {
	flag.Parse()
	if *versionFlag {
		printVersion(os.Args[0])
		return
	}

	msgFile := ""
	if len(os.Args) > 1 {
		msgFile = os.Args[1]
	}

	validator.Validate(msgFile)
}

func printVersion(cmd string) {
	fmt.Printf("%s %s\n", filepath.Base(cmd), versionStr)
}
