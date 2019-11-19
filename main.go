package main

import (
	"daemon/structs"
	"daemon/vendors"
	_2ch "daemon/vendors/2ch"
	_4chan "daemon/vendors/4chan"
	"log"
	"sync"
)

func GetFiles(grabberSchema []structs.Board) map[string][]structs.File {
	var vendorInstances = map[string]vendors.Interface{
		"2ch":   _2ch.Instance(),
		"4chan": _4chan.Instance(),
	}

	var filesChannel = make(chan structs.FileChannelMessage)
	var waitGroup sync.WaitGroup
	var response = make(map[string][]structs.File)

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
	return response
}
