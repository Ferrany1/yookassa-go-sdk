package confirmation

type ConfirmationType string

const (
	EmbeddedConfirmationType          ConfirmationType = "embedded"
	ExternalConfirmationType          ConfirmationType = "external"
	MobileApplicationConfirmationType ConfirmationType = "mobile_application"
	QRConfirmationType                ConfirmationType = "qr"
	RedirectConfirmationType          ConfirmationType = "redirect"
)

type GenericConfirmation interface {
	EmbeddedConfirmation |
		ExternalConfirmation |
		MobileApplicationConfirmation |
		QRConfirmation |
		RedirectConfirmation
}

type EmbeddedConfirmation struct {
	// Код сценария подтверждения.
	Type ConfirmationType `json:"type"`
	// Язык интерфейса, писем и смс, которые будет видеть или получать пользователь. Формат соответствует ISO/IEC 15897.
	// Возможные значения: ru_RU, en_US. Регистр важен.
	Locale Locale `json:"locale"`

	// Токен для инициализации платежного виджета ЮKassa.
	ConfirmationToken string `json:"confirmation_token"`
}

func NewEmbeddedConfirmation() EmbeddedConfirmation {
	return EmbeddedConfirmation{Type: EmbeddedConfirmationType}
}

func (c EmbeddedConfirmation) WithLocale(locale Locale) EmbeddedConfirmation {
	c.Locale = locale
	return c
}

func (c EmbeddedConfirmation) WithConfirmationToken(confirmationToken string) EmbeddedConfirmation {
	c.ConfirmationToken = confirmationToken
	return c
}

type ExternalConfirmation struct {
	// Код сценария подтверждения.
	Type ConfirmationType `json:"type"`
	// Язык интерфейса, писем и смс, которые будет видеть или получать пользователь. Формат соответствует ISO/IEC 15897.
	// Возможные значения: ru_RU, en_US. Регистр важен.
	Locale Locale `json:"locale"`
}

func NewExternalConfirmation() ExternalConfirmation {
	return ExternalConfirmation{Type: ExternalConfirmationType}
}

func (c ExternalConfirmation) WithLocale(locale Locale) ExternalConfirmation {
	c.Locale = locale
	return c
}

type MobileApplicationConfirmation struct {
	// Код сценария подтверждения.
	Type ConfirmationType `json:"type"`
	// Язык интерфейса, писем и смс, которые будет видеть или получать пользователь. Формат соответствует ISO/IEC 15897.
	// Возможные значения: ru_RU, en_US. Регистр важен.
	Locale Locale `json:"locale"`

	// Диплинк на мобильное приложение, в котором пользователь подтверждает платеж.
	ConfirmationURL string `json:"confirmation_url"`
}

func NewMobileApplicationConfirmation() MobileApplicationConfirmation {
	return MobileApplicationConfirmation{Type: MobileApplicationConfirmationType}
}

func (c MobileApplicationConfirmation) WithLocale(locale Locale) MobileApplicationConfirmation {
	c.Locale = locale
	return c
}

func (c MobileApplicationConfirmation) WithConfirmationURL(confirmationURL string) MobileApplicationConfirmation {
	c.ConfirmationURL = confirmationURL
	return c
}

type QRConfirmation struct {
	// Код сценария подтверждения.
	Type ConfirmationType `json:"type"`
	// Язык интерфейса, писем и смс, которые будет видеть или получать пользователь. Формат соответствует ISO/IEC 15897.
	// Возможные значения: ru_RU, en_US. Регистр важен.
	Locale Locale `json:"locale"`

	// Данные для генерации QR-кода.
	ConfirmationData string `json:"confirmation_data"`
}

func NewQRConfirmation() QRConfirmation {
	return QRConfirmation{Type: QRConfirmationType}
}

func (c QRConfirmation) WithLocale(locale Locale) QRConfirmation {
	c.Locale = locale
	return c
}

func (c QRConfirmation) WithConfirmationData(confirmationData string) QRConfirmation {
	c.ConfirmationData = confirmationData
	return c
}

type RedirectConfirmation struct {
	// Код сценария подтверждения.
	Type ConfirmationType `json:"type"`
	// Язык интерфейса, писем и смс, которые будет видеть или получать пользователь. Формат соответствует ISO/IEC 15897.
	// Возможные значения: ru_RU, en_US. Регистр важен.
	Locale Locale `json:"locale"`
	// URL, на который необходимо перенаправить пользователя для подтверждения оплаты.
	ConfirmationURL string `json:"confirmation_url"`
	// Запрос на проведение платежа с аутентификацией по 3-D Secure.
	// Будет работать, если оплату банковской картой вы по умолчанию принимаете без подтверждения платежа пользователем.
	// В остальных случаях аутентификацией по 3-D Secure будет управлять ЮKassa.
	// Если хотите принимать платежи без дополнительного подтверждения пользователем, напишите вашему менеджеру ЮKassa.
	Enforce bool `json:"enforce"`
	// URL, на который вернется пользователь после подтверждения или отмены платежа на веб-странице.
	ReturnURL string `json:"return_url"`
}

func NewRedirectConfirmation() RedirectConfirmation {
	return RedirectConfirmation{Type: RedirectConfirmationType}
}

func (c RedirectConfirmation) WithLocale(locale Locale) RedirectConfirmation {
	c.Locale = locale
	return c
}

func (c RedirectConfirmation) WithConfirmationURL(confirmationURL string) RedirectConfirmation {
	c.ConfirmationURL = confirmationURL
	return c
}

func (c RedirectConfirmation) WithEnforce() RedirectConfirmation {
	c.Enforce = true
	return c
}

func (c RedirectConfirmation) WithReturnURL(returnURL string) RedirectConfirmation {
	c.ReturnURL = returnURL
	return c
}
