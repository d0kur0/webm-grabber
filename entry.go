package main

import "daemon/vendors/implementation"

func main() {
	testVendor := implementation.TestVendorFactory()
	_ = testVendor
}
