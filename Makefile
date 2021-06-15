APPNAME := commit-msg
LDFLAGS += -s -w

.PHONY: all gen build upx

all: gen build upx

gen:
	go generate ./...

build:
	go build -trimpath -ldflags "$(LDFLAGS) -X 'main.version=$(TAG)' -X 'main.goVersion=$(shell go version)' -X 'main.commitHash=$(shell git show -s --format=%H)' -X 'main.buildTime=$(shell date "+%Y-%m-%d %T%z")'"

upx:
	upx `find -maxdepth 1 -name "$(APPNAME)" -o -name "$(APPNAME).exe"`
