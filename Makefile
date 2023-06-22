GOOS ?= linux
GOARCH ?= amd64
VERSION = $(shell git describe --tags)
BINARY = booru-dl-${VERSION}-${GOOS}-${GOARCH}
LDFLAGS = -ldflags="-s -w -extldflags=-static -X github.com/ikigai-gh/booru-dl/cmd/booru.version=${VERSION}"

booru-dl:
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=0 go build ${LDFLAGS} -o ${BINARY}
lint:
	go fmt .
run:
	./${BINARY} posts -t "gothic_lolita katana" > /tmp/urls.txt
	./${BINARY} posts -f /tmp/urls.txt
clean:
	rm -f /tmp/urls.txt /tmp/*.png /tmp/*.jpg
