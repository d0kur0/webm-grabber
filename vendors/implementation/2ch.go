package implementation

import (
	"daemon/vendors"
)

type VendorImplementation struct {
	request vendors.Request
}

func (vendor *VendorImplementation) FetchThreads(boardName string) map[int]string {
	return map[int]string{}
}

func (vendor *VendorImplementation) FetchVideos(threadId int) map[int]string {
	return map[int]string{}
}

func Instance2ch() vendors.Interface {
	return &VendorImplementation{
		vendors.Request{Address: "http://example.com"},
	}
}
