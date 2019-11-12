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

	for _, localBoard := range getGrabberSchema() {
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

				go func(vendor vendors.Interface, board string, threadId int, wg *sync.WaitGroup) {
					defer wg.Done()

					threadFiles, fetchFilesErr := vendor.FetchFiles(board, threadId)
					if fetchFilesErr != nil {
						log.Println("Error fetchFiles:", fetchFilesErr)
						return
					}

					filesChannel <- structs.FileChannelMessage{
						LocalBoard: localBoard.Name,
						Files:      threadFiles,
					}
				}(desiredVendor, localBoard.Name, threadId, &waitGroup)
			}
		}
	}

	for file := range filesChannel {
		spew.Dump(file)
	}

	waitGroup.Wait()
	log.Println("All operations end")
}

func getGrabberSchema() (grabberSchema []structs.Board) {
	grabberSchema = []structs.Board{
		{
			Name:        "b",
			Description: "...",
			SourceBoards: []structs.SourceBoard{
				{"2ch", "b"},
				{"4chan", "b"},
			},
		},
		{
			Name:        "a",
			Description: "...",
			SourceBoards: []structs.SourceBoard{
				{"2ch", "a"},
				{"4chan", "a"},
				{"4chan", "c"},
			},
		},
		{
			Name:        "s",
			Description: "...",
			SourceBoards: []structs.SourceBoard{
				{"4chan", "s"},
				{"4chan", "c"},
			},
		},
	}

	return
}
