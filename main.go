package webmGrabber

import (
	"fmt"
	"log"
	"sync"

	"github.com/d0kur0/webm-grabber/types"
	"github.com/pkg/errors"
)

var channel chan types.ChannelMessage
var done chan bool
var waitGroup sync.WaitGroup

func catchingFilesChannel(output *types.Output) {
	for {
		select {
		case message := <-channel:
			for _, file := range message.Files {
				item := types.OutputItem{
					VendorName:   message.VendorName,
					BoardName:    message.Thread.Board.Name,
					SourceThread: message.SourceThread,
					File:         file,
				}

				*output = append(*output, item)
			}
		case <-done:
			return
		}
	}
}

func fetch(vendor types.VendorInterface, thread types.Thread) {
	defer waitGroup.Done()

	files, err := vendor.FetchFiles(thread)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("FetchFiles error, vendor: %s, thread: %d, board: %s", vendor.VendorName(), thread.ID, thread.Board.Name))
		return
	}

	if len(files) == 0 {
		return
	}

	channel <- types.ChannelMessage{
		VendorName:   vendor.VendorName(),
		SourceThread: vendor.GetThreadUrl(thread),
		Thread:       thread,
		Files:        files,
	}
}

func GrabberProcess(grabberSchemas []types.GrabberSchema) (output types.Output) {
	channel = make(chan types.ChannelMessage)
	done = make(chan bool)

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
	done <- true
	return output
}
