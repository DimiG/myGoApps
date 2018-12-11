# This is a Makefile to compile the application
# Just help to Automate the compilation process
# With `make' and `make clean'
# Created by: DimiG
# Date: 2018-12-11

# GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
# GO_BUILD example for Windows platform
# For Linux or Mac OS .exe IS NOT necessary
GO_BUILD=videotomp4.exe

all: clean
	windres -i main.rc -O coff -o main-res.syso
	$(GO_BUILD_ENV) go build -o $(GO_BUILD) .

clean:
	@[ -f $(GO_BUILD) ] && rm -v $(GO_BUILD) main-res.syso || true

