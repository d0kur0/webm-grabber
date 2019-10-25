package vendors

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	Address string
}

func (r *Request) BuildUri(uri string) string {
	return r.Address + "/" + strings.Trim(uri, "/")
}

func (r *Request) Exec(uri string) map[string]interface{} {
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

	var jsonData map[string]interface{}
	if error := json.Unmarshal(body, &jsonData); error != nil {
		panic(error)
	}

	return jsonData
}
