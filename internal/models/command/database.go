package command

type Database struct {
	Name   string   `json:"name"`
	Tables []string `json:"tables"`
}
