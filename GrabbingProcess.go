package main

import (
	"daemon/sources/fourChannel"
	"daemon/sources/twoChannel"
	"daemon/sources/types"
	"sync"

	"github.com/davecgh/go-spew/spew"

	"github.com/ztrue/tracerr"
)

var channel = make(chan types.ChannelMessage)
var waitGroup sync.WaitGroup

func catchingFilesChannel(output *types.Output) {
	for {
		message := <-channel
		output.Push(&message)
	}
}

func fetch(vendor types.Interface, thread types.Thread) {
	defer waitGroup.Done()

	files, err := vendor.FetchFiles(thread)
	if err != nil {
		tracerr.Print(tracerr.Wrap(err))
		return
	}

	if len(files) == 0 {
		return
	}

	channel <- types.ChannelMessage{
		VendorName: vendor.VendorName(),
		Thread:     thread,
		Files:      files,
	}
}

func GrabberProcess() types.Output {
	allowedExtensions := types.AllowedExtensions{".webm", ".mp4"}

	grabberSchemas := []types.GrabberSchema{
		{
			twoChannel.Make(allowedExtensions),
			[]types.Board{"b", "a"},
		},
		{
			fourChannel.Make(allowedExtensions),
			[]types.Board{"b"},
		},
	}

	output := types.MakeOutput(grabberSchemas)
	go catchingFilesChannel(&output)

	for _, schema := range grabberSchemas {
		for _, board := range schema.Boards {
			threads, err := schema.Vendor.FetchThreads(board)
			if err != nil {
				tracerr.PrintSourceColor(tracerr.Wrap(err))
				continue
			}

			for _, thread := range threads {
				waitGroup.Add(1)
				go fetch(schema.Vendor, thread)
			}
		}
	}

	waitGroup.Wait()

	return output
}

func main() {
	files := GrabberProcess()

	spew.Dump(files)
}
