package types

type ChannelMessage struct {
	VendorName   string `json:"vendorName"`
	Files        []File `json:"files"`
	Thread       Thread `json:"thread"`
	SourceThread string `json:"sourceThread"`
}
