GOFILES = $(shell find . -name '*.go')

default: build

workdir:
	mkdir -p workdir

build: workdir/tk

build-native: $(GOFILES)
	go build -o workdir/native-toolkit .

workdir/tk: $(GOFILES)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o workdir/toolkit .

clean:
	rm -rf workdir