package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/url"
	"strconv"
	"strings"

	"github.com/ueihvn/go-devduo/model"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type IdResponse struct {
	Id uint64
}

type PageData struct {
	Content       []Mentor `json:"content,omitempty"`
	Page          uint64   `json:"page,omitempty"`
	ContentLength uint64   `json:"content_length"`
}

var parseJsonError = "err deserialize data. Check request"
var ParseIDError = "err parse Id. Check request"
var notFoundError = "err records not found. Check request"
var serverInternalError = "err server. Please try again later"

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

func extractDataFromURL(val url.Values) (*model.FilterSortPage, error) {
	var fsp model.FilterSortPage
	strTechs := val.Get("tech")
	if strTechs != "" {
		fsp.Filter = map[string]string{
			"tech": strTechs,
		}
	}

	strFields := val.Get("field")
	if strFields != "" {
		// add filterFields
		if fsp.Filter == nil {
			fsp.Filter = map[string]string{
				"field": strFields,
			}
		} else {
			fsp.Filter["field"] = strFields
		}

	}

	strSorts := val.Get("sort")
	if strSorts != "" {
		sorts, err := checksSort(strSorts)
		if err != nil {
			return nil, err
		}
		fsp.Sort = sorts

	}

	strPage := val.Get("page")
	if strPage != "" {
		iPage, err := parseID(strPage)
		if err != nil {
			return nil, err
		}
		fsp.Page = iPage

	} else {
		fsp.Page = 1
	}

	return &fsp, nil
}

func checksSort(strSorts string) (map[string]string, error) {
	res := map[string]string{}
	if strSorts == "" {
		return nil, errors.New("sort nil")
	}
	sorts := strings.Trim(strSorts, ",")

	arrSorts := strings.Split(sorts, ",")

	for _, sort := range arrSorts {
		sSort := strings.Split(sort, ".")
		if len(sSort) != 2 {
			return nil, errors.New("wrong formart, sort must be field.order")
		}

		if sSort[0] != "full_name" {
			return nil, errors.New("wrong sort field")
		}
		if sSort[1] != "desc" && sSort[1] != "asc" {
			return nil, errors.New("sort order must be asc or desc")
		}

		if _, ok := res[sSort[0]]; ok {
			return nil, errors.New("duplicate sort field")
		}

		res[sSort[0]] = sSort[1]
	}
	return res, nil
}
