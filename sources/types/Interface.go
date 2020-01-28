package types

type Interface interface {
	FetchThreads(board Board) (threads []Thread, err error)
	FetchFiles(thread Thread) (files []File, err error)
	VendorName() string
}
