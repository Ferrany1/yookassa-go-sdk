package industry_details //nolint:revive // better package readability

import "time"

type IndustryDetail struct {
	FederalID      string    `json:"federal_id"`
	DocumentDate   time.Time `json:"document_date"`
	DocumentNumber string    `json:"document_number"`
	Value          string    `json:"value"`
}

func NewIndustryDetail() *IndustryDetail {
	return &IndustryDetail{}
}

func (d *IndustryDetail) WithFederalID(federalID string) *IndustryDetail {
	d.FederalID = federalID
	return d
}

func (d *IndustryDetail) WithDocumentDate(documentDate time.Time) *IndustryDetail {
	d.DocumentDate = documentDate
	return d
}

func (d *IndustryDetail) WithDocumentNumber(documentNumber string) *IndustryDetail {
	d.DocumentNumber = documentNumber
	return d
}

func (d *IndustryDetail) WithValue(value string) *IndustryDetail {
	d.Value = value
	return d
}

type IndustryDetails []*IndustryDetail
