package twoChannel

import (
	"daemon/sources/types"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"

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

	for _, thread := range responseThreads.Threads {
		threadId, convertError := strconv.Atoi(thread.Id)
		if convertError != nil {
			tracerr.PrintSourceColor(tracerr.Wrap(convertError))
		}

		threads = append(threads, types.Thread{
			ID:    threadId,
			Board: board,
		})
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

	for _, post := range responsePosts.Threads[0].Posts {
		if post.Files == nil {
			continue
		}

		for _, file := range post.Files {
			if !vendor.allowedExtensions.Contains(filepath.Ext(file.Path)) {
				continue
			}

			files = append(files, types.File{
				Name:     file.Name,
				Path:     vendor.basedAddress + "/" + file.Path,
				Preview:  vendor.basedAddress + "/" + file.Preview,
				ThreadId: thread.ID,
			})
		}
	}

	return
}

func (vendor *implement) VendorName() string {
	return "2ch"
}

func Make(allowedExtensions types.AllowedExtensions) types.Interface {
	return &implement{"http://2ch.hk", allowedExtensions}
}
