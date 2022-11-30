package payment_mode //nolint:revive // better package readability

type PaymentMode string

const (
	// FullPrepaymentPaymentMode Полная предоплата
	FullPrepaymentPaymentMode PaymentMode = "full_prepayment"
	// PartialPrepaymentPaymentMode Частичная предоплата
	PartialPrepaymentPaymentMode PaymentMode = "partial_prepayment"
	// AdvancePaymentMode Аванс
	AdvancePaymentMode PaymentMode = "advance"
	// FullPaymentPaymentMode Полный расчет
	FullPaymentPaymentMode PaymentMode = "full_payment"
	// PartialPaymentPaymentMode Частичный расчет и кредит
	PartialPaymentPaymentMode PaymentMode = "partial_payment"
	// CreditPaymentMode Кредит
	CreditPaymentMode PaymentMode = "credit"
	// CreditPaymentPaymentMode Выплата по кредиту
	CreditPaymentPaymentMode PaymentMode = "credit_payment"
)
