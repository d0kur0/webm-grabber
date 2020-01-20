package fourChannel

import (
	"daemon/sources/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ztrue/tracerr"
)

type implement struct {
	basedAddress      string
	allowedExtensions types.AllowedExtensions
}

func (vendor *implement) FetchThreads(board types.Board) (threads []types.Thread, err error) {
	response, err := http.Get(vendor.basedAddress + "/" + board.String() + "/threads.json")
	if err != nil {
		return
	}

	defer func() {
		err = response.Body.Close()
		if err != nil {
			tracerr.PrintSourceColor(tracerr.Wrap(err))
		}
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	var responseThreads ResponseThreads
	if err = json.Unmarshal(body, &responseThreads); err != nil {
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
	response, err := http.Get(vendor.basedAddress + "/" + thread.Board.String() + "/res/" + thread.StringId() + ".json")
	if err != nil {
		return
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		err = tracerr.New("Status code out of range 200-299")
		return
	}

	defer func() {
		err = response.Body.Close()
		if err != nil {
			tracerr.PrintSourceColor(tracerr.Wrap(err))
		}
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	var responsePosts ResponsePosts
	if err = json.Unmarshal(body, &responsePosts); err != nil {
		return
	}

	for _, post := range responsePosts.Posts {
		if !vendor.allowedExtensions.Contains(post.FileExtension) {
			continue
		}

		files = append(files, types.File{
			Name:     post.Filename,
			Path:     "https://i.4cdn.org/" + thread.Board.String() + "/" + fmt.Sprint(post.FileId) + post.FileExtension,
			Preview:  "https://i.4cdn.org/" + thread.Board.String() + "/" + fmt.Sprint(post.FileId) + "s" + post.FileExtension,
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
