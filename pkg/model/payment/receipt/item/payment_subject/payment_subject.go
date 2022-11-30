package payment_subject //nolint:revive // better package readability

type PaymentSubject string

const (
	// Доступно, начиная с ФФД 1.05.

	// CommodityPaymentSubject Товар
	CommodityPaymentSubject PaymentSubject = "commodity"
	// ExcisePaymentSubject Подакцизный товар
	ExcisePaymentSubject PaymentSubject = "excise"
	// JobPaymentSubject Работа
	JobPaymentSubject PaymentSubject = "job"
	// ServicePaymentSubject Услуга
	ServicePaymentSubject PaymentSubject = "service"
	// PaymentPaymentSubject Платеж
	PaymentPaymentSubject PaymentSubject = "payment"
	// CasinoPaymentSubject Платеж казино
	CasinoPaymentSubject PaymentSubject = "casino"
	// GamblingBetPaymentSubject Ставка в азартной игре
	GamblingBetPaymentSubject PaymentSubject = "gambling_bet"
	// GamblingPrizePaymentSubject Выигрыш азартной игры
	GamblingPrizePaymentSubject PaymentSubject = "gambling_prize"
	// LotteryPaymentSubject Лотерейный билет
	LotteryPaymentSubject PaymentSubject = "lottery"
	// LotteryPrizePaymentSubject Выигрыш в лотерею
	LotteryPrizePaymentSubject PaymentSubject = "lottery_prize"
	// IntellectualActivityPaymentSubject Результаты интеллектуальной деятельности
	IntellectualActivityPaymentSubject PaymentSubject = "intellectual_activity"
	// AgentCommissionPaymentSubject Агентское вознаграждение
	AgentCommissionPaymentSubject PaymentSubject = "agent_commission"
	// PropertyRightPaymentSubject Имущественное право
	PropertyRightPaymentSubject PaymentSubject = "property_right"
	// NonOperatingGainPaymentSubject Внереализационный доход
	NonOperatingGainPaymentSubject PaymentSubject = "non_operating_gain"
	// InsurancePremiumPaymentSubject Страховой сбор
	InsurancePremiumPaymentSubject PaymentSubject = "insurance_premium"
	// SalesTaxPaymentSubject Торговый сбор
	SalesTaxPaymentSubject PaymentSubject = "sales_tax"
	// ResortFeePaymentSubject Курортный сбор
	ResortFeePaymentSubject PaymentSubject = "resort_fee"
	// CompositePaymentSubject Несколько вариантов
	CompositePaymentSubject PaymentSubject = "composite"
	// AnotherPaymentSubject Другое
	AnotherPaymentSubject PaymentSubject = "another"

	// Доступно, начиная с ФФД 1.2.

	// MarkedPaymentSubject Товар, подлежащий маркировке средством идентификации, имеющим код маркировки,
	// за исключением подакцизного товара (в чеке — ТМ). Пример: обувь, духи, товары легкой промышленности
	MarkedPaymentSubject PaymentSubject = "marked"
	// NonMarkedPaymentSubject Товар, подлежащий маркировке средством идентификации, не имеющим кода маркировки,
	// за исключением подакцизного товара (в чеке — ТНМ). Пример: меховые изделия
	NonMarkedPaymentSubject PaymentSubject = "non_marked"
	// MarkedExcisePaymentSubject Подакцизный товар, подлежащий маркировке средством идентификации,
	// имеющим код маркировки (в чеке — АТМ). Пример: табак
	MarkedExcisePaymentSubject PaymentSubject = "marked_excise"
	// NonMarkedExcisePaymentSubject Подакцизный товар, подлежащий маркировке средством идентификации,
	// не имеющим кода маркировки (в чеке — АТНМ). Пример: алкогольная продукция
	NonMarkedExcisePaymentSubject PaymentSubject = "non_marked_excise"
	// FinePaymentSubject Выплата
	FinePaymentSubject PaymentSubject = "fine"
	// TaxPaymentSubject Страховые взносы
	TaxPaymentSubject PaymentSubject = "tax"
	// LienPaymentSubject Залог
	LienPaymentSubject PaymentSubject = "lien"
	// CostPaymentSubject Расход
	CostPaymentSubject PaymentSubject = "cost"
	// AgentWithdrawalsPaymentSubject Выдача денежных средств
	AgentWithdrawalsPaymentSubject PaymentSubject = "agent_withdrawals"
	// PensionInsuranceWithoutPayoutsPaymentSubject Взносы на обязательное пенсионное страхование ИП
	PensionInsuranceWithoutPayoutsPaymentSubject PaymentSubject = "pension_insurance_without_payouts"
	// PensionInsuranceWithPayoutsPaymentSubject Взносы на обязательное пенсионное страхование
	PensionInsuranceWithPayoutsPaymentSubject PaymentSubject = "pension_insurance_with_payouts"
	// HealthInsuranceWithoutPayoutsPaymentSubject Взносы на обязательное медицинское страхование ИП
	HealthInsuranceWithoutPayoutsPaymentSubject PaymentSubject = "health_insurance_without_payouts"
	// HealthInsuranceWithPayoutsPaymentSubject Взносы на обязательное медицинское страхование
	HealthInsuranceWithPayoutsPaymentSubject PaymentSubject = "health_insurance_with_payouts"
	// HealthInsurancePaymentSubject Взносы на обязательное социальное страхование
	HealthInsurancePaymentSubject PaymentSubject = "health_insurance"
)
