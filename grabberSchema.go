package main

import "daemon/structs"

var GrabberSchema = []structs.Board{
	{
		Name:        "b",
		Description: "Bred",
		SourceBoards: []structs.SourceBoard{
			{"2ch", "b"},
			{"4chan", "b"},
		},
	},
	{
		Name:        "a",
		Description: "Anime",
		SourceBoards: []structs.SourceBoard{
			{"2ch", "a"},
			{"4chan", "a"},
			{"4chan", "c"},
		},
	},
	{
		Name:        "s",
		Description: "Sexuality",
		SourceBoards: []structs.SourceBoard{
			{"4chan", "s"},
			{"4chan", "c"},
		},
	},
}
