package vendors

import "daemon/structs"

type Interface interface {
	FetchThreads(boardName string) ([]int, error)
	FetchVideos(boardName string, threadId int) (videos []structs.Video, err error)
}
