package payment

import "github.com/Ferrany1/yookassa-go-sdk-fed/pkg/model/amount"

type RawPaymentPart struct {
	// Сумма платежа.
	// Иногда партнеры ЮKassa берут с пользователя дополнительную комиссию, которая не входит в эту сумму.
	Amount amount.Amount `json:"amount,omitempty"`
	// Описание транзакции (не более 128 символов), которое вы увидите в личном кабинете ЮKassa,
	// а пользователь — при оплате. Например: «Оплата заказа № 72 для user@yoomoney.ru».
	Description string `json:"description,omitempty"`
	// Получатель платежа.
	Recipient *Recipient `json:"recipient,omitempty"`
	// Данные о сделке, в составе которой проходит платеж.
	// Присутствует, если вы проводите Безопасную сделку.
	Deal *Deal `json:"deal,omitempty"`
	// Идентификатор покупателя в вашей системе, например электронная почта или номер телефона.
	// Не более 200 символов.
	// Присутствует, если вы хотите запомнить банковскую карту и отобразить ее при повторном платеже в виджете ЮKassa.
	MerchantCustomerID string `json:"merchant_customer_id,omitempty"`
	// Данные о распределении денег — сколько и в какой магазин нужно перевести.
	// Присутствует, если вы используете Сплитование платежей.
	Transfers *Transfers `json:"transfers,omitempty"`
	// Данные для формирования чека в онлайн-кассе (для соблюдения 54-ФЗ ).
	// Необходимо передавать, если вы отправляете данные для формирования чеков по одному из сценариев:
	// Платеж и чек одновременно или Сначала чек, потом платеж.
	Receipt *Receipt `json:"receipt,omitempty"`
	// Одноразовый токен для проведения оплаты, сформированный с помощью веб или мобильного SDK.
	PaymentToken string `json:"payment_token,omitempty"`
	// Идентификатор сохраненного способа оплаты.
	PaymentMethodID string `json:"payment_method_id,omitempty"`
	// Сохранение платежных данных (с их помощью можно проводить повторные безакцептные списания.
	// Значение true инициирует создание многоразового payment_method.
	SavePaymentMethod bool `json:"save_payment_method,omitempty"`
	// Автоматический прием поступившего платежа.
	Capture bool `json:"capture,omitempty"`
	// IPv4 или IPv6-адрес пользователя. Если не указан, используется IP-адрес TCP-подключения.
	ClientIP string `json:"client_ip,omitempty"`
	// Любые дополнительные данные, которые нужны вам для работы (например, номер заказа).
	// Передаются в виде набора пар «ключ-значение» и возвращаются в ответе от ЮKassa.
	// Ограничения: максимум 16 ключей, имя ключа не больше 32 символов,
	// значение ключа не больше 512 символов, тип данных — строка в формате UTF-8.
	Metadata *Metadata `json:"metadata,omitempty"`
	// Информация о пассажирах и билетах передается при создании платежа.
	// В запросе обязательно указывается или номер билета (ticket_number),
	// или номер брони (booking_reference), если номера билета пока нет.
	AirlineTicketData *AirlineTicketData `json:"airline,omitempty"`
}

type RawPaymentRequest struct {
	RawPaymentPart
	// Данные для оплаты конкретным способом.
	// Вы можете не передавать этот объект в запросе.
	// В этом случае пользователь будет выбирать способ оплаты на стороне ЮKassa.
	PaymentMethodData *map[string]interface{} `json:"payment_method_data,omitempty"`
	// Данные, необходимые для инициирования выбранного сценария подтверждения платежа пользователем.
	Confirmation *map[string]interface{} `json:"confirmation,omitempty"`
}
