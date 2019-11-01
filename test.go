package main

type ThreadChannelMessage struct {
	ThreadId  int
	BoardName string
	Vendor    string
}

type File struct {
	Name string
	Path string
}

type Vendor struct {
	Name string
}

func (vendor *Vendor) FetchFiles (boardName string, threadId int) (files []File) {
	for i := 0; i <= 100; i++ {

	}
}

var vendors = map[string][]Vendor{
	"2ch": { Vendor{"2ch.hk"} },
}

func main() {
	var threadsChannel = make(chan []ThreadChannelMessage)
	var filesChannel   = make(chan []File)

	// Getting threads
	go asyncThreads(threadsChannel)

	// Catch new threads in channel
	go func (threadsChannel chan) {

		for {
			newThreadId := <- threadsChannel
			go asyncVideos(newThreadId, filesChannel)
		}

	}(threadsChannel)
}

func asyncThreads(threadChannel chan) {
	for i := 0; i <= 100; i++ {
		threadChannel <- ThreadChannelMessage{
			ThreadId:	 i,
			BoardName: 	"b",
			Vendor: 	"2ch",
		}
	}
}

func asyncVideos(thread ThreadChannelMessage, filesChannel chan) {
	if desiredVendor, vendorExists := vendors[thread.Vendor]; vendorExists {
		filesChannel <- desiredVendor.FetchFiles(thread.BoardName, thread.ThreadId)
	}
}
