package transfer

import (
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/amount"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/metadata"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/status"
)

type Transfer struct {
	// Идентификатор магазина, в пользу которого вы принимаете оплату.
	// Выдается ЮKassa, отображается в разделе Продавцы личного кабинета (столбец shopId).
	AccountID string `json:"account_id,omitempty"`
	// Сумма, которую необходимо перечислить магазину.
	Amount *amount.Amount `json:"amount,omitempty"`
	// Статус распределения денег между магазинами.
	Status status.Status `json:"status,omitempty"`
	// Комиссия за проданные товары и услуги, которая удерживается с магазина в вашу пользу.
	PlatformFeeAmount *amount.Amount `json:"platform_fee_amount,omitempty"`
	// Любые дополнительные данные, которые нужны вам для работы (например, номер заказа).
	// Передаются в виде набора пар «ключ-значение» и возвращаются в ответе от ЮKassa.
	// Ограничения: максимум 16 ключей, имя ключа не больше 32 символов,
	// значение ключа не больше 512 символов, тип данных — строка в формате UTF-8.
	Metadata *metadata.Metadata `json:"metadata,omitempty"`
}

func NewTransfer() *Transfer {
	return &Transfer{}
}

func (t *Transfer) WithAccountID(accountID string) *Transfer {
	t.AccountID = accountID
	return t
}

func (t *Transfer) WithAmount(amount *amount.Amount) *Transfer {
	t.Amount = amount
	return t
}

func (t *Transfer) WithStatus(status status.Status) *Transfer {
	t.Status = status
	return t
}

func (t *Transfer) WithPlatformFeeAmount(platformFeeAmount *amount.Amount) *Transfer {
	t.PlatformFeeAmount = platformFeeAmount
	return t
}

func (t *Transfer) WithMetadata(metadata *metadata.Metadata) *Transfer {
	t.Metadata = metadata
	return t
}

// Transfers Данные о распределении денег — сколько и в какой магазин нужно перевести.
// Присутствует, если вы используете Сплитование платежей.
type Transfers []*Transfer

func NewTransfers(transfers ...*Transfer) *Transfers {
	t := Transfers(transfers)
	return &t
}
