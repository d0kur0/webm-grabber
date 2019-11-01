package main

import (
	"sync"

	"github.com/davecgh/go-spew/spew"
)

type Video int

func main() {
	var videos []Video
	var waitGroup sync.WaitGroup
	var queue = make(chan []Video, 1)

	waitGroup.Add(1)
	go func() { queue <- asyncThreads() }()

	go func() {
		for video := range queue {
			videos = append(videos, video...)
			waitGroup.Done()
		}
	}()

	waitGroup.Wait()
	spew.Dump(videos)
}

func asyncThreads() (videos []Video) {
	for i := 0; i < 100; i++ {
		videos = append(videos, Video(i))
	}

	return
}

func asyncVideos() {

}
