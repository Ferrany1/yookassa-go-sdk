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

func NewCard() *Card {
	return &Card{}
}

func (c *Card) WithNumber(number string) *Card {
	c.Number = number
	return c
}

func (c *Card) WithCSC(csc string) *Card {
	c.CSC = csc
	return c
}

func (c *Card) WithCardholder(cardholder string) *Card {
	c.Cardholder = cardholder
	return c
}

func (c *Card) WithFirst6(first6 string) *Card {
	c.First6 = first6
	return c
}

func (c *Card) WithLast4(last4 string) *Card {
	c.Last4 = last4
	return c
}

func (c *Card) WithExpiryMonth(expiryMonth string) *Card {
	c.ExpiryMonth = expiryMonth
	return c
}

func (c *Card) WithExpiryYear(expiryYear string) *Card {
	c.ExpiryYear = expiryYear
	return c
}

func (c *Card) WithCardType(cardType string) *Card {
	c.CardType = cardType
	return c
}

func (c *Card) WithIssuerCountry(issuerCountry string) *Card {
	c.IssuerCountry = issuerCountry
	return c
}

func (c *Card) WithIssuerName(issuerName string) *Card {
	c.IssuerName = issuerName
	return c
}
