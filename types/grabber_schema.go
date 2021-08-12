package types

type GrabberSchema struct {
	Vendor VendorInterface `json:"vendor"`
	Boards []Board         `json:"boards"`
}
