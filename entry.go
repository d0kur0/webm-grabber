package main

import (
	"daemon/vendors"
	_2ch "daemon/vendors/2ch"
	"fmt"
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
		_ = instance
	}
}
