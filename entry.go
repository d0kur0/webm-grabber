package main

import (
	"daemon/vendors/implementation"
	"fmt"
)

func main() {
	defer func() {
		if error := recover(); error != nil {
			fmt.Printf("Panic: %s", error)
		}
	}()

	testVendor := implementation.Instance2ch()
	threads, err := testVendor.FetchThreads("b")
	if err != nil {
		panic(err)
	}

	_ = threads
}
