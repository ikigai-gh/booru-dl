package booru

import (
    "strings"
    "strconv"
    "net/url"
	"encoding/json"
)

type Api struct {
    BaseUrl string
    PostsPerPageLimit int
    TagLimit int
}

func (a *Api) Posts(tagString string, useLargeFileBaseUrls bool, maxPages int) ([]Post, error) {
	if len(strings.Split(tagString, " ")) > a.TagLimit {
		panic("Only " + strconv.Itoa(a.TagLimit) + " tags are allowed!")
	}
    postsUrl := a.BaseUrl + "posts.json"

    query := url.Values{}
    query.Add("tags", tagString)
    data, err := Request(postsUrl + "?" + query.Encode())

    if err != nil {
        return nil, err
    }

    var posts []Post
    // TODO: Move unmarshalling logic to method
    // var urlString string
    // if useLargeFileUrls {
    //    urlString = p.LargeFileUrl
    // } else {
    //    urlString = p.PreviewFileUrl
    // }
    // Search only for pics
    //if urlString != "" && (strings.HasSuffix(urlString, ".jpg") || strings.HasSuffix(urlString, ".png")) {
    //    fmt.Println(urlString)
    //}
    err = json.Unmarshal(data, &posts)

    if err != nil {
        return nil, err
    }

    return posts, nil
}

func (a *Api) Tags() []Tag {
    return nil
}
