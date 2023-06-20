booru-dl: main.go
	go build .
lint:
	go fmt .
run:
	./booru-dl posts -t "gothic_lolita katana" > /tmp/urls.txt
	./booru-dl posts -f /tmp/urls.txt
clean:
	rm -f /tmp/urls.txt /tmp/*.png /tmp/*.jpg
