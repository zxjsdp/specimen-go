APP_NAME := specimen-go
APP_VERSION := v1.9.1

PLATFORMS := linux/amd64 linux/386 darwin/amd64 darwin/386 windows/amd64 windows/386
temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

help:
	@echo '=========================================================================='
	@echo 'Makefile for specimen-go'
	@echo ''
	@echo 'Usage:'
	@echo '    make build             compile specimen-go to executable'
	@echo '    make cross-compiling   cross-compiling specimen-go'
	@echo '    make windows-gui       compile specimen-go to Windows GUI executable'
	@echo '    make release           cross-compiling & windows-gui'
	@echo '    make clean             do cleaning job'
	@echo '=========================================================================='


.PHONY: build
build:
	@echo 'compile specimen-go to Windows GUI executable'
	go build -o specimen-go


.PHONY: cross-compiling $(PLATFORMS)
cross-compiling: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p bin
	if [ $(os) == "windows" ]; then \
		GOOS=$(os) GOARCH=$(arch) go build -o bin/$(APP_NAME)-$(APP_VERSION)-$(os)-$(arch).exe; \
	else \
		GOOS=$(os) GOARCH=$(arch) go build -o bin/$(APP_NAME)-$(APP_VERSION)-$(os)-$(arch); \
	fi


.PHONY: windows-gui
windows-gui:
	@echo 'compile specimen-go to Windows GUI executable'
	cd gui/windows
	if [[ ! -e rsrc.syso ]]; then rsrc -manifest specimen-go-gui.exe.manifest -o rsrc.syso -ico "../resources/icon.ico"; fi
	go build -ldflags="-H windowsgui" -o specimen-go.exe


.PHONY: release
release: cross-compiling windows-gui


.PHONY: clean
clean:
	@echo 'Start cleaning...'
	rm -rf ./bin/
	if [[ -e specimen-go ]] ; then rm specimen-go; echo 'specimen-go executable deleted!'; fi

