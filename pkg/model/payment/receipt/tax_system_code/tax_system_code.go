package tax_system_code //nolint:revive // better package readability

type TaxSystemCode int

const (
	// GeneralTaxSystemCode Общая система налогообложения
	GeneralTaxSystemCode TaxSystemCode = 1
	// SimplifiedIncomeTaxSystemCode Упрощенная (УСН, доходы)
	SimplifiedIncomeTaxSystemCode TaxSystemCode = 2
	// SimplifiedIncomeMinusExpensesTaxSystemCode Упрощенная (УСН, доходы минус расходы)
	SimplifiedIncomeMinusExpensesTaxSystemCode TaxSystemCode = 3
	// SingleTaxOnImputedIncomeTaxSystemCode Единый налог на вмененный доход (ЕНВД)
	SingleTaxOnImputedIncomeTaxSystemCode TaxSystemCode = 4
	// SingleAgriculturalTaxTaxSystemCode Единый сельскохозяйственный налог (ЕСН)
	SingleAgriculturalTaxTaxSystemCode TaxSystemCode = 5
	// PatentTaxationSystemTaxSystemCode Патентная система налогообложения
	PatentTaxationSystemTaxSystemCode TaxSystemCode = 6
)
