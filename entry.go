package main

import "daemon/vendors/implementation"

func main() {
	testVendor := implementation.Instance2ch()
	_ = testVendor
}
