linux:
	go build -o build/linux_5words
windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -ldflags="-H windowsgui" -o build/windows_5words
mac:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o build/mac_5words

embed:
	pkger -include /webnode/dist

.PHONY: embed linux windows mac
