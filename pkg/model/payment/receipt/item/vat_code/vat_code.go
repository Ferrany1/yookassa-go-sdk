package vat_code //nolint:revive // better package readability

type VatCode int

const (
	// NoNDSVatCode Без НДС
	NoNDSVatCode VatCode = 1
	// NDS0PercentageVatCode НДС по ставке 0%
	NDS0PercentageVatCode VatCode = 2
	// NDS10PercentageVatCode НДС по ставке 10%
	NDS10PercentageVatCode VatCode = 3
	// NDS20PercentageVatCode НДС чека по ставке 20%
	NDS20PercentageVatCode VatCode = 4
	// NDSCheque10on110PercentageVatCode НДС чека по расчетной ставке 10/110
	NDSCheque10on110PercentageVatCode VatCode = 5
	// NDSCheque20on120PercentageVatCode НДС чека по расчетной ставке 20/120
	NDSCheque20on120PercentageVatCode VatCode = 6
)
