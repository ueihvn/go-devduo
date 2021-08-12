package handler

import (
	"strconv"
)

func parseID(idJSON string) (uint64, error) {

	id, err := strconv.ParseUint(idJSON, 10, 8)
	if err != nil {
		return 0, err
	}

	return id, nil
}
