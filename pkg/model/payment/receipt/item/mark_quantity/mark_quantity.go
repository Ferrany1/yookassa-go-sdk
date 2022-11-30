package mark_quantity //nolint:revive // better package readability

type MarkQuantity struct {
	Numerator   int `json:"numerator"`
	Denominator int `json:"denominator"`
}

func NewMarkQuantity() *MarkQuantity {
	return &MarkQuantity{}
}

func (q *MarkQuantity) WithNumerator(numerator int) *MarkQuantity {
	q.Numerator = numerator
	return q
}

func (q *MarkQuantity) WithDenominator(denominator int) *MarkQuantity {
	q.Denominator = denominator
	return q
}
