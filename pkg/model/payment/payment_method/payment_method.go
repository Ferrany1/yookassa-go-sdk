package payment_method //nolint:revive // better package readability

import "github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/amount"

type PaymentMethodType string

const ( // Payment type.
	AlfabankPaymentMethodType      PaymentMethodType = "alfabank"
	MobileBalancePaymentMethodType PaymentMethodType = "mobile_balance"
	BankCardPaymentMethodType      PaymentMethodType = "bank_card"
	InstallmentsPaymentMethodType  PaymentMethodType = "installments"
	CashPaymentMethodType          PaymentMethodType = "cash"
	B2BSberbankPaymentMethodType   PaymentMethodType = "b2b_sberbank"
	SberbankPaymentMethodType      PaymentMethodType = "sberbank"
	TinkoffBankPaymentMethodType   PaymentMethodType = "tinkoff_bank"
	YooMoneyPaymentMethodType      PaymentMethodType = "yoo_money"
	ApplePayPaymentMethodType      PaymentMethodType = "apple_pay"
	GooglePayPaymentMethodType     PaymentMethodType = "google_pay"
	QiwiPaymentMethodType          PaymentMethodType = "qiwi"
	WeChatPaymentMethodType        PaymentMethodType = "wechat"
	WebmoneyPaymentMethodType      PaymentMethodType = "webmoney"
)

type GenericPaymentMethod interface {
	AlfabankPaymentMethod |
		MobileBalancePaymentMethod |
		BankCardPaymentMethod |
		InstallmentsPaymentMethod |
		CashPaymentMethod |
		B2BSberbankPaymentMethod |
		SberbankPaymentMethod |
		TinkoffBankPaymentMethod |
		YooMoneyPaymentMethod |
		ApplePayPaymentMethod |
		GooglePayPaymentMethod |
		QiwiPaymentMethod |
		WeChatPaymentMethod |
		WebMoneyPaymentMethod
}

type BasePaymentMethod struct {
	// Код способа оплаты.
	Type PaymentMethodType `json:"type"`
	// Идентификатор способа оплаты.
	ID string `json:"id"`
	// С помощью сохраненного способа оплаты можно проводить безакцептные списания.
	Saved bool `json:"saved"`
	// Название способа оплаты.
	Title string `json:"title"`
}

func NewBasePaymentMethod() BasePaymentMethod {
	return BasePaymentMethod{}
}

func (m *BasePaymentMethod) WithType(methodType PaymentMethodType) *BasePaymentMethod {
	m.Type = methodType
	return m
}

func (m *BasePaymentMethod) WithID(id string) *BasePaymentMethod {
	m.ID = id
	return m
}

func (m *BasePaymentMethod) WithSaved() *BasePaymentMethod {
	m.Saved = true
	return m
}

func (m *BasePaymentMethod) WithTitle(title string) *BasePaymentMethod {
	m.Title = title
	return m
}

type AlfabankPaymentMethod struct {
	BasePaymentMethod
	// Логин пользователя в Альфа-Клике (привязанный телефон или дополнительный логин).
	Login *string `json:"login"`
}

func NewAlfabankPaymentMethod() AlfabankPaymentMethod {
	return AlfabankPaymentMethod{BasePaymentMethod: BasePaymentMethod{Type: AlfabankPaymentMethodType}}
}

func (m *AlfabankPaymentMethod) WithLogin(login string) *AlfabankPaymentMethod {
	m.Login = &login
	return m
}

type MobileBalancePaymentMethod struct {
	BasePaymentMethod
	// Телефон, с баланса которого осуществляется платеж. Указывается в формате ITU-T E.164, например 79000000000.
	Phone *string `json:"phone"`
}

func NewMobileBalancePaymentMethod() MobileBalancePaymentMethod {
	return MobileBalancePaymentMethod{BasePaymentMethod: BasePaymentMethod{Type: MobileBalancePaymentMethodType}}
}

func (m *MobileBalancePaymentMethod) WithPhone(phone string) *MobileBalancePaymentMethod {
	m.Phone = &phone
	return m
}

type BankCardPaymentMethod struct {
	BasePaymentMethod
	// Данные банковской карты.
	Card *Card `json:"card"`
}

func NewBankCardPaymentMethod() BankCardPaymentMethod {
	return BankCardPaymentMethod{BasePaymentMethod: BasePaymentMethod{Type: BankCardPaymentMethodType}}
}

func (m *BankCardPaymentMethod) WithCard(card *Card) *BankCardPaymentMethod {
	m.Card = card
	return m
}

type InstallmentsPaymentMethod struct {
	BasePaymentMethod
	// Телефон, с баланса которого осуществляется платеж. Указывается в формате ITU-T E.164, например 79000000000.
	// Поле можно оставить пустым: пользователь сможет заполнить его при оплате на стороне ЮKassa.
	Phone *string `json:"phone"`
}

func NewInstallmentsPaymentMethod() InstallmentsPaymentMethod {
	return InstallmentsPaymentMethod{BasePaymentMethod: BasePaymentMethod{Type: InstallmentsPaymentMethodType}}
}

func (m *InstallmentsPaymentMethod) WithPhone(phone string) *InstallmentsPaymentMethod {
	m.Phone = &phone
	return m
}

type CashPaymentMethod struct {
	BasePaymentMethod
}

func NewCashPaymentMethod() CashPaymentMethod {
	return CashPaymentMethod{BasePaymentMethod: BasePaymentMethod{Type: CashPaymentMethodType}}
}

type B2BSberbankPaymentMethod struct {
	BasePaymentMethod

	PayerBankDetails *PayerBankDetails `json:"payer_bank_details"`
	// Назначение платежа (не больше 210 символов).
	PaymentPurpose *string `json:"payment_purpose"`
	// Данные о налоге на добавленную стоимость (НДС). Платеж может облагаться и не облагаться НДС.
	// Товары могут облагаться по одной ставке НДС или по разным.
	VATData *VATData `json:"vat_data"`
}

func NewB2BSberbankPaymentMethod() B2BSberbankPaymentMethod {
	return B2BSberbankPaymentMethod{BasePaymentMethod: BasePaymentMethod{Type: B2BSberbankPaymentMethodType}}
}

func (m *B2BSberbankPaymentMethod) WithPayerBankDetails(payerBankDetails *PayerBankDetails) *B2BSberbankPaymentMethod {
	m.PayerBankDetails = payerBankDetails
	return m
}

func (m *B2BSberbankPaymentMethod) WithPaymentPurpose(paymentPurpose string) *B2BSberbankPaymentMethod {
	m.PaymentPurpose = &paymentPurpose
	return m
}

func (m *B2BSberbankPaymentMethod) WithVATData(vatData *VATData) *B2BSberbankPaymentMethod {
	m.VATData = vatData
	return m
}

type PayerBankDetails struct {
	// Полное наименование организации.
	FullName string `json:"full_name"`
	// Сокращенное наименование организации.
	ShortName string `json:"short_name"`
	// Адрес организации.
	Address string `json:"address"`
	// Индивидуальный налоговый номер (ИНН) организации.
	INN string `json:"inn"`
	// Код причины постановки на учет (КПП) организации.
	KPP string `json:"kpp"`
	// Наименование банка организации.
	BankName string `json:"bank_name"`
	// Отделение банка организации.
	BankBranch string `json:"bank_branch"`
	// Банковский идентификационный код (БИК) банка организации.
	BankBIK string `json:"bank_bik"`
	// Номер счета организации.
	Account string `json:"account"`
}

type VATData struct {
	// Код способа расчета НДС.
	Type string `json:"type"`
	// Сумма НДС.
	Amount amount.Amount `json:"amount"`
	// Налоговая ставка (в процентах).
	Rate string `json:"rate"`
}

type SberbankPaymentMethod struct {
	BasePaymentMethod
	// Телефон пользователя, на который зарегистрирован аккаунт в СберБанк Онлайн/SberPay.
	// Указывается в формате ITU-T E.164, например 79000000000.
	Phone *string `json:"phone"`
}

func NewSberbankPaymentMethod() SberbankPaymentMethod {
	return SberbankPaymentMethod{BasePaymentMethod: BasePaymentMethod{Type: SberbankPaymentMethodType}}
}

func (m *SberbankPaymentMethod) WithPhone(phone string) *SberbankPaymentMethod {
	m.Phone = &phone
	return m
}

type TinkoffBankPaymentMethod struct {
	BasePaymentMethod
}

func NewTinkoffBankPaymentMethod() TinkoffBankPaymentMethod {
	return TinkoffBankPaymentMethod{BasePaymentMethod: BasePaymentMethod{Type: TinkoffBankPaymentMethodType}}
}

type YooMoneyPaymentMethod struct {
	BasePaymentMethod
	// Номер кошелька ЮMoney, из которого заплатил пользователь.
	AccountNumber *string `json:"account_number"`
}

func NewYooMoneyPaymentMethod() YooMoneyPaymentMethod {
	return YooMoneyPaymentMethod{BasePaymentMethod: BasePaymentMethod{Type: YooMoneyPaymentMethodType}}
}

func (m *YooMoneyPaymentMethod) WithAccountNumber(accountNumber string) *YooMoneyPaymentMethod {
	m.AccountNumber = &accountNumber
	return m
}

type ApplePayPaymentMethod struct {
	BasePaymentMethod
	// Содержимое поля paymentData из объекта PKPaymentToken (платежный токен Apple Pay). Приходит в формате Base64.
	PaymentData *string `json:"payment_data"`
}

func NewApplePayPaymentMethod() ApplePayPaymentMethod {
	return ApplePayPaymentMethod{BasePaymentMethod: BasePaymentMethod{Type: ApplePayPaymentMethodType}}
}

func (m *ApplePayPaymentMethod) WithPaymentData(paymentData string) *ApplePayPaymentMethod {
	m.PaymentData = &paymentData
	return m
}

type GooglePayPaymentMethod struct {
	BasePaymentMethod
	// Требуется криптограмма Payment Token Cryptography для проведения оплаты через Google Pay.
	// Чтобы ее получить, необходимо вызвать метод PaymentData.getPaymentMethodToken().getToken()
	// в мобильном приложении на устройстве пользователя.
	PaymentMethodToken *string `json:"payment_method_token"`
}

func NewGooglePayPaymentMethod() GooglePayPaymentMethod {
	return GooglePayPaymentMethod{BasePaymentMethod: BasePaymentMethod{Type: GooglePayPaymentMethodType}}
}

func (m *GooglePayPaymentMethod) WithPaymentData(paymentMethodToken string) *GooglePayPaymentMethod {
	m.PaymentMethodToken = &paymentMethodToken
	return m
}

type QiwiPaymentMethod struct {
	BasePaymentMethod
	// Телефон, на который зарегистрирован аккаунт в QIWI. Указывается в формате ITU-T E.164, например 79000000000.
	Phone *string `json:"phone"`
}

func NewQiwiPaymentMethod() QiwiPaymentMethod {
	return QiwiPaymentMethod{BasePaymentMethod: BasePaymentMethod{Type: QiwiPaymentMethodType}}
}

func (m *QiwiPaymentMethod) WithPhone(phone string) *QiwiPaymentMethod {
	m.Phone = &phone
	return m
}

type WeChatPaymentMethod struct {
	BasePaymentMethod
}

func NewWeChatPaymentMethod() WeChatPaymentMethod {
	return WeChatPaymentMethod{BasePaymentMethod: BasePaymentMethod{Type: WeChatPaymentMethodType}}
}

type WebMoneyPaymentMethod struct {
	BasePaymentMethod
}

func NewWebMoneyPaymentMethod() WebMoneyPaymentMethod {
	return WebMoneyPaymentMethod{BasePaymentMethod: BasePaymentMethod{Type: WebmoneyPaymentMethodType}}
}
