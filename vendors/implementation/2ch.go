package implementation

import (
	"daemon/vendors"
	"encoding/json"
)

type vendor struct {
	request *vendors.Request
}

func (v *vendor) FetchThreads(boardName string) (response map[int]string, err error) {
	jsonData, err := v.request.Exec(boardName + "/catalog_num.json")
	if err != nil {
		return
	}

	var jsonResponse struct {
		Threads interface{} `json:"threads"`
	}

	if err = json.Unmarshal(jsonData, &jsonResponse); err != nil {
		return
	}
	//spew.Dump(jsonResponse.Threads)

	return
}

func (v *vendor) FetchVideos(threadId int) (response map[int]string, err error) {
	return
}

func Instance2ch() vendors.Interface {
	return &vendor{
		vendors.RequestFactory("http://2ch.hk"),
	}
}
