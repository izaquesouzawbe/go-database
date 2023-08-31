package request

type InsertRequest struct {
	Table       string        `json:"table"`
	PrimaryKeys []string      `json:"primary_keys"`
	Records     []interface{} `json:"records"`
}
