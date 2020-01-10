package daemon

import (
	"daemon/types"
	"daemon/vendors"
)

type GrabberSchema struct {
	Vendor vendors.Interface
	Boards []types.Board
}
