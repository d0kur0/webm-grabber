package sources

import (
	"fmt"
)

type Thread struct {
	ID    int
	Board Board
}

func (thread *Thread) StringId() string {
	return fmt.Sprint(thread.ID)
}
