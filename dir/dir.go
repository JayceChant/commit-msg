package dir

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

const (
	hookDir = "./.git/hooks/"
)

func FindFiles(file string) []string {
	paths := make([]string, 0)
	if home, err := homedir.Dir(); err == nil {
		paths = append(paths, filepath.Join(home, file))
	}

	f, err := os.Stat(hookDir)
	if (err == nil || os.IsExist(err)) && f.IsDir() {
		// working dir is on project root
		paths = append(paths, filepath.Join(hookDir, file))
	} else {
		// work around for test
		paths = append(paths, file)
	}
	return paths
}

func FindFirstExist(file string) string {
	if isFile(file) {
		return file
	}

	hookFile := filepath.Join(hookDir, file)
	if isFile(hookFile) {
		return hookFile
	}

	if home, err := homedir.Dir(); err == nil {
		homeFile := filepath.Join(home, file)
		if isFile(homeFile) {
			return homeFile
		}
	}

	return ""
}

func isFile(path string) bool {
	fi, err := os.Stat(path)
	if err != nil && !os.IsExist(err) {
		return false
	}
	return !fi.IsDir()
}
