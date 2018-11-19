# This is a Makefile to compile the WEB application
# Created by: DimiG
# Date: 2018-11-19

# GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
GO_BUILD=config.exe

all: clean
	$(GO_BUILD_ENV) go build -o $(GO_BUILD) .

clean:
	@[ -f $(GO_BUILD) ] && rm -v $(GO_BUILD) || true
