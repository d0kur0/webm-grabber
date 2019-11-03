package _2ch

import (
	"daemon/functions"
	"daemon/structs"
	"daemon/vendors"
	"encoding/json"
	"path/filepath"
	"strconv"
)

type vendor struct {
	request *vendors.Request
}

func Instance() vendors.Interface {
	return &vendor{
		vendors.RequestFactory("http://2ch.hk"),
	}
}

func (vendor *vendor) FetchThreads(board string) (threads []int, err error) {
	jsonData, err := vendor.request.Exec(board + "/threads.json")
	if err != nil {
		return
	}

	var ThreadsStruct struct {
		Threads []struct {
			Id string `json:"num"`
		}
	}

	if err = json.Unmarshal(jsonData, &ThreadsStruct); err != nil {
		return
	}

	for _, thread := range ThreadsStruct.Threads {
		threadId, parseErr := strconv.ParseInt(thread.Id, 10, 64)
		if parseErr != nil {
			continue
		}

		threads = append(threads, int(threadId))
	}

	return
}

func (vendor *vendor) FetchFiles(board string, threadId int) (files []structs.File, err error) {
	jsonData, err := vendor.request.Exec(board + "/res/" + strconv.Itoa(threadId) + ".json")
	if err != nil {
		return
	}

	var PostsStruct struct {
		Threads []struct {
			Posts []struct {
				Files []struct {
					Name    string `json:"fullname"`
					Path    string `json:"path"`
					Preview string `json:"thumbnail"`
				}
			}
		}
	}

	if err = json.Unmarshal(jsonData, &PostsStruct); err != nil {
		return
	}

	for _, post := range PostsStruct.Threads[0].Posts {
		if post.Files == nil {
			continue
		}

		for _, file := range post.Files {
			if exists, _ := functions.InArray(filepath.Ext(file.Path), vendors.AllowFileTypes); exists {
				files = append(files, structs.File{
					ThreadId: threadId,
					Path:     "https://2ch.hk" + file.Path,
					Name:     file.Name,
					Preview:  "https://2ch.hk" + file.Preview,
				})
			}
		}
	}

	return
}
