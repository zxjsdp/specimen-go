help:
	@echo '=================================================================='
	@echo 'Makefile for specimen-go'
	@echo ''
	@echo 'Usage:'
	@echo '    make install     compile specimen-go to executable'
	@echo '    make gui-win     compile specimen-go to Windows GUI executable'
	@echo '    make clean       do clean job'
	@echo '=================================================================='

install:
	@echo 'compile specimen-go to Windows GUI executable'
	go build -o specimen-go

gui-win:
	@echo 'compile specimen-go to Windows GUI executable'
	cd gui/windows
	if [[ ! -e rsrc.syso ]]; then rsrc -manifest specimen-go-gui.exe.manifest -o rsrc.syso -ico "../resources/icon.ico"; fi
	go build -ldflags="-H windowsgui" -o specimen-go.exe

clean:
	@echo 'Start cleaning...'
	if [[ -e specimen-go ]] ; then rm specimen-go; echo 'specimen-go executable deleted!'; fi
