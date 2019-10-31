package _4chan

import (
	"daemon/functions"
	"daemon/structs"
	"daemon/vendors"
	"encoding/json"
	"strconv"
)

type vendor struct {
	request *vendors.Request
}

func (v *vendor) FetchThreads(boardName string) (threads []int, err error) {
	jsonData, err := v.request.Exec(boardName + "/" + "threads.json")
	if err != nil {
		return
	}

	var ThreadsStruct []struct {
		Threads []struct {
			Id int `json:"no"`
		}
	}

	if err = json.Unmarshal(jsonData, &ThreadsStruct); err != nil {
		return
	}

	for _, threadPage := range ThreadsStruct {
		for _, thread := range threadPage.Threads {
			threads = append(threads, thread.Id)
		}
	}

	return
}

func (v *vendor) FetchVideos(boardName string, threadId int) (videos []structs.Video, err error) {
	jsonData, err := v.request.Exec(boardName + "/" + "thread" + "/" + strconv.Itoa(threadId) + ".json")
	if err != nil {
		return
	}

	var PostsStruct struct {
		Posts []struct {
			Filename      string `json:"filename"`
			FileExtension string `json:"ext"`
		}
	}

	err = json.Unmarshal(jsonData, &PostsStruct)
	if err != nil {
		return
	}

	for _, post := range PostsStruct.Posts {
		if exists, _ := functions.InArray(post.FileExtension, vendors.AllowFileTypes); exists {
			videos = append(videos, structs.Video{
				ThreadId: threadId,
				Path:     "https://i.4cdn.org/" + boardName + "/" + post.Filename + post.FileExtension,
				Name:     post.Filename,
				Preview:  "https://i.4cdn.org/" + boardName + "/" + post.Filename + "s" + post.FileExtension,
			})
		}
	}

	return
}

func Instance() vendors.Interface {
	return &vendor{
		vendors.RequestFactory("https://a.4cdn.org"),
	}
}
