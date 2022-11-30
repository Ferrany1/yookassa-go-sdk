package payment

import (
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/airlane"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/amount"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/confirmation"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/deal"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/fraud_data"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/metadata"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/payment_method"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/receipt"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/recipient"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/transfer"
)

type PaymentRequest[
	PaymentMethod payment_method.GenericPaymentMethod,
	Confirmation confirmation.GenericConfirmation,
] struct {
	RawPaymentRequestPart
	// Данные для оплаты конкретным способом.
	// Вы можете не передавать этот объект в запросе.
	// В этом случае пользователь будет выбирать способ оплаты на стороне ЮKassa.
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
	amount *amount.Amount,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.Amount = amount
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithDescription(
	description string,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.Description = description
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithRecipient(
	recipient *recipient.Recipient,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.Recipient = recipient
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithDeal(
	deal *deal.Deal,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.Deal = deal
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithMerchantCustomerID(
	merchantCustomerID string,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.MerchantCustomerID = merchantCustomerID
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithTransfers(
	transfers *transfer.Transfers,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.Transfers = transfers
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithReceipt(
	receipt *receipt.Receipt,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.Receipt = receipt
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
	metadata *metadata.Metadata,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.Metadata = metadata
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithFraudData(
	fraudData *fraud_data.FraudData,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.FraudData = fraudData
	return p
}

func (p *PaymentRequest[PaymentMethod, Confirmation]) WithAirlineTicketData(
	airlineTicket *airlane.AirlineTicket,
) *PaymentRequest[PaymentMethod, Confirmation] {
	p.Airline = airlineTicket
	return p
}
