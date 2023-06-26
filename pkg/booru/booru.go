package booru

import (
	"bufio"
	"encoding/json"
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
const TagsStr = "tags.json?limit=200"

type Post struct {
	LargeFileUrl   string `json:"large_file_url"`
	PreviewFileUrl string `json:"preview_file_url"`
}

type Tag struct {
    Name string `json:"name"`
}

type Environment string

func try(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func GetPosts(tagString string, useLargeFileUrls bool, urlsFile string, maxPages int) {
	// download images from file
	if urlsFile != "" {
		var urls []string
		file, err := os.Open(urlsFile)
		defer file.Close()
		try(err)

		sc := bufio.NewScanner(file)
		for sc.Scan() {
			urls = append(urls, sc.Text())
		}

		b := bar.NewOptions(len(urls), bar.OptionSetWidth(50), bar.OptionSetDescription("downloading images..."))

        fmt.Println("Number of images to download: " + strconv.Itoa(len(urls)))
		var wg sync.WaitGroup
		wg.Add(len(urls))

		for idx, u := range urls {
			ext := "." + strings.Split(u, ".")[len(strings.Split(u, "."))-1]
            filePath := "/tmp/" + strconv.Itoa(idx) + ext
            go func(url string, wg *sync.WaitGroup) {
                bytesResp, err := Request(url)
                try(err)
                file, err := os.Create(filePath)
                defer file.Close()
                try(err)
                _, err = file.Write(bytesResp)
                try(err)
                wg.Done()
                }(u, &wg)
			b.Add(1)
			time.Sleep(5 * time.Millisecond)
		}

		wg.Wait()
		// just print urls to stdout
	} else {
        // TODO: Add remaining sites
        a := Api{BaseUrl: "https://danbooru.donmai.us/", PostsPerPageLimit: 200, TagLimit: 2}
        posts, err := a.Posts(tagString, useLargeFileUrls, maxPages)
        try(err)
        for _, p := range posts {
            fmt.Println(p.PreviewFileUrl)
        }
    }
}

func GetTags() {
    pageNumber := 1
    for {
        tagsResp, err := http.Get("" + "/" + TagsStr + "&page=" + strconv.Itoa(pageNumber))
        try(err)

        defer tagsResp.Body.Close()

        tagsBytes, err := io.ReadAll(tagsResp.Body)
        try(err)

        var tags []Tag
        err = json.Unmarshal(tagsBytes, &tags)
        if len(tags) == 0 {
            break
        }
        try(err)
        for _, t := range tags {
            fmt.Println(t)
        }
        pageNumber++
    }
}
