package implementation

import (
	"daemon/vendors"
	"fmt"
	"net/http"
	"strings"
)

type VendorImplementation struct {
	address string
}

func (vendor *VendorImplementation) buildUri(uri string) string {
	return vendor.address + "/" + strings.Trim(uri, "/")
}

func (vendor *VendorImplementation) execRequest(uri string) string {
	uri = vendor.buildUri(uri)

	response, error := http.Get(uri)
	if error != nil {
		// Fuck!
	}

	fmt.Printf("%v", response)
}

func (vendor *VendorImplementation) FetchThreads(boardName string) map[int]string {
	return map[int]string{}
}

func (vendor *VendorImplementation) FetchVideos(threadId int) map[int]string {
	return map[int]string{}
}

func TestFactory() vendors.Interface {
	return &VendorImplementation{address: "http://example.com"}
}
