package request

type InsertRequest struct {
	Table        string              `json:"table"`
	Records      []map[string]string `json:"records"`
	UpdateRecord bool                `json:"update_record"`
}
