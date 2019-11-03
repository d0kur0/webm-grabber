package main

import "daemon/structs"

var GrabberSchema = []structs.Board{
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
