# booru-dl
View, search &amp; download your awesome waifus from plenty of boards

## List of currently supported boards

- [x] [Danbooru](https://danbooru.donmai.us)

## Build
Compile this utility at least with go 1.20:

```bash
    cd booru-dl && go build
```

## Run
This utility supports two modes: one for just printing out links to stdout, like this:

```bash
    booru-dl --tags "katana open_mouth" > /tmp/urls.txt
```
By default booru-dl will print urls of **preview** images. You can change this behaviour by setting `--large` flag.

To download images from file containing urls just pass the `--file` flag like this:
```bash
    booru-dl --file /tmp/urls.txt
```
