package main

import (
	"daemon/structs"
	"daemon/vendors"
	_2ch "daemon/vendors/2ch"
	"log"
	"sync"

	"github.com/davecgh/go-spew/spew"
)

var vendorInstances = map[string]vendors.Interface{
	"2ch": _2ch.Instance(),
	//"4chan": _4chan.Instance(),
}

func main() {
	var filesChannel = make(chan structs.FileChannelMessage)
	var waitGroup sync.WaitGroup

	for _, localBoard := range GrabberSchema {
		for _, sourceBoard := range localBoard.SourceBoards {
			desiredVendor, vendorExists := vendorInstances[sourceBoard.Vendor]
			if !vendorExists {
				log.Println("Not found vendor", sourceBoard.Vendor)
				continue
			}

			threads, fetchThreadsErr := desiredVendor.FetchThreads(sourceBoard.Board)
			if fetchThreadsErr != nil {
				log.Println("Error FetchThreads:", fetchThreadsErr)
				continue
			}

			for _, threadId := range threads {
				waitGroup.Add(1)

				go func() {
					defer waitGroup.Done()

					threadFiles, fetchFilesErr := desiredVendor.FetchFiles(sourceBoard.Board, threadId)
					if fetchFilesErr != nil {
						log.Println("Error fetchFiles:", fetchFilesErr)
						return
					}

					filesChannel <- structs.FileChannelMessage{
						LocalBoard: localBoard.Name,
						Files:      threadFiles,
					}
				}()
			}
		}
	}

	waitGroup.Wait()
	spew.Dump(filesChannel)
}
