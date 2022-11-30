package fraud_data //nolint:revive // better package readability

// FraudData Информация для проверки операции на мошенничество
type FraudData struct {
	ToppedUpPhone string `json:"topped_up_phone"`
}

func NewFraudData() *FraudData {
	return &FraudData{}
}

func (d *FraudData) WithToppedUpPhone(toppedUpPhone string) *FraudData {
	d.ToppedUpPhone = toppedUpPhone
	return d
}
