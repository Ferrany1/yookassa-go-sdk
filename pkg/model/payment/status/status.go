package status

// Status Статус платежа.
// Подробнее про жизненный цикл платежа: https://yookassa.ru/developers/payment/payment-process#lifecycle.
type Status string

const (
	Pending           Status = "pending"
	WaitingForCapture Status = "waiting_for_capture"
	Succeeded         Status = "succeeded"
	Canceled          Status = "canceled"
)
