package vendors

type Interface interface {
	FetchThreads(boardName string) map[int]string
	FetchVideos(threadId int) map[int]string
}
