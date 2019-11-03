package main

import (
	"daemon/vendors"
	_2ch "daemon/vendors/2ch"
	"log"
)

var vendorInstances = map[string]vendors.Interface{
	"2ch": _2ch.Instance(),
	//"4chan": _4chan.Instance(),
}

func main() {
	for _, localBoard := range GrabberSchema {
		for _, sourceBoard := range localBoard.SourceBoards {
			desiredVendor, vendorExists := vendorInstances[sourceBoard.Vendor]
			if !vendorExists {
				log.Println("Not found vendor", sourceBoard.Vendor)
				continue
			}

			threads, fetchThreadsErr := desiredVendor.FetchThreads(sourceBoard.Board)
			if fetchThreadsErr != nil {
				log.Println("Error FetchThreads:", fetchThreadsErr)
				continue
			}

			for _, threadId := range threads {
				// Async fetching Files
				_ = threadId
			}
		}
	}
}
