package webmGrabber

import (
	"fmt"
	"log"
	"sync"

	"github.com/pkg/errors"

	"github.com/d0kur0/webm-grabber/sources/types"
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
		err = errors.Wrap(err, fmt.Sprintf("FetchFiles error, vendor: %s, thread: %d, board: %s", vendor.VendorName(), thread, thread.Board.String()))
		log.Println(err)
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

func GrabberProcess(grabberSchemas []types.GrabberSchema) types.Output {
	output := types.MakeOutput(grabberSchemas)
	go catchingFilesChannel(&output)

	for _, schema := range grabberSchemas {
		for _, board := range schema.Boards {
			threads, err := schema.Vendor.FetchThreads(board)
			if err != nil {
				err = errors.Wrap(err, fmt.Sprintf("FetchThreads error: vendor %s, board: %s", schema.Vendor.VendorName(), board.String()))
				log.Println(err)
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
