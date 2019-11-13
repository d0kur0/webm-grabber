package main

import (
	"daemon/structs"
	"daemon/vendors"
	_2ch "daemon/vendors/2ch"
	_4chan "daemon/vendors/4chan"
	"log"
	"sync"
)

func main() {
	var vendorInstances = map[string]vendors.Interface{
		"2ch":   _2ch.Instance(),
		"4chan": _4chan.Instance(),
	}

	var filesChannel = make(chan structs.FileChannelMessage)
	var response = make(map[string][]structs.File)
	var grabberSchema = getGrabberSchema()
	var waitGroup sync.WaitGroup

	go func() {
		for {
			file := <-filesChannel

			if _, exists := response[file.LocalBoard]; exists {
				response[file.LocalBoard] = append(response[file.LocalBoard], file.Files...)
			}
		}
	}()

	for _, localBoard := range grabberSchema {
		response[localBoard.Name] = nil

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

				go func(vendor vendors.Interface, board string, threadId int) {
					defer waitGroup.Done()

					threadFiles, fetchFilesErr := vendor.FetchFiles(board, threadId)
					if fetchFilesErr != nil {
						log.Println("Error fetchFiles:", fetchFilesErr)
						return
					}

					filesChannel <- structs.FileChannelMessage{
						LocalBoard: localBoard.Name,
						Files:      threadFiles,
					}
				}(desiredVendor, localBoard.Name, threadId)
			}
		}
	}

	waitGroup.Wait()
	log.Println("QUEUE IS EMPTY")
	var counter = 0

	for _, boards := range response {
		counter += len(boards)
	}

	log.Println("Result files: ", counter)
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
