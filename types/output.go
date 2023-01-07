package types

type OutputItem struct {
	VendorName   string `json:"vendorName"`
	BoardName    string `json:"boardName"`
	SourceThread string `json:"sourceThread"`
	File         `json:"file"`
}

type Output = []OutputItem
