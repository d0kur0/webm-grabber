package main

import (
	"daemon/structs"
	"daemon/vendors"
	_2ch "daemon/vendors/2ch"
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic: %s", err)
		}
	}()

	instances := map[string]vendors.Interface{
		"2ch": _2ch.Instance(),
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
					continue
				}

				for _, thread := range threads {
					videos, err := instance.FetchVideos(sourceBoard.Board, thread)
					if err != nil {
						continue
					}

					responseBoard.Videos = append(responseBoard.Videos, videos...)
				}
			}
		}

		responseBoards = append(responseBoards, responseBoard)
	}

	spew.Dump(responseBoards)
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
				{"2ch", "e"},
				{"4chan", "s"},
				{"4chan", "c"},
			},
		},
		{
			Name:        "c",
			Description: "...",
			SourceBoards: []structs.SourceBoard{
				{"2ch", "h"},
				{"2ch", "fur"},
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
