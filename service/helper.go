package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ueihvn/go-devduo/model"
)

const resultPerPage = 20

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

func parseID(strID string) (uint64, error) {

	id, err := strconv.ParseUint(strID, 10, 8)
	if err != nil {
		return 0, err
	}

	return id, nil
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

func buildProfileFilterSortPageQuery(fsp *model.FilterSortPage) (string, []interface{}) {
	query := getProfile
	args := []interface{}{}

	if fsp.Filter != nil {
		strFields, existField := fsp.Filter["field"]
		strTechs, existTech := fsp.Filter["tech"]

		if existField && existTech {
			query = filterProfileByFieldTech
			fields, _ := fromStrIDsToArrUnitIDs(strFields)
			techs, _ := fromStrIDsToArrUnitIDs(strTechs)
			args = append(args, fields, techs)
		} else if existField && !existTech {
			query = filterProfileByField
			fields, _ := fromStrIDsToArrUnitIDs(strFields)
			args = append(args, fields)
		} else if !existField && existTech {
			query = filterProfileByTech
			techs, _ := fromStrIDsToArrUnitIDs(strTechs)
			args = append(args, techs)
		}
	}

	sort := "order by"
	if fsp.Sort != nil {
		for field, order := range fsp.Sort {
			sort = fmt.Sprint(sort, " ", field, " ", order, ",")
		}
		sort = strings.Trim(sort, ",")
	}

	page := ""
	if fsp.Page != 0 {
		// page = "offset " + strconv.FormatUint((fsp.Page-1)*resultPerPage, 10) + "limit" + strconv.Itoa(resultPerPage)
		page = fmt.Sprint("offset ", ((fsp.Page - 1) * resultPerPage), " limit ", resultPerPage)
	}

	return fmt.Sprint(query, sort, " ", page), args
}
