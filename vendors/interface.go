package vendors

type Interface interface {
	FetchThreads(boardName string) (response map[int]string, err error)
	FetchVideos(threadId int) (response map[int]string, err error)
}
