package vendors

import (
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

func (r *Request) Exec(uri string) (body []byte, err error) {
	uri = r.BuildUri(uri)

	response, err := http.Get(uri)
	if err != nil {
		return
	}

	defer func() {
		er := response.Body.Close()
		if er != nil {
			// TODO: Catch error
		}
	}()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	return
}

func RequestFactory(address string) *Request {
	return &Request{address}
}
