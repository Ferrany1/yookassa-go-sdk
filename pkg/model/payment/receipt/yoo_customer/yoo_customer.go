package yoo_customer //nolint:revive // better package readability

type YooCustomer struct {
	FullName string `json:"full_name,omitempty"`
	INN      string `json:"inn,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

func NewYooCustomer() *YooCustomer {
	return &YooCustomer{}
}

func (c *YooCustomer) WithFullName(fullName string) *YooCustomer {
	c.FullName = fullName
	return c
}

func (c *YooCustomer) WithINN(inn string) *YooCustomer {
	c.INN = inn
	return c
}

func (c *YooCustomer) WithEmail(email string) *YooCustomer {
	c.Email = email
	return c
}

func (c *YooCustomer) WithPhone(phone string) *YooCustomer {
	c.Phone = phone
	return c
}
