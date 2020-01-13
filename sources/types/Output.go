package types

type outputItem struct {
	thread Thread
	files  []File
}

type Output struct {
	vendors map[string]struct {
		boards map[string]outputItem
	}
}

func (output *Output) Push(message ChannelMessage) {
	if _, isVendorExists := output.vendors[message.VendorName]; !isVendorExists {
		output.vendors[message.VendorName] = struct{ boards map[string]outputItem }{boards: nil}
	}

	vendor := output.vendors[message.VendorName]
	board := message.Thread.Board.String()

	if _, isBoardExists := vendor.boards[board]; !isBoardExists {
		vendor.boards[board] = outputItem{
			thread: message.Thread,
			files:  nil,
		}
	} else {
		vendor.boards[board].files = nil
	}
}
