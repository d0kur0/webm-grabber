package main

import (
	"daemon/structs"

	"github.com/davecgh/go-spew/spew"
)

var grabberSchema = []structs.Board{
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

func main() {
	files := GetFiles(grabberSchema)
	spew.Dump(files)
}
