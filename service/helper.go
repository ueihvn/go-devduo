package service

import (
	"strconv"
	"strings"

	"github.com/ueihvn/go-devduo/model"
)

func fromFieldsToStringId(fields []model.Field) string {
	var arrID []string
	for _, field := range fields {
		strID := strconv.Itoa(int(field.ID))
		arrID = append(arrID, strID)
	}

	return strings.Join(arrID, ",")
}

func fromTechnologiesToStringId(technologies []model.Technology) string {
	var arrID []string
	for _, technology := range technologies {
		strID := strconv.Itoa(int(technology.ID))
		arrID = append(arrID, strID)
	}

	return strings.Join(arrID, ",")
}

func fromFieldsToArrFieldId(fields []model.Field) []uint64 {
	var arrID []uint64
	for _, field := range fields {
		arrID = append(arrID, field.ID)
	}

	return arrID
}

func fromTechnologiesToArrTechnonogyId(techs []model.Technology) []uint64 {
	var arrID []uint64
	for _, tech := range techs {
		arrID = append(arrID, tech.ID)
	}

	return arrID
}
