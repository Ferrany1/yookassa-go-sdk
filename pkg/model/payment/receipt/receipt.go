package receipt

import (
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/receipt/industry_details"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/receipt/item"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/receipt/operational_details"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/receipt/tax_system_code"
	"github.com/Ferrany1/yookassa-go-sdk/pkg/model/payment/receipt/yoo_customer"
)

// Receipt Данные для формирования чека в онлайн-кассе (для соблюдения 54-ФЗ ).
// Необходимо передавать, если вы отправляете данные для формирования чеков по одному из сценариев:
// Платеж и чек одновременно или Сначала чек, потом платеж.
type Receipt struct {
	Customer           yoo_customer.YooCustomer               `json:"customer,omitempty"`
	Items              item.Items                             `json:"items,omitempty"`
	TaxSystemCode      tax_system_code.TaxSystemCode          `json:"tax_system_code,omitempty"`
	IndustryDetails    industry_details.IndustryDetails       `json:"receipt_industry_details,omitempty"`
	OperationalDetails operational_details.OperationalDetails `json:"receipt_operational_details,omitempty"`
}

func NewReceipt() *Receipt {
	return &Receipt{}
}

func (r *Receipt) WithCustomer(customer *yoo_customer.YooCustomer) *Receipt {
	r.Customer = *customer
	return r
}

func (r *Receipt) WithItems(items ...*item.Item) *Receipt {
	r.Items = items
	return r
}

func (r *Receipt) WithTaxSystemCode(taxSystemCode tax_system_code.TaxSystemCode) *Receipt {
	r.TaxSystemCode = taxSystemCode
	return r
}

func (r *Receipt) WithIndustryDetails(industryDetails ...*industry_details.IndustryDetail) *Receipt {
	r.IndustryDetails = industryDetails
	return r
}

func (r *Receipt) WithOperationalDetails(operationalDetails ...*operational_details.OperationalDetail) *Receipt {
	r.OperationalDetails = operationalDetails
	return r
}
