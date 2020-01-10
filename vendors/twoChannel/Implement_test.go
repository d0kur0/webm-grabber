package twoChannel

import (
	"daemon/types"
	"testing"
)

var testImplement = Make(types.AllowedExtensions{".webm", ".mp4"})

func TestFetchThreads(t *testing.T) {
	threads, err := testImplement.FetchThreads("b")
	if err != nil {
		t.Error("Fetch threads error", err)
	}

	if len(threads) == 0 {
		t.Error("FetchingThreads retun a 0 rows")
	}
}
