package deal

// Deal Данные о сделке, в составе которой проходит платеж.
// Присутствует, если вы проводите Безопасную сделку.
type Deal struct {
	// Идентификатор сделки.
	ID string `json:"id,omitempty"`
	// Данные о распределении денег.
	Settlements Settlements `json:"settlements,omitempty"`
}

func NewDeal() *Deal {
	return &Deal{}
}

func (d *Deal) WithID(id string) *Deal {
	d.ID = id
	return d
}

func (d *Deal) WithSettlements(settlements ...*Settlement) *Deal {
	d.Settlements = settlements
	return d
}
