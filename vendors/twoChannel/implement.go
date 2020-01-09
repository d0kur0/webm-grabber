package _2ch

import (
	"daemon/types"
	"daemon/types/twoChannel"
	"daemon/vendors"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ztrue/tracerr"
)

type implement struct {
	basedAddress        string
	extensionsForFilter []string
}

func (vendor *implement) FetchThreads(board string) (threads []int, err error) {
	response, err := http.Get(vendor.basedAddress + "/" + board + "/threads.json")
	if err != nil {
		return
	}

	defer func() {
		err = response.Body.Close()
		tracerr.PrintSourceColor(tracerr.Wrap(err))
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	var responseThreads twoChannel.ResponseThreads
	if err = json.Unmarshal(body, &responseThreads); err != nil {
		return
	}

	for _, thread := range responseThreads.Threads {
		threadId, convertError := strconv.Atoi(thread.Id)
		if convertError != nil {
			tracerr.PrintSourceColor(tracerr.Wrap(convertError))
		}

		threads = append(threads, threadId)
	}

	return
}

func (vendor *implement) FetchFiles(board string, threadId int) (files []types.File, err error) {
	response, err := http.Get(vendor.basedAddress + "/" + board + "/" + strconv.Itoa(threadId) + ".json")
	if err != nil {
		return
	}

	defer func() {
		err = response.Body.Close()
		tracerr.PrintSourceColor(tracerr.Wrap(err))
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	var responsePosts twoChannel.ResponsePosts
	if err = json.Unmarshal(body, &responsePosts); err != nil {
		return
	}

	for _, post := range responsePosts.Threads[0].Posts {
		if post.Files == nil {
			continue
		}

		for _, file := range post.Files {

		}
	}

	return
}

func Make(extensionForFilter []string) vendors.Interface {
	return &implement{"http://2ch.hk", extensionForFilter}
}
