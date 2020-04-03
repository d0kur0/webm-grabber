package webmGrabber

import (
	"fmt"
	"github.com/d0kur0/webm-grabber/sources/types"
	"github.com/pkg/errors"
	"log"
	"sync"
)

var channel chan types.ChannelMessage
var waitGroup sync.WaitGroup

func catchingFilesChannel(output *types.Output) {
	for {
		message := <-channel
		if err := output.Push(&message); err != nil {
			log.Println(errors.Wrap(err, "Push in output error"))
		}
	}
}

func fetch(vendor types.Interface, thread types.Thread) {
	defer waitGroup.Done()

	files, err := vendor.FetchFiles(thread)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("FetchFiles error, vendor: %s, thread: %d, board: %s", vendor.VendorName(), thread.ID, thread.Board.Name))
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

	channel = make(chan types.ChannelMessage)
	output := types.MakeOutput(grabberSchemas)
	go catchingFilesChannel(&output)

	for _, schema := range grabberSchemas {
		for _, board := range schema.Boards {
			threads, err := schema.Vendor.FetchThreads(board)
			if err != nil {
				err = errors.Wrap(err, fmt.Sprintf("FetchThreads error: vendor %s, board: %s", schema.Vendor.VendorName(), board.Name))
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