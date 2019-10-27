package vendors

type Interface interface {
	FetchThreads(boardName string) ([]int, error)
	FetchVideos(threadId int) (response map[int]string, err error)
}
