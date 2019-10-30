package main

import (
	"daemon/structs"
	"daemon/vendors"
	_4chan "daemon/vendors/4chan"
	"fmt"
	"log"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic: %s", err)
		}
	}()

	instances := map[string]vendors.Interface{
		//"2ch":   _2ch.Instance(),
		"4chan": _4chan.Instance(),
	}

	var responseBoards []structs.ResponseBoards
	for _, board := range getGrabberSchema() {
		var responseBoard = structs.ResponseBoards{
			BoardName:   board.Name,
			Description: board.Description,
			Videos:      []structs.Video{},
		}

		for _, sourceBoard := range board.SourceBoards {
			if instance, exists := instances[sourceBoard.Vendor]; exists {
				threads, err := instance.FetchThreads(sourceBoard.Board)
				if err != nil {
					log.Println("FetchThreads return error:", err, "BoardName:", sourceBoard.Board)
					continue
				}

				for _, thread := range threads {
					videos, err := instance.FetchVideos(sourceBoard.Board, thread)
					if err != nil {
						log.Println("FetchVideos return error:", err, "BoardName:", sourceBoard.Board, "ThreadId:", thread)
						continue
					}

					responseBoard.Videos = append(responseBoard.Videos, videos...)
				}
			} else {
				log.Println("A nonexistent vendor is called: ", sourceBoard.Vendor)
			}
		}

		responseBoards = append(responseBoards, responseBoard)
	}

	//spew.Dump(responseBoards)
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
		{
			Name:        "c",
			Description: "...",
			SourceBoards: []structs.SourceBoard{
				{"4chan", "h"},
				{"4chan", "u"},
				{"4chan", "d"},
				{"4chan", "e"},
				{"4chan", "aco"},
			},
		},
	}

	return
}
