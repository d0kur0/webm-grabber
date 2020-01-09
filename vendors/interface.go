package vendors

import "daemon/types"

type Interface interface {
	FetchThreads(board string) (threads []int, err error)
	FetchFiles(board string, threadId int) (files []types.File, err error)
}
