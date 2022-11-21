package payment_method //nolint:revive // better package readability

type Card struct {
	Number        string `json:"number,omitempty"`
	CSC           string `json:"csc,omitempty"`
	Cardholder    string `json:"cardholder,omitempty"`
	First6        string `json:"first6,omitempty"`
	Last4         string `json:"last4,omitempty"`
	ExpiryMonth   string `json:"expiry_month,omitempty"`
	ExpiryYear    string `json:"expiry_year,omitempty"`
	CardType      string `json:"card_type,omitempty"`
	IssuerCountry string `json:"issuer_country,omitempty"`
	IssuerName    string `json:"issuer_name,omitempty"`
}
