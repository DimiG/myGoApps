# This is a Makefile to compile the application
# Just help to Automate the compilation process
# With `make', `make macos' and `make clean'
# Created by: DimiG
# Date: 2018-12-12

### VARIABLEs ###
### GO_BUILD_ENV_M and GO_BUILD_ENV_L MacOS and Linux related ###
TARGET = videotomp4
TARGET_W = $(TARGET).exe
WIN_RES = main-res.syso
GO_BUILD_ENV_M = GOOS=darwin GOARCH=amd64
GO_BUILD_ENV_L = CGO_ENABLED=0 GOOS=linux GOARCH=amd64

.PHONY: all macos clean

all: clean
	windres -i main.rc -O coff -o $(WIN_RES)
	$(GO_BUILD_ENV_W) go build -o $(TARGET_W) .

macos:
	$(GO_BUILD_ENV_M) go build .

clean:
	@[ -f $(TARGET_W) ] && rm -v $(TARGET_W) $(WIN_RES) || true	# Windows related
	@[ -f $(TARGET) ] && rm -v $(TARGET) || true	# MacOS and Linux related
