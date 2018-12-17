# This is a Makefile to compile the application
# Just help to Automate the compilation process
# With `make', `make macos' and `make clean'
# Created by: DimiG
# Date: 2018-12-17

### VARIABLEs ###
### GO_BUILD_ENV_M and GO_BUILD_ENV_L MacOS and Linux related ###
TARGET = videotomp4
TARGET_W = $(TARGET).exe
COMPRESSOR_W = upx.exe
WIN_RES = main-res.syso
GO_BUILD_ENV_M = GOOS=darwin GOARCH=amd64
GO_BUILD_ENV_L = CGO_ENABLED=0 GOOS=linux GOARCH=amd64

ifeq ($(OS),Windows_NT)
  DETECTED_OS = Windows
else
  DETECTED_OS = $(shell sh -c 'uname 2>/dev/null || echo Unknown')
endif

.PHONY: all win mac clean

all: clean
ifeq ($(DETECTED_OS),Windows)
	$(info =>>> Compilation for Windows is in progress...)
	windres -i main.rc -O coff -o $(WIN_RES)
	go build -o $(TARGET_W) .
endif

ifeq ($(detected_OS),Darwin)
	$(info =>>> Compilation for Mac OS is in progress...)
	go build -o $(TARGET) .
endif

ifeq ($(detected_OS),Linux)
	$(info =>>> Compilation for Linux is in progress...)
	go build -o $(TARGET) .
endif

win: clean
	$(info =>>> Optimized for Windows platform)
	windres -i main.rc -O coff -o $(WIN_RES)
	$(GO_BUILD_ENV_W) go build -ldflags="-s -w" -o $(TARGET_W) .
	$(COMPRESSOR_W) --brute $(TARGET_W)

mac:
	$(info =>>> Optimized for MacOS platform)
	$(GO_BUILD_ENV_M) go build -ldflags="-s -w" .

clean:
	$(info =>>> Cleaning UP)
	@[ -f $(WIN_RES) ] && rm -v $(WIN_RES) || true		# Windows related
	@[ -f $(TARGET_W) ] && rm -v $(TARGET_W) || true	# Windows related
	@[ -f $(TARGET) ] && rm -v $(TARGET) || true		# MacOS and Linux related
