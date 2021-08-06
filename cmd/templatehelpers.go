package main

import (
	"fmt"
	"strings"
	"text/template"

	"EDteam/go-wizard/model"

	"github.com/stoewer/go-strcase"
)

var templateFuncs = template.FuncMap{
	"parseToUpperCamelCase": func(v string) string {
		return strcase.UpperCamelCase(v)
	},
	"parseToUpper": func(v string) string {
		return strings.ToUpper(v)
	},
	"parseToKebabCase": func(v string) string {
		return strcase.KebabCase(v)
	},
	"parseToLowerCamelCase": func(v string) string {
		return strcase.LowerCamelCase(v)
	},
	"increment": func(v int) int {
		return v + 1
	},
	"decrement": func(v int) int {
		return v - 1
	},
	"parseToSqlType": func(v string) string {
		switch v {
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
			return "INTEGER"
		case "float64", "float32":
			return "NUMERIC(SIZE)"
		case "string":
			return "VARCHAR(SIZE)"
		case "bool":
			return "BOOLEAN"
		case "time.Time":
			return "TIMESTAMP"
		}
		return "CHANGE-THIS-TYPE"
	},
	"handleNull": func(f model.Field) string {
		field := strcase.UpperCamelCase(f.Name)

		if strings.TrimSpace(strings.ToUpper(f.NotNull)) == "NOT NULL" {
			return fmt.Sprintf("m.%s", field)
		}

		switch f.Type {
		case "string":
			return fmt.Sprintf("sqlutil.StringToNull(m.%s)", field)
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
			return fmt.Sprintf("sqlutil.IntToNull(int64(m.%s))", field)
		case "float64", "float32":
			return fmt.Sprintf("sqlutil.FloatToNull(float64(m.%s))", field)
		case "time.Time":
			return fmt.Sprintf("sqlutil.TimeToNull(m.%s)", field)
		default:
			return fmt.Sprintf("invalid data type: %s", f.Type)
		}
	},
	"handleNullOnScan": func(f model.Field) string {
		if strings.TrimSpace(strings.ToUpper(f.NotNull)) == "NOT NULL" {
			return ""
		}

		switch f.Type {
		case "string":
			return fmt.Sprintf("%s := sql.NullString{}", f.Name)
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
			return fmt.Sprintf("%s := sql.NullInt64{}", f.Name)
		case "float64", "float32":
			return fmt.Sprintf("%s := sql.NullFloat64{}", f.Name)
		case "time.Time":
			return fmt.Sprintf("%s := pq.NullTime{}", f.Name)
		case "bool":
			return fmt.Sprintf("%s := sql.NullBool{}", f.Name)
		default:
			return fmt.Sprintf("invalid data type: %s", f.Type)
		}
	},
	"setNullFieldsOnScan": func(f model.Field) string {
		field := strcase.UpperCamelCase(f.Name)
		if f.NotNull == "NOT NULL" {
			return ""
		}

		switch f.Type {
		case "string":
			return fmt.Sprintf("m.%s = %s.String", field, f.Name)
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
			return fmt.Sprintf("m.%s = %s(%s.Int64)", field, f.Type, f.Name)
		case "float64", "float32":
			return fmt.Sprintf("m.%s = %s(%s.Float64)", field, f.Type, f.Name)
		case "time.Time":
			return fmt.Sprintf("m.%s = %s.Time", field, f.Name)
		case "bool":
			return fmt.Sprintf("m.%s = %s.Bool", field, f.Name)
		default:
			return fmt.Sprintf("invalid data type: %s", f.Type)
		}
	},
}
