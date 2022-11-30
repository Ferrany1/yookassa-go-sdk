package payment

import (
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/airlane"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/amount"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/deal"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/fraud_data"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/metadata"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/receipt"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/recipient"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/transfer"
)

type RawPaymentRequestPart struct {
	Amount *amount.Amount `json:"amount,omitempty"`
	// Описание транзакции (не более 128 символов), которое вы увидите в личном кабинете ЮKassa,
	// а пользователь — при оплате. Например: «Оплата заказа № 72 для user@yoomoney.ru».
	Description string               `json:"description,omitempty"`
	Receipt     *receipt.Receipt     `json:"receipt,omitempty"`
	Recipient   *recipient.Recipient `json:"recipient,omitempty"`
	// Одноразовый токен для проведения оплаты
	PaymentToken string `json:"payment_token,omitempty"`
	// Одноразовый токен для проведения оплаты, сформированный с помощью веб или мобильного SDK.
	// Идентификатор сохраненного способа оплаты.
	PaymentMethodID string `json:"payment_method_id,omitempty"`
	// Сохранение платежных данных (с их помощью можно проводить повторные безакцептные списания.
	// Значение true инициирует создание многоразового payment_method.
	SavePaymentMethod bool `json:"save_payment_method,omitempty"`
	// Автоматический прием поступившего платежа.
	Capture bool `json:"capture,omitempty"`
	// IPv4 или IPv6-адрес пользователя. Если не указан, используется IP-адрес TCP-подключения.
	ClientIP  string                 `json:"client_ip,omitempty"`
	Metadata  *metadata.Metadata     `json:"metadata,omitempty"`
	Airline   *airlane.AirlineTicket `json:"airline,omitempty"`
	Transfers *transfer.Transfers    `json:"transfers,omitempty"`
	FraudData *fraud_data.FraudData  `json:"fraud_data"`
	Deal      *deal.Deal             `json:"deal,omitempty"`
	// Идентификатор покупателя в вашей системе, например электронная почта или номер телефона.
	// Не более 200 символов.
	// Присутствует, если вы хотите запомнить банковскую карту и отобразить ее при повторном платеже в виджете ЮKassa.
	MerchantCustomerID string `json:"merchant_customer_id,omitempty"`
}

type RawPaymentRequest struct {
	RawPaymentRequestPart
	// Данные для оплаты конкретным способом.
	// Вы можете не передавать этот объект в запросе.
	// В этом случае пользователь будет выбирать способ оплаты на стороне ЮKassa.
	PaymentMethodData *map[string]interface{} `json:"payment_method_data,omitempty"`
	// Данные, необходимые для инициирования выбранного сценария подтверждения платежа пользователем.
	Confirmation *map[string]interface{} `json:"confirmation,omitempty"`
}
