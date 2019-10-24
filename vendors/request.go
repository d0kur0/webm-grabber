package vendors

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	address string
}

func (r *Request) BuildUri(uri string) string {
	return r.address + "/" + strings.Trim(uri, "/")
}

func (r *Request) Exec(uri string) byte {
	uri = r.BuildUri(uri)

	response, error := http.Get(uri)
	if error != nil {
		panic(error)
	}

	defer response.Body.Close()
	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		panic(error)
	}

	var jsonData = map[string]interface{}
	if error := json.Unmarshal(response, jsonData); error != nil {
		panic(error)
	}

}
