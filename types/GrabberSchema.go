package types

import (
	"daemon/vendors"
)

type GrabberSchema struct {
	Vendor vendors.Interface
	Boards []Board
}
