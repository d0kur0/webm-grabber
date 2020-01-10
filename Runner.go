package daemon

import (
	"daemon/types"
	"daemon/vendors/twoChannel"
)

func Runner() {
	allowedExtensions := types.AllowedExtensions{"webm", "mp4"}

	GrabberSchemas := []GrabberSchema{
		{
			twoChannel.Make(allowedExtensions),
			[]types.Board{"b", "a", "g"},
		},
	}
}
