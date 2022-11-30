package deal

import "github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/amount"

type Settlement struct {
	// Тип операции. Фиксированное значение: payout — выплата продавцу.
	Type SettlementType `json:"type,omitempty"`
	// Сумма вознаграждения продавца.
	Amount amount.Amount `json:"amount,omitempty"`
}

func NewSettlement() *Settlement {
	return &Settlement{}
}

func (s *Settlement) WithID(settlementType SettlementType) *Settlement {
	s.Type = settlementType
	return s
}

func (s *Settlement) WithAmount(amount amount.Amount) *Settlement {
	s.Amount = amount
	return s
}

type Settlements []*Settlement
