package models

type Sequence struct {
	Name           string `json:"name"`
	LasValue       int    `json:"last_value"`
	IncrementValue int    `json:"increment_value"`
}
