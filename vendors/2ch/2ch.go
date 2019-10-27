package _2ch

import (
	"daemon/vendors"
	"encoding/json"
	"strconv"
)

type vendor struct {
	request *vendors.Request
}

func (v *vendor) FetchThreads(boardName string) (threads []int, err error) {
	jsonData, err := v.request.Exec(boardName + "/catalog_num.json")
	if err != nil {
		return
	}

	var ThreadsStruct struct {
		Threads []struct {
			Id string `json:"num"`
		}
	}

	err = json.Unmarshal(jsonData, &ThreadsStruct)
	if err != nil {
		return
	}

	for _, thread := range ThreadsStruct.Threads {
		threadId, err := strconv.ParseInt(thread.Id, 10, 64)
		if err != nil {
			continue
		}

		threads = append(threads, int(threadId))
	}

	return
}

func (v *vendor) FetchVideos(threadId int) (response map[int]string, err error) {
	return
}

func Instance() vendors.Interface {
	return &vendor{
		vendors.RequestFactory("http://2ch.hk"),
	}
}
