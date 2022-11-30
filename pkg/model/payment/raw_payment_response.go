package payment

import (
	"time"

	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/amount"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/authorization"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/cancellation"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/deal"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/metadata"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/recipient"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/status"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/transfer"
)

type RawPaymentResponsePart struct {
	// Идентификатор платежа в ЮKassa.
	ID     string        `json:"id"`
	Status status.Status `json:"status"`
	Amount amount.Amount `json:"amount"`
	// Сумма платежа, которую получит магазин, — значение amount
	// за вычетом комиссии ЮKassa.
	// Если вы партнер и для аутентификации запросов используете OAuth-токен,
	// запросите у магазина право на получение информации о комиссиях при платежах.
	IncomeAmount amount.Amount `json:"income_amount"`
	// Описание транзакции (не более 128 символов), которое вы увидите в личном кабинете ЮKassa,
	// а пользователь — при оплате. Например: «Оплата заказа № 72 для user@yoomoney.ru».
	Description string              `json:"description"`
	Recipient   recipient.Recipient `json:"recipient"`
	// Время подтверждения платежа. Указывается по UTC и передается в формате ISO 8601.
	CapturedAt time.Time `json:"captured_at"`
	// Время создания заказа. Указывается по UTC и передается в формате ISO 8601. Пример: 2017-11-03T11:52:31.827Z
	CreatedAt time.Time `json:"created_at"`
	// Время, до которого вы можете бесплатно отменить или подтвердить платеж.
	// В указанное время платеж в статусе waiting_for_capture
	// будет автоматически отменен.
	ExpiresAt time.Time `json:"expires_at"`
	// Способ оплаты, который был использован для этого платежа.
	// Признак тестовой операции.
	Test bool `json:"test"`
	// Сумма, которая вернулась пользователю.
	// Присутствует, если у этого платежа есть успешные возвраты.
	RefundedAmount amount.Amount `json:"refunded_amount"`
	// Признак оплаты заказа.
	Paid bool `json:"paid"`
	// Возможность провести возврат по API.
	Refundable bool `json:"refundable"`
	// Статус доставки данных для чека в онлайн-кассу.
	// Присутствует, если вы используете решение ЮKassa для работы по 54-ФЗ.
	ReceiptRegistration  status.Status                      `json:"receipt_registration"`
	Metadata             metadata.Metadata                  `json:"metadata"`
	CancellationDetails  cancellation.CancellationDetails   `json:"cancellation_details"`
	AuthorizationDetails authorization.AuthorizationDetails `json:"authorization_details"`
	Transfers            transfer.Transfers                 `json:"transfers"`
	Deal                 deal.Deal                          `json:"deal"`
	// Идентификатор покупателя в вашей системе, например электронная почта или номер телефона.
	// Не более 200 символов.
	// Присутствует, если вы хотите запомнить банковскую карту и отобразить ее при повторном платеже в виджете ЮKassa.
	MerchantCustomerId string `json:"merchant_customer_id"`
}

type RawPaymentResponse struct {
	RawPaymentResponsePart
	// Способ оплаты, который был использован для этого платежа.
	PaymentMethod map[string]interface{} `json:"payment_method,omitempty"`
	// Выбранный способ подтверждения платежа.
	// Присутствует, когда платеж ожидает подтверждения от пользователя.
	// Подробнее о сценариях подтверждения: https://yookassa.ru/developers/payment/payment-process#user-confirmation
	Confirmation map[string]interface{} `json:"confirmation,omitempty"`
}
