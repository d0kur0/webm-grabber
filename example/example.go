package main

import (
	"fmt"
	"github.com/d0kur0/webm-grabber/vendors/fourChannel"
	"github.com/d0kur0/webm-grabber/vendors/twoChannel"

	webmGrabber "github.com/d0kur0/webm-grabber"
	"github.com/d0kur0/webm-grabber/types"
)

func main() {
	allowedExtensions := types.AllowedExtensions{".mp4", ".webm"}

	grabberSchema := []types.GrabberSchema{
		{
			Vendor: twoChannel.Make(allowedExtensions),
			Boards: []types.Board{
				{"b", "Бред"},
				{"media", "Media"},
			},
		},
		{
			Vendor: fourChannel.Make(allowedExtensions),
			Boards: []types.Board{
				{"b", "Бред"},
			},
		},
	}

	result := webmGrabber.GrabberProcess(grabberSchema)

	for _, item := range result {
		fmt.Printf("Vendor: %s; Board: %s; Thread URL: %s \n", item.VendorName, item.BoardName, item.SourceThread)
		fmt.Printf("File name: %s; \n", item.File.Name)
		fmt.Printf("File preview: %s; \n", item.File.Preview)
		fmt.Printf("File path: %s; \n", item.File.Path)
		fmt.Println("---------------")
	}

	fmt.Println("Count of items: ", len(result))
}
