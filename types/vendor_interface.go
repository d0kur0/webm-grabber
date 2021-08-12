package types

type VendorInterface interface {
	FetchThreads(board Board) (threads []Thread, err error)
	FetchFiles(thread Thread) (files []File, err error)
	GetThreadUrl(thread Thread) (url string)
	VendorName() string
}
