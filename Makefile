ASSETS_FOLDER = pub/
WINDOWS_EXE_NAME = consulstate.exe
LINUX_EXE_NAME = consulstate_linux

.PHONY: build
build: test bindata
	go build

.PHONY: build_windows
build_windows: bindata
	GOOS=windows GOARCH=amd64 go build -o $(WINDOWS_EXE_NAME)

.PHONY: build_linux
build_linux: bindata
	GOOS=linux GOARCH=amd64 go build -o $(LINUX_EXE_NAME)

.PHONY: clean
clean:
	go clean
	-rm $(WINDOWS_EXE_NAME)

.PHONY: test
test: bindata
	go test -race -coverprofile=coverage.txt -covermode=atomic

.PHONY: test_all
test_all: bindata
	go test -race -coverprofile=coverage.txt -covermode=atomic -tags=integration

.PHONY: bindata
bindata:
	go-bindata $(ASSETS_FOLDER)
	