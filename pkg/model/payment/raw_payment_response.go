package payment

import (
	"time"

	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/amount"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/status"
)

type RawPaymentResponse struct {
	// Идентификатор платежа в ЮKassa.
	ID string `json:"id,omitempty"`
	// В случае, если сервер пришлёт ошибку, это может быть записано в Type.
	Type string `json:"type,omitempty"`
	// Сумма платежа.
	// Иногда партнеры ЮKassa берут с пользователя дополнительную комиссию, которая не входит в эту сумму.
	Amount amount.Amount `json:"amount,omitempty"`
	// Описание транзакции (не более 128 символов), которое вы увидите в личном кабинете ЮKassa,
	// а пользователь — при оплате. Например: «Оплата заказа № 72 для user@yoomoney.ru».
	Description string `json:"description,omitempty"`
	// Получатель платежа.
	Recipient Recipient `json:"recipient,omitempty"`
	// Данные о сделке, в составе которой проходит платеж.
	// Присутствует, если вы проводите Безопасную сделку.
	Deal Deal `json:"deal,omitempty"`
	// Идентификатор покупателя в вашей системе, например электронная почта или номер телефона.
	// Не более 200 символов.
	// Присутствует, если вы хотите запомнить банковскую карту и отобразить ее при повторном платеже в виджете ЮKassa.
	MerchantCustomerId string `json:"merchant_customer_id,omitempty"`
	// Данные о распределении денег — сколько и в какой магазин нужно перевести.
	// Присутствует, если вы используете Сплитование платежей.
	Transfers Transfers `json:"transfers,omitempty"`
	// Статус платежа.
	// Подробнее про жизненный цикл платежа: https://yookassa.ru/developers/payment/payment-process#lifecycle.
	Status status.Status `json:"status,omitempty"`
	// Сумма платежа, которую получит магазин, — значение amount
	// за вычетом комиссии ЮKassa.
	// Если вы партнер и для аутентификации запросов используете OAuth-токен,
	// запросите у магазина право на получение информации о комиссиях при платежах.
	IncomeAmount amount.Amount `json:"income_amount,omitempty"`
	// Способ оплаты, который был использован для этого платежа.
	PaymentMethod map[string]interface{} `json:"payment_method,omitempty"`
	// Время подтверждения платежа. Указывается по UTC и передается в формате ISO 8601.
	CapturedAt time.Time `json:"captured_at,omitempty"`
	// Время создания заказа. Указывается по UTC и передается в формате ISO 8601. Пример: 2017-11-03T11:52:31.827Z
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Время, до которого вы можете бесплатно отменить или подтвердить платеж.
	// В указанное время платеж в статусе waiting_for_capture
	// будет автоматически отменен.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// Выбранный способ подтверждения платежа.
	// Присутствует, когда платеж ожидает подтверждения от пользователя.
	// Подробнее о сценариях подтверждения: https://yookassa.ru/developers/payment/payment-process#user-confirmation
	Confirmation map[string]interface{} `json:"confirmation,omitempty"`
	// Признак тестовой операции.
	Test bool `json:"test,omitempty"`
	// Сумма, которая вернулась пользователю.
	// Присутствует, если у этого платежа есть успешные возвраты.
	RefundedAmount amount.Amount `json:"refunded_amount,omitempty"`
	// Признак оплаты заказа.
	Paid bool `json:"paid,omitempty"`
	// Возможность провести возврат по API.
	Refundable bool `json:"refundable,omitempty"`
	// Статус доставки данных для чека в онлайн-кассу.
	// Присутствует, если вы используете решение ЮKassa для работы по 54-ФЗ.
	ReceiptRegistration status.Status `json:"receipt_registration,omitempty"`
	// Любые дополнительные данные, которые нужны вам для работы (например, номер заказа).
	// Передаются в виде набора пар «ключ-значение» и возвращаются в ответе от ЮKassa.
	// Ограничения: максимум 16 ключей, имя ключа не больше 32 символов,
	// значение ключа не больше 512 символов, тип данных — строка в формате UTF-8.
	Metadata Metadata `json:"metadata,omitempty"`
	// Комментарий к статусу canceled: кто отменил платеж и по какой причине.
	// Подробнее про неуспешные платежи:
	// https://yookassa.ru/developers/payment/declined-payment
	CancellationDetails CancellationDetails `json:"cancellation_details,omitempty"`
	// Данные об авторизации платежа.
	AuthorizationDetails AuthorizationDetails `json:"authorization_details,omitempty"`
}
