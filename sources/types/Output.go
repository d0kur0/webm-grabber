package types

import (
	"errors"
)

type outputVendors map[string][]outputBoard

type outputBoard struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Threads     []outputThread
}

type outputThread struct {
	Id    int    `json:"id"`
	Files []File `json:"files"`
}

type Output struct {
	Vendors outputVendors `json:"vendors"`
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

	o.Vendors[vendor][desiredBoardIndex].Threads = append(o.Vendors[vendor][desiredBoardIndex].Threads, outputThread{
		Id:    message.Thread.ID,
		Files: message.Files,
	})

	return nil
}

func MakeOutput(schemas []GrabberSchema) (o Output) {
	o = Output{}
	o.Vendors = make(outputVendors, len(schemas))

	for _, schema := range schemas {
		var boards []outputBoard

		for _, board := range schema.Boards {
			boards = append(boards, outputBoard{
				Name:        board.Name,
				Description: board.Description,
				Threads:     []outputThread{},
			})
		}

		o.Vendors[schema.Vendor.VendorName()] = boards
	}

	return
}
