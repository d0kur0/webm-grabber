package twoChannel

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"webmGrabber/sources/types"

	"github.com/ztrue/tracerr"
)

type implement struct {
	basedAddress      string
	allowedExtensions types.AllowedExtensions
	authToken         string
}

func (vendor *implement) request(url string) (responseData []byte, err error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	request.Header.Set("Cookie", "usercode_auth="+vendor.authToken+"; path=/; domain=.2ch.hk;")

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		return
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		err = tracerr.New("Status code out of range 200-299 for url: " + url)
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
	response, err := vendor.request(vendor.basedAddress + "/" + thread.Board.String() + "/res/" + thread.StringId() + ".json")
	if err != nil {
		return
	}

	var responsePosts ResponsePosts
	if err = json.Unmarshal(response, &responsePosts); err != nil {
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
	return &implement{"https://2ch.hk", allowedExtensions, "5c46087a5952919e3740736f355b0515"}
}
