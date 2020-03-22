APPNAME = commit-msg
LDFLAGS += -s -w

.PHONY: all build upx

all: build

build:
	go build -trimpath -ldflags '$(LDFLAGS)'

upx:
	upx `find -maxdepth 1 -name "$(APPNAME)" -o -name "$(APPNAME).exe"`
