package main

import (
	"daemon/sources/twoChannel"
	"daemon/sources/types"
	"log"
	"sync"

	"github.com/ztrue/tracerr"
)

var channel = make(chan types.ChannelMessage)
var waitGroup sync.WaitGroup

func catchingFilesChannel(output *types.Output) {
	message := <-channel
	log.Println(message)
	output.Push(&message)
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

func GrabberProcess() {
	allowedExtensions := types.AllowedExtensions{".webm", ".mp4"}

	grabberSchemas := []types.GrabberSchema{
		{
			twoChannel.Make(allowedExtensions),
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

	log.Print("All jobs done")
}

func main() {
	GrabberProcess()
}
