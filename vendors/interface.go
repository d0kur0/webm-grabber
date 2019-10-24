package vendors

type Interface interface {
	execRequest(uri string) string
	buildUri(uri string) string
	FetchThreads(boardName string) map[int]string
	FetchVideos(threadId int) map[int]string
}
