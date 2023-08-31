package command

type Field struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	NotNull int    `json:"not_null"`
}
