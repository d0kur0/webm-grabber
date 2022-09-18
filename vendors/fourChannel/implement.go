package fourChannel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/d0kur0/webm-grabber/types"

	"github.com/pkg/errors"
)

type fourChannel struct {
	CDNBaseAddress    string
	boardsBaseAddress string
	allowedExtensions types.AllowedExtensions
}

const VendorName = "4chan"

func (vendor *fourChannel) request(url string) (responseData []byte, err error) {
	response, err := http.Get(url)
	if err != nil {
		return
	}

	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode < 200 || response.StatusCode > 299 {
		err = errors.New("Request returned code out of range 200-299, url:" + url)
		return
	}

	return ioutil.ReadAll(response.Body)
}

func (vendor *fourChannel) FetchThreads(board types.Board) (threads []types.Thread, err error) {
	url := fmt.Sprintf("https://a.%s/%s/threads.json", vendor.CDNBaseAddress, board.Name)
	response, err := vendor.request(url)
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

func (vendor *fourChannel) FetchFiles(thread types.Thread) (files []types.File, err error) {
	url := fmt.Sprintf("https://a.%s/%s/res/%d.json", vendor.CDNBaseAddress, thread.Board.Name, thread.ID)
	response, err := vendor.request(url)
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
			Path:     fmt.Sprintf("https://i.%s/%s/%d%s", vendor.CDNBaseAddress, thread.Board.Name, post.FileId, post.FileExtension),
			Preview:  fmt.Sprintf("https://i.%s/%s/%d%s", vendor.CDNBaseAddress, thread.Board.Name, post.FileId, "s.jpg"),
			ThreadId: thread.ID,
		})
	}

	return
}

func (vendor *fourChannel) VendorName() string {
	return VendorName
}

func (vendor *fourChannel) GetThreadUrl(thread types.Thread) string {
	return fmt.Sprintf("https://%s/%s/thread/%d", vendor.boardsBaseAddress, thread.Board.Name, thread.ID)
}

func Make(allowedExtension types.AllowedExtensions) types.VendorInterface {
	return &fourChannel{
		"4cdn.org",
		"boards.4channel.org",
		allowedExtension,
	}
}
