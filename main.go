package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	bar "github.com/schollz/progressbar/v3"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const TestDomain = "https://testbooru.donmai.us"
const ProdDomain = "https://danbooru.donmai.us"

// 200 is a maximum per page on danbooru
const PostLimit = 200

// 2 is a maximum tag limit on danbooru
const TagLimit = 2
const PostsStr = "posts.json?limit=200"

// Command line flags
var tagString = flag.String("tags", "", "list of tags, separated by space")
var useLargeFileUrls = flag.Bool("large", false, "whether to save full sized image urls")
var urlsFile = flag.String("file", "", "file with links to images to download")

type Post struct {
	LargeFileUrl   string `json:"large_file_url"`
	PreviewFileUrl string `json:"preview_file_url"`
}

type Environment string

const (
	DEV  Environment = "DEV"
	PROD Environment = "PROD"
)

func try(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func downloadImg(url string, filePath string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	try(err)
	defer resp.Body.Close()
	bytesResp, err := io.ReadAll(resp.Body)
	file, err := os.Create(filePath)
	defer file.Close()
	try(err)
	_, err = file.Write(bytesResp)
	try(err)
	wg.Done()
}

func main() {
	env := Environment(os.Getenv("BOORU_ENV"))
	var url string
	if env == DEV {
		url = TestDomain
	} else {
		url = ProdDomain
	}

	flag.Parse()

	if len(strings.Split(*tagString, " ")) > TagLimit {
		panic("Only " + strconv.Itoa(TagLimit) + " tags are allowed!")
	}

	// download images from file
	if *urlsFile != "" {
		var urls []string
		file, err := os.Open(*urlsFile)
		defer file.Close()
		try(err)

		sc := bufio.NewScanner(file)
		for sc.Scan() {
			urls = append(urls, sc.Text())
		}

		b := bar.NewOptions(len(urls), bar.OptionSetWidth(50), bar.OptionSetDescription("downloading images..."))

		var wg sync.WaitGroup
		wg.Add(len(urls))

		for idx, u := range urls {
			ext := "." + strings.Split(u, ".")[len(strings.Split(u, "."))-1]
			fileName := strconv.Itoa(idx) + ext
			go downloadImg(u, "/tmp/"+fileName, &wg)
			b.Add(1)
			time.Sleep(5 * time.Millisecond)
		}

		wg.Wait()
		// just print urls to stdout
	} else {
		pageNumber := 1
		for {
			postsResp, err := http.Get(url + "/" + PostsStr + "&page=" + strconv.Itoa(pageNumber) + "&tags=" + *tagString)
			try(err)

			defer postsResp.Body.Close()

			postsBytes, err := io.ReadAll(postsResp.Body)
			try(err)

			var posts []Post
			err = json.Unmarshal(postsBytes, &posts)
			if len(posts) == 0 {
				break
			}
			try(err)
			for _, p := range posts {
				var urlString string
				if *useLargeFileUrls {
					urlString = p.LargeFileUrl
				} else {
					urlString = p.PreviewFileUrl
				}
				// Search only for pics
				if strings.HasSuffix(urlString, ".jpg") || strings.HasSuffix(urlString, ".png") {
					fmt.Println(urlString)
				}
			}
			pageNumber++
		}
	}
}
