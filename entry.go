package main

import (
	"daemon/vendors"
	_2ch "daemon/vendors/2ch"
	"fmt"
	"log"
)

func main() {
	defer func() {
		if error := recover(); error != nil {
			fmt.Printf("Panic: %s", error)
		}
	}()

	instances := []vendors.Interface{
		_2ch.Instance(),
	}

	for _, instance := range instances {
		err, threads := instance.FetchThreads("b")
		if err != nil {
			log.Fatal(err)
			continue
		}

		_ = threads
	}
}
