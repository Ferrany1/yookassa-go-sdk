package payment

type Error struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Parameter   string `json:"parameter"`
}
