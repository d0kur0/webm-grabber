package types

import (
	"errors"
)

type OutputVendors map[string][]OutputBoard

type OutputBoard struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Threads     []OutputThread `json:"threads"`
}

type OutputThread struct {
	Id    int    `json:"id"`
	Files []File `json:"files"`
}

type Output struct {
	Vendors OutputVendors `json:"vendors"`
}

func (o *Output) Push(message *ChannelMessage) error {
	vendor := message.VendorName
	board := message.Thread.Board.Name

	var desiredBoardIndex = -1
	for vendorBoardIndex, vendorBoard := range o.Vendors[vendor] {
		if vendorBoard.Name == board {
			desiredBoardIndex = vendorBoardIndex
			break
		}
	}

	if desiredBoardIndex == -1 {
		return errors.New("board not found")
	}

	o.Vendors[vendor][desiredBoardIndex].Threads = append(o.Vendors[vendor][desiredBoardIndex].Threads, OutputThread{
		Id:    message.Thread.ID,
		Files: message.Files,
	})

	return nil
}

func MakeOutput(schemas []GrabberSchema) (o Output) {
	o = Output{}
	o.Vendors = make(OutputVendors, len(schemas))

	for _, schema := range schemas {
		var boards []OutputBoard

		for _, board := range schema.Boards {
			boards = append(boards, OutputBoard{
				Name:        board.Name,
				Description: board.Description,
				Threads:     []OutputThread{},
			})
		}

		o.Vendors[schema.Vendor.VendorName()] = boards
	}

	return
}
