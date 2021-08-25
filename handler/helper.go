package handler

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type IdResponse struct {
	Id uint64
}

func parseID(strID string) (uint64, error) {

	id, err := strconv.ParseUint(strID, 10, 8)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func ToJSON(i interface{}, w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(i)
}

func FromJSON(i interface{}, r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(i)
}

func fromStrIDsToArrUnitIDs(strIDs string) ([]uint64, error) {
	var res []uint64
	sIDs := strings.Split(strings.Trim(strIDs, ","), ",")
	for _, strID := range sIDs {
		id, err := parseID(strID)

		if err != nil {
			return nil, err
		}
		res = append(res, id)

	}
	return res, nil
}
