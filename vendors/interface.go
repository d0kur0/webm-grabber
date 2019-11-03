package vendors

import "daemon/structs"

type Interface interface {
	FetchThreads(board string) (threads []int, err error)
	FetchFiles(board string, threadId int) (files []structs.File, err error)
}
