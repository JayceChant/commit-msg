package validator

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/JayceChant/commit-msg/state"
	homedir "github.com/mitchellh/go-homedir"
)

const (
	configFileName = ".commit-msg.json"
	hookDir        = "./.git/hooks/"
)

type globalConfig struct {
	Lang          string   `json:"lang,omitempty"`
	BodyRequired  bool     `json:"bodyRequired,omitempty"`
	LineLimit     int      `json:"lineLimit,omitempty"`
	Types         []string `json:"types,omitempty"`
	DenyTypes     []string `json:"denyTypes,omitempty"`
	ScopeRequired bool     `json:"scopeRequired,omitempty"`
	Scopes        []string `json:"scopes,omitempty"`
}

type empty struct{}

var (
	// Config ...
	Config *globalConfig = &globalConfig{Lang: "en", BodyRequired: false, LineLimit: 80}
	// TypeSet ...
	TypeSet = map[string]empty{
		"feat":     {}, // new feature 新功能
		"fix":      {}, // fix bug 修复
		"docs":     {}, // documentation 文档
		"style":    {}, // changes not affect logic 格式（不影响代码运行的变动）
		"refactor": {}, // refactoring 重构（既不是新增功能，也不是修改bug的代码变动）
		"perf":     {}, // performance 提升性能
		"test":     {}, // 增加测试
		"chore":    {}, // 辅助工具的变动'
		"build":    {}, // build process 构建过程
		"ci":       {}, // continuous integration 持续集成相关
		"docker":   {}, // 容器相关
		"revert":   {}, // 撤销以前的 commit
		"Revert":   {}, // 有些工具生成的 revert 首字母大写
	}
	// TypesStr ...
	TypesStr string
)

func locateConfigs() []string {
	paths := make([]string, 0)
	if home, err := homedir.Dir(); err == nil {
		paths = append(paths, filepath.Join(home, configFileName))
	}

	f, err := os.Stat(hookDir)
	if (err == nil || os.IsExist(err)) && f.IsDir() {
		// working dir is on project root
		paths = append(paths, filepath.Join(hookDir, configFileName))
	} else {
		// work around for test
		paths = append(paths, configFileName)
	}
	return paths
}

func loadConfig(path string, cfg *globalConfig) *globalConfig {
	f, err := os.Open(path)
	if err != nil && !os.IsExist(err) {
		return cfg
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	if err := dec.Decode(cfg); err != nil {
		log.Println(err)
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

	for _, t := range Config.Types {
		TypeSet[t] = empty{}
	}

	for _, t := range Config.DenyTypes {
		delete(TypeSet, t)
	}

	types := make([]string, 0, len(TypeSet))
	for t := range TypeSet {
		types = append(types, t)
	}
	TypesStr = strings.Join(types, ", ")

	state.Init(Config.Lang, TypesStr)
}
