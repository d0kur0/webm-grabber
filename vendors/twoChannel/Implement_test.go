package twoChannel

import (
	"daemon/types"
	"testing"
)

var implement = Make(types.AllowedExtensions{".webm", ".mp4"})

func TestFetchThreads(t *testing.T) {
	threads, err := implement.FetchThreads("b")

}
