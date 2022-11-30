package adapter

import "fmt"

type Error struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Parameter   string `json:"parameter"`
}

func (e Error) Error() string {
	const format = "type: %s, id: %s, code: %s, description: %s, parameter: %s"
	return fmt.Sprintf(format, e.Type, e.ID, e.Code, e.Description, e.Parameter)
}
