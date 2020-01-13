package daemon

import (
	"daemon/types"
	"daemon/vendors"
	"daemon/vendors/twoChannel"
	"log"
	"sync"

	"github.com/ztrue/tracerr"
)

var filesChannel = make(chan []types.File)
var waitGroup sync.WaitGroup

func catchingFilesChannel() {
	file := <-filesChannel

	log.Println("New file in channel: ", file)
}

func fetch(vendor vendors.Interface, thread types.Thread) {
	defer waitGroup.Done()

	files, err := vendor.FetchFiles(thread)
	if err != nil {
		tracerr.PrintSourceColor(tracerr.Wrap(err))
	}

	filesChannel <- files
}

func GrabberProcess() {
	allowedExtensions := types.AllowedExtensions{"webm", "mp4"}

	grabberSchemas := []types.GrabberSchema{
		{
			twoChannel.Make(allowedExtensions),
			[]types.Board{"b", "a", "g"},
		},
	}

	go catchingFilesChannel()

	for _, schema := range grabberSchemas {
		for _, board := range schema.Boards {
			threads, err := schema.Vendor.FetchThreads(board)
			if err != nil {
				tracerr.PrintSourceColor(tracerr.Wrap(err))
				continue
			}

			for _, thread := range threads {
				go fetch(schema.Vendor, thread)
			}
		}
	}
}
