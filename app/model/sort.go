package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

const (
	descending = "desc"
	ascending  = "asc"
)

// Sort format: field_1:mode_1;field_2:mode_2...
type Sort struct {
	SortList  string      `json:"-"`
	SortValue []SortValue `json:"sort_value"`
}

type SortValue struct {
	Field string
	Mode  string
}

func (s *Sort) SetDefault(fieldConversions map[string]string) {
	// get list sort
	var list []string

	if s.SortList == "" {
		list = append(list, "_id:desc")
	} else {
		// split by ';'
		list = strings.Split(s.SortList, ";")
	}

	// convert
	var sortValues []SortValue
	for _, val := range list {
		v := strings.Split(val, ":")

		// length of array after split should be two
		if len(v) == 2 {
			// field
			field := v[0]
			if fieldConversions != nil && fieldConversions[field] != "" {
				field = fieldConversions[field]
			}

			// mode
			mode := v[1]

			if strings.ToLower(mode) == descending {
				mode = descending
			} else {
				mode = ascending
			}

			// add to array
			sortValues = append(sortValues, SortValue{
				Field: strings.ToLower(field),
				Mode:  strings.ToLower(mode),
			})
		}
	}

	s.SortValue = sortValues
}

func (s *Sort) Option() *options.FindOptions {
	// loop for add sort, format: field:mode (ex: _id:desc)
	sortOption := bson.D{}
	for _, val := range s.SortValue {
		if val.Mode == descending {
			sortOption = append(sortOption, bson.E{Key: val.Field, Value: -1})
		} else {
			sortOption = append(sortOption, bson.E{Key: val.Field, Value: 1})
		}
	}

	return &options.FindOptions{
		Sort: sortOption,
	}
}
