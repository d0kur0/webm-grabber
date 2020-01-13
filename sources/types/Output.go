package types

type Output struct {
	items map[string]struct {
		Thread Thread
		Files []File
	}
}

func (output *Output) Push(message *ChannelMessage) {
	output.items[message.VendorName] = struct {
		Thread Thread
		Files []File
	}{Thread: message.Thread, Files: message.Files}
}

func (output *Output)