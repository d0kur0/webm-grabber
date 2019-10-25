package implementation

import (
	"daemon/vendors"
	"fmt"
	"net/http"
	"strings"
)

type VendorImplementation struct {
	request vendors.Request{ address: "http://example.com" }
}

func (vendor *VendorImplementation) FetchThreads(boardName string) map[int]string {
	return map[int]string{}
}

func (vendor *VendorImplementation) FetchVideos(threadId int) map[int]string {
	return map[int]string{}
}

func TestFactory() vendors.Interface {
	return &VendorImplementation{}
}
