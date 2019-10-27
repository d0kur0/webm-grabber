package implementation

import (
	"daemon/vendors"
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
)

type vendor struct {
	request *vendors.Request
}

func (v *vendor) FetchThreads(boardName string) (response map[int]string, err error) {
	jsonData, err := v.request.Exec(boardName + "/catalog_num.json")
	if err != nil {
		return
	}

	var result map[string]interface{}
	json.Unmarshal(jsonData, &result)

	threads := result["threads"].(map[string]interface{})

	spew.Dump(threads)

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
