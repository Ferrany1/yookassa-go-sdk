package recipient

// Recipient Получатель платежа.
type Recipient struct {
	// Идентификатор магазина в ЮKassa.
	AccountID string `json:"account_id,omitempty"`
	// Идентификатор субаккаунта.
	// Используется для разделения потоков платежей в рамках одного аккаунта.
	GatewayID string `json:"gateway_id,omitempty"`
}

func NewRecipient() *Recipient {
	return &Recipient{}
}

func (r *Recipient) WithGatewayID(gatewayID string) *Recipient {
	r.GatewayID = gatewayID
	return r
}
