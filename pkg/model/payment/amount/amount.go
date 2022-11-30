package amount

import "strconv"

// Amount Сумма платежа.
// Иногда партнеры ЮKassa берут с пользователя дополнительную комиссию, которая не входит в эту сумму.
type Amount struct {
	// Сумма в выбранной валюте. Выражается в виде строки и пишется через точку, например 10.00.
	// Количество знаков после точки зависит от выбранной валюты.
	Value string `json:"value"`
	// Код валюты в формате ISO-4217. Должен соответствовать валюте субаккаунта (recipient.gateway_id),
	// если вы разделяете потоки платежей, и валюте аккаунта (shopId в личном кабинете), если не разделяете.
	Currency Currency `json:"currency"`
}

func NewAmount() *Amount {
	return &Amount{}
}

func (a *Amount) WithFloatValue(value float64) *Amount {
	a.Value = strconv.FormatFloat(value, 'f', 2, 64)
	return a
}

func (a *Amount) WithStringValue(value string) *Amount {
	a.Value = value
	return a
}

func (a *Amount) WithCurrency(currency Currency) *Amount {
	a.Currency = currency
	return a
}
