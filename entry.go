package main

import (
	"daemon/structs"
	"daemon/vendors"
	_2ch "daemon/vendors/2ch"
	_4chan "daemon/vendors/4chan"
	"fmt"
	"log"
	"sync"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic: %s", err)
		}
	}()

	vendors := map[string]vendors.Interface{
		"2ch":   _2ch.Instance(),
		"4chan": _4chan.Instance(),
	}

	//var responseBoards []structs.ResponseBoards
	var waitGroup sync.WaitGroup

	var queueThreads = make(chan []int, 1)
	var queueVideos = make(chan []structs.Video, 1)

	waitGroup.Add(1)
	for _, boardStruct := range getGrabberSchema() {
		for _, sourceBoard := range boardStruct.SourceBoards {
			if desiredVendor, vendorExists := vendors[sourceBoard.Vendor]; vendorExists {
				go func() { queueThreads <- asyncFetchThreads(desiredVendor, sourceBoard.Board) }()
				go func() {
					for threadId := range queueThreads {
						queueVideos = append(queueVideos)
					}
				}()
			}
		}
	}

	waitGroup.Wait()
}

func asyncFetchThreads(vendor vendors.Interface, boardName string) (threadsList []int) {
	threadsList, threadsError := vendor.FetchThreads(boardName)
	if threadsError != nil {
		log.Println("Error (FetchThreads): ", "BoardName: ", boardName)
		return
	}

	return
}

func asyncFetchVideos(vendor vendors.Interface, boardName string, threadId int) (videosList []structs.Video) {
	videosList, videosError := vendor.FetchVideos(boardName, threadId)
	if videosError != nil {
		log.Println("Error (FetchVideos): ", "BoardName: ", boardName, "ThreadId:", threadId)
		return
	}

	return
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
