package models

type Database struct {
	Name   string   `json:"name"`
	Tables []string `json:"tables"`
}
