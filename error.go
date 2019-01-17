package datamall

import (
	"fmt"
)

type Error struct {
	StatusCode int
}

func (e Error) Error() string {
	return fmt.Sprintf("datamall: received %d status code from server", e.StatusCode)
}
