package handler

import (
	"strconv"
)

func parseID(idJSON string) uint64 {

	id, _ := strconv.ParseUint(idJSON, 10, 8)
	return id
}
