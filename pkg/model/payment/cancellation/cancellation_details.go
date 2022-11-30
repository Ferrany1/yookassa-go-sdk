package cancellation

// CancellationDetails Комментарий к статусу canceled: кто отменил платеж и по какой причине.
// Подробнее про неуспешные платежи:
// https://yookassa.ru/developers/payment/declined-payment
type CancellationDetails struct {
	// Участник процесса платежа, который принял решение об отмене транзакции.
	// Подробнее про инициаторов отмены платежа:
	// https://yookassa.ru/developers/payments/declined-payments#cancellation-details-party
	Party string `json:"party,omitempty"`
	// Причина отмены платежа.
	// Перечень и описание возможных значений:
	// https://yookassa.ru/developers/payments/declined-payments#cancellation-details-reason
	Reason string `json:"reason,omitempty"`
}
