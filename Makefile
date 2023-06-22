GOOS ?= linux
GOARCH ?= amd64
BINARY = booru-dl-${GOOS}-${GOARCH}
LDFLAGS = -ldflags="-s -w -extldflags=-static"

booru-dl:
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=0 go build ${LDFLAGS} -o ${BINARY}
lint:
	go fmt .
run:
	./booru-dl posts -t "gothic_lolita katana" > /tmp/urls.txt
	./booru-dl posts -f /tmp/urls.txt
clean:
	rm -f /tmp/urls.txt /tmp/*.png /tmp/*.jpg
