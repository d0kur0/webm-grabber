package fourChannel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"webmGrabber/sources/types"

	"github.com/ztrue/tracerr"
)

type implement struct {
	basedAddress      string
	allowedExtensions types.AllowedExtensions
}

func (vendor *implement) request(url string) (responseData []byte, err error) {
	response, err := http.Get(url)
	if err != nil {
		return
	}

	defer func() {
		err = response.Body.Close()
		if err != nil {
			tracerr.PrintSourceColor(tracerr.Wrap(err))
		}
	}()

	return ioutil.ReadAll(response.Body)
}

func (vendor *implement) FetchThreads(board types.Board) (threads []types.Thread, err error) {
	response, err := vendor.request(vendor.basedAddress + "/" + board.String() + "/threads.json")
	if err != nil {
		return
	}

	var responseThreads ResponseThreads
	if err = json.Unmarshal(response, &responseThreads); err != nil {
		return
	}

	for _, threadPage := range responseThreads {
		for _, thread := range threadPage.Threads {
			threads = append(threads, types.Thread{
				ID:    thread.Id,
				Board: board,
			})
		}
	}

	return
}

func (vendor *implement) FetchFiles(thread types.Thread) (files []types.File, err error) {
	response, err := vendor.request(vendor.basedAddress + "/" + thread.Board.String() + "/res/" + thread.StringId() + ".json")
	if err != nil {
		return
	}

	var responsePosts ResponsePosts
	if err = json.Unmarshal(response, &responsePosts); err != nil {
		return
	}

	for _, post := range responsePosts.Posts {
		if !vendor.allowedExtensions.Contains(post.FileExtension) {
			continue
		}

		files = append(files, types.File{
			Name:     post.Filename,
			Path:     "https://i.4cdn.org/" + thread.Board.String() + "/" + fmt.Sprint(post.FileId) + post.FileExtension,
			Preview:  "https://i.4cdn.org/" + thread.Board.String() + "/" + fmt.Sprint(post.FileId) + "s.png",
			ThreadId: thread.ID,
		})
	}

	return
}

func (vendor *implement) VendorName() string {
	return "4chan"
}

func Make(allowedExtension types.AllowedExtensions) types.Interface {
	return &implement{"https://a.4cdn.org", allowedExtension}
}
