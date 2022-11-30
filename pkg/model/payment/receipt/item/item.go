package item

import (
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/amount"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/receipt/industry_details"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/receipt/item/country_of_origin_code"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/receipt/item/mark_code_info"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/receipt/item/mark_quantity"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/receipt/item/payment_mode"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/receipt/item/payment_subject"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/receipt/item/vat_code"
)

type Item struct {
	Description                   string                                     `json:"description,omitempty"`
	Amount                        amount.Amount                              `json:"amount,omitempty"`
	VatCode                       vat_code.VatCode                           `json:"vat_code,omitempty"`
	Quantity                      string                                     `json:"quantity,omitempty"`
	Measure                       string                                     `json:"measure,omitempty"`
	MarkQuantity                  mark_quantity.MarkQuantity                 `json:"mark_quantity,omitempty"`
	PaymentSubject                payment_subject.PaymentSubject             `json:"payment_subject,omitempty"`
	PaymentMode                   payment_mode.PaymentMode                   `json:"payment_mode,omitempty"`
	CountryOfOriginCode           country_of_origin_code.CountryOfOriginCode `json:"country_of_origin_code,omitempty"`
	CustomsDeclarationNumber      string                                     `json:"customs_declaration_number,omitempty"`
	Excise                        string                                     `json:"excise,omitempty"`
	ProductCode                   string                                     `json:"product_code,omitempty"`
	MarkCodeInfo                  mark_code_info.MarkCodeInfo                `json:"mark_code_info"`
	MarkMode                      string                                     `json:"mark_mode"`
	PaymentSubjectIndustryDetails industry_details.IndustryDetails           `json:"payment_subject_industry_details"`
}

func NewItem() *Item {
	return &Item{}
}

func (i *Item) WithDescription(description string) *Item {
	i.Description = description
	return i
}

func (i *Item) WithAmount(amount *amount.Amount) *Item {
	i.Amount = *amount
	return i
}

func (i *Item) WithVatCode(vatCode vat_code.VatCode) *Item {
	i.VatCode = vatCode
	return i
}

func (i *Item) WithQuantity(quantity string) *Item {
	i.Quantity = quantity
	return i
}

func (i *Item) WithMeasure(measure string) *Item {
	i.Measure = measure
	return i
}

func (i *Item) WithMarkQuantity(markQuantity *mark_quantity.MarkQuantity) *Item {
	i.MarkQuantity = *markQuantity
	return i
}

func (i *Item) WithPaymentSubject(paymentSubject payment_subject.PaymentSubject) *Item {
	i.PaymentSubject = paymentSubject
	return i
}

func (i *Item) WithPaymentMode(paymentMode payment_mode.PaymentMode) *Item {
	i.PaymentMode = paymentMode
	return i
}

func (i *Item) WithCountryOfOriginCode(countryOfOriginCode country_of_origin_code.CountryOfOriginCode) *Item {
	i.CountryOfOriginCode = countryOfOriginCode
	return i
}

func (i *Item) WithCustomsDeclarationNumber(customsDeclarationNumber string) *Item {
	i.CustomsDeclarationNumber = customsDeclarationNumber
	return i
}

func (i *Item) WithExcise(excise string) *Item {
	i.Excise = excise
	return i
}

func (i *Item) WithProductCode(productCode string) *Item {
	i.ProductCode = productCode
	return i
}

func (i *Item) WithMarkCodeInfo(markCodeInfo *mark_code_info.MarkCodeInfo) *Item {
	i.MarkCodeInfo = *markCodeInfo
	return i
}

func (i *Item) WithMarkMode() *Item {
	i.MarkMode = "0"
	return i
}

func (i *Item) WithPaymentSubjectIndustryDetails(
	paymentSubjectIndustryDetails ...*industry_details.IndustryDetail,
) *Item {
	i.PaymentSubjectIndustryDetails = paymentSubjectIndustryDetails
	return i
}

type Items []*Item
