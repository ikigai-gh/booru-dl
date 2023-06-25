# booru-dl
View, search &amp; download your awesome waifus from plenty of boards

## List of currently supported boards

- [x] [Danbooru](https://danbooru.donmai.us)

## Build (linux/amd64)
Compile this utility at least with go 1.20:

```bash
cd booru-dl && make
```

For other operating systems and architectures set `GOOS` and `GOARCH` environment variables correspondently.

## Run
This utility supports two modes: one for just printing out links to stdout, like this:

```bash
./booru-dl posts -t "katana open_mouth" > /tmp/urls.txt
```
By default booru-dl will print urls of **preview** images. You can change this behaviour by setting `-l` flag.

To download images from file containing urls just pass the `-f` flag like this:
```bash
./booru-dl posts -f /tmp/urls.txt
```
