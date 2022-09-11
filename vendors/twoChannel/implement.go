package twoChannel

import (
	"encoding/json"
	"fmt"
	"github.com/d0kur0/webm-grabber/types"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type twoChannel struct {
	CDNBaseAddress    string
	boardsBaseAddress string
	allowedExtensions types.AllowedExtensions
	authToken         string
}

func (vendor *twoChannel) request(url string) (responseData []byte, err error) {
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
		err = errors.New("Request returned code out of range 200-299, url:" + url)
		return
	}

	defer func() {
		_ = response.Body.Close()
	}()

	return ioutil.ReadAll(response.Body)
}

func (vendor *twoChannel) FetchThreads(board types.Board) (threads []types.Thread, err error) {
	url := fmt.Sprintf("https://%s/%s/threads.json", vendor.CDNBaseAddress, board.Name)
	response, err := vendor.request(url)
	if err != nil {
		return
	}

	var responseThreads ResponseThreads
	if err = json.Unmarshal(response, &responseThreads); err != nil {
		return
	}

	for _, thread := range responseThreads.Threads {
		threads = append(threads, types.Thread{
			ID:    thread.Id,
			Board: board,
		})
	}

	return
}

func (vendor *twoChannel) FetchFiles(thread types.Thread) (files []types.File, err error) {
	url := fmt.Sprintf("https://%s/%s/res/%d.json", vendor.CDNBaseAddress, thread.Board.Name, thread.ID)
	response, err := vendor.request(url)
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
				Path:     fmt.Sprintf("https://%s/%s", vendor.CDNBaseAddress, file.Path),
				Preview:  fmt.Sprintf("https://%s/%s", vendor.CDNBaseAddress, file.Preview),
				ThreadId: thread.ID,
			})
		}
	}

	return
}

func (vendor *twoChannel) VendorName() string {
	return "2ch"
}

func (vendor *twoChannel) GetThreadUrl(thread types.Thread) string {
	return fmt.Sprintf("https://%s/%s/res/%d.html", vendor.boardsBaseAddress, thread.Board.Name, thread.ID)
}

func Make(allowedExtensions types.AllowedExtensions) types.VendorInterface {
	return &twoChannel{
		"2ch.hk",
		"2ch.hk",
		allowedExtensions,
		"5c46087a5952919e3740736f355b0515",
	}
}
