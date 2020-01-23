package types

type outputVendors map[string]outputBoards

type outputBoards map[string][]outputThread

type outputThread struct {
	Id    int
	files []File
}

type Output struct {
	Vendors outputVendors
}

func (o *Output) Push(message *ChannelMessage) {
	vendor := message.VendorName
	board := message.Thread.Board.String()

	o.Vendors[vendor][board] = append(o.Vendors[vendor][board], outputThread{
		Id:    message.Thread.ID,
		files: message.Files,
	})
}

func MakeOutput(schemas []GrabberSchema) (o Output) {
	o = Output{}
	o.Vendors = make(outputVendors, len(schemas))

	for _, schema := range schemas {
		var boards = make(outputBoards, len(schema.Boards))

		for _, board := range schema.Boards {
			boards[board.String()] = []outputThread{}
		}

		o.Vendors[schema.Vendor.VendorName()] = boards
	}

	return
}
