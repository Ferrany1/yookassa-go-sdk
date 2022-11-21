package payment

import (
	"time"

	"github.com/Ferrany1/yookassa-go-sdk-fed/pkg/model/amount"
	"github.com/Ferrany1/yookassa-go-sdk-fed/pkg/model/status"
)

type Item struct {
	Description              string        `json:"description,omitempty"`
	Quantity                 string        `json:"quantity,omitempty"`
	Amount                   amount.Amount `json:"amount,omitempty"`
	VatCode                  int           `json:"vat_code,omitempty"`
	PaymentSubject           string        `json:"payment_subject,omitempty"`
	PaymentMode              string        `json:"payment_mode,omitempty"`
	ProductCode              string        `json:"product_code,omitempty"`
	CountryOfOriginCode      string        `json:"country_of_origin_code,omitempty"`
	CustomsDeclarationNumber string        `json:"customs_declaration_number,omitempty"`
	Excise                   string        `json:"excise,omitempty"`
}

type Items []*Item

type YooCustomer struct {
	FullName string `json:"full_name,omitempty"`
	INN      string `json:"inn,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

type Recipient struct {
	// Идентификатор магазина в ЮKassa.
	AccountID string `json:"account_id,omitempty"`
	// Идентификатор субаккаунта.
	// Используется для разделения потоков платежей в рамках одного аккаунта.
	GatewayID string `json:"gateway_id,omitempty"`
}

type Deal struct {
	// Идентификатор сделки.
	ID string `json:"id,omitempty"`
	// Данные о распределении денег.
	Settlements Settlements `json:"settlements,omitempty"`
}

type Settlement struct {
	// Тип операции. Фиксированное значение: payout — выплата продавцу.
	Type string `json:"type,omitempty"`
	// Сумма вознаграждения продавца.
	Amount amount.Amount `json:"amount,omitempty"`
}

type Settlements []*Settlement

type Transfer struct {
	// Идентификатор магазина, в пользу которого вы принимаете оплату.
	// Выдается ЮKassa, отображается в разделе Продавцы личного кабинета (столбец shopId).
	AccountID string `json:"account_id,omitempty"`
	// Сумма, которую необходимо перечислить магазину.
	Amount amount.Amount `json:"amount,omitempty"`
	// Статус распределения денег между магазинами.
	Status status.Status `json:"status,omitempty"`
	// Комиссия за проданные товары и услуги, которая удерживается с магазина в вашу пользу.
	PlatformFeeAmount amount.Amount `json:"platform_fee_amount,omitempty"`
	// Любые дополнительные данные, которые нужны вам для работы (например, номер заказа).
	// Передаются в виде набора пар «ключ-значение» и возвращаются в ответе от ЮKassa.
	// Ограничения: максимум 16 ключей, имя ключа не больше 32 символов,
	// значение ключа не больше 512 символов, тип данных — строка в формате UTF-8.
	Metadata Metadata `json:"metadata,omitempty"`
}

type Transfers []*Transfer

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

type AuthorizationDetails struct {
	// Retrieval Reference Number — уникальный идентификатор транзакции в системе эмитента.
	// Используется при оплате банковской картой.
	RRN string `json:"rrn,omitempty"`
	// Код авторизации банковской карты.
	// Выдается эмитентом и подтверждает проведение авторизации.
	AuthCode string `json:"auth_code,omitempty"`
}

type Metadata map[string]string

type Receipt struct {
	Customer      YooCustomer `json:"customer,omitempty"`
	Items         Items       `json:"items,omitempty"`
	TaxSystemCode int         `json:"tax_system_code,omitempty"`
	Phone         string      `json:"phone,omitempty"`
	Email         string      `json:"email,omitempty"`
}

type AirlineTicketData struct {
	TicketNumber     string                      `json:"ticket_number,omitempty"`
	BookingReference string                      `json:"booking_reference,omitempty"`
	Passengers       AirlineTicketPassengersData `json:"passengers,omitempty"`
	Legs             AirlineTicketLegsData       `json:"legs,omitempty"`
}

type AirlineTicketPassengerData struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

type AirlineTicketPassengersData []*AirlineTicketPassengerData

type AirlineTicketLegData struct {
	DepartureAirport   string    `json:"departure_airport,omitempty"`
	DepartureDate      time.Time `json:"departure_date,omitempty"`
	DestinationAirport string    `json:"destination_airport,omitempty"`
	CarrierCode        string    `json:"carrier_code,omitempty"`
}

// Информация о перелетах. Перелет — это фрагмент маршрута.
// Если пользователь летит без пересадки, это один перелет. Если есть одна пересадка — два перелета.
type AirlineTicketLegsData []*AirlineTicketLegData
