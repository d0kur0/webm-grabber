package vendors

import "daemon/types"

type Interface interface {
	FetchThreads(board types.Board) (threads []types.Thread, err error)
	FetchFiles(thread types.Thread) (files []types.File, err error)
	VendorName() string
}
