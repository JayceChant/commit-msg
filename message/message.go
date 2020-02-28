package message

var (
	// Lang ...
	Lang *LangPack
	// Types ...
	Types string
)

// Config ...
func Config(l *LangPack, t string) {
	Lang = l
	Types = t
}
