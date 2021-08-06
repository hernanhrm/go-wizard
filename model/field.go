package model

import (
	"errors"
	"strings"
)

var ErrInvalidField = errors.New("invalid len of field")

const fieldSeparator = ":"

const (
	nameFieldIndex = 0
	typeFieldIndex = 1
)

type Field struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	NotNull string `json:"not_null"`
}

type Fields []Field

func NewFieldsFromSliceString(fieldsStr []string) (Fields, error) {
	fields := Fields{}

	for _, fieldStr := range fieldsStr {
		splitField := strings.Split(fieldStr, fieldSeparator)
		if !isValidLen(splitField) {
			return fields, ErrInvalidField
		}

		field := Field{
			Name: splitField[nameFieldIndex],
			Type: splitField[typeFieldIndex],
		}

		if isNameTypeAndNotNull(splitField) {
			field.NotNull = "NOT NULL"
		}

		fields = append(fields, field)
	}

	return fields, nil
}

func isValidLen(splitField []string) bool {
	fieldsLen := len(splitField)
	return fieldsLen >= 2 && fieldsLen <= 3
}

func isNameTypeAndNotNull(splitField []string) bool {
	return len(splitField) == 2
}
