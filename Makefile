ASSETS_FOLDER = pub/
WINDOWS_EXE_NAME = consulstate.exe

.PHONY: build
build: test bindata
	go build

.PHONY: build_windows
build_windows: bindata
	GOOS=windows GOARCH=amd64 go build -o $(WINDOWS_EXE_NAME)

.PHONY: clean
clean:
	go clean
	-rm $(WINDOWS_EXE_NAME)

.PHONY: test
test: bindata
	go test

.PHONY: bindata
bindata:
	go-bindata $(ASSETS_FOLDER)
	