package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/JayceChant/commit-msg/validator"
)

var (
	versionFlag = flag.Bool("version", false, "")
	version     string
	goVersion   string
	commitHash  string
	buildTime   string
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
	fmt.Println(filepath.Base(cmd), version)
	if goVersion != "" {
		fmt.Println(goVersion)
	}
	if commitHash != "" {
		fmt.Println("commit hash :", commitHash)
	}
	if buildTime != "" {
		fmt.Println("build at :", buildTime)
	}
}
