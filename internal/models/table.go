package models

type Table struct {
	TableName string  `json:"table_name"`
	Fields    []Field `json:"fields"`
}
