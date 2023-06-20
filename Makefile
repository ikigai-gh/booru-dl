booru-dl: main.go
	go build .
lint:
	go fmt .
run:
	./booru-dl --tags "gothic_lolita katana" > /tmp/urls.txt
	./booru-dl --file /tmp/urls.txt
clean:
	rm -f /tmp/urls.txt /tmp/*.png /tmp/*.jpg
