package payment

import (
	"strconv"

	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/amount"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/confirmation"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment_method"
)

type PaymentRequest[
	PaymentMethod payment_method.GenericPaymentMethod,
	Confirmation confirmation.GenericConfirmation,
] struct {
	RawPaymentPart
	PaymentMethodData *PaymentMethod `json:"payment_method_data"`
	// Данные, необходимые для инициирования выбранного сценария подтверждения платежа пользователем.
	Confirmation *Confirmation `json:"confirmation"`
}

func NewPaymentRequest[
	PaymentMethod payment_method.GenericPaymentMethod,
	Confirmation confirmation.GenericConfirmation,
]() *PaymentRequest[PaymentMethod, Confirmation] {
	return &PaymentRequest[PaymentMethod, Confirmation]{}
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithAmount(
	value float64,
	currency amount.Currency,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.Amount = amount.Amount{
		Value:    strconv.FormatFloat(value, 'f', 2, 64),
		Currency: currency,
	}
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithDescription(
	description string,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.Description = description
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithRecipient(
	recipient Recipient,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.Recipient = &recipient
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithDeal(
	deal Deal,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.Deal = &deal
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithMerchantCustomerID(
	merchantCustomerID string,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.MerchantCustomerID = merchantCustomerID
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithTransfers(
	transfers Transfers,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.Transfers = &transfers
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithReceipt(
	receipt Receipt,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.Receipt = &receipt
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithPaymentToken(
	paymentToken string,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.PaymentToken = paymentToken
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithPaymentMethodID(
	paymentMethodID string,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.PaymentMethodID = paymentMethodID
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithPaymentMethodData(
	paymentMethodData PaymentMethod,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.PaymentMethodData = &paymentMethodData
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithConfirmation(
	confirmation Confirmation,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.Confirmation = &confirmation
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithSavePaymentMethod() *PaymentRequest[PaymentMethod, Confirmation] {
	p.SavePaymentMethod = true
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithCapture() *PaymentRequest[PaymentMethod, Confirmation] {
	p.Capture = true
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithClientIP(
	clientIP string,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.ClientIP = clientIP
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithMetadata(
	metadata Metadata,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.Metadata = &metadata
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithAirlineTicketData(
	airlineTicketData AirlineTicketData,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.AirlineTicketData = &airlineTicketData
	return p
}
