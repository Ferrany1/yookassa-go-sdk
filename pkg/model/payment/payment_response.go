package payment

import (
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/confirmation"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/payment_method"
)

type PaymentResponse[
	PaymentMethod payment_method.GenericPaymentMethod,
	Confirmation confirmation.GenericConfirmation,
] struct {
	RawPaymentResponsePart
	PaymentMethod PaymentMethod `json:"payment_method"`
	// Выбранный способ подтверждения платежа.
	// Присутствует, когда платеж ожидает подтверждения от пользователя.
	// Подробнее о сценариях подтверждения: https://yookassa.ru/developers/payment/payment-process#user-confirmation
	Confirmation Confirmation `json:"confirmation"`
}
