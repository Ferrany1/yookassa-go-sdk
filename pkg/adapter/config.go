package adapter

import "errors"

const (
	paymentsEndpoint = "payments/"
	captureEndpoint  = "/capture"
	declineEndpoint  = "/cancel"
)

type Config struct {
	// Идентификатор вашего магазина.
	ShopID string `json:"shop_id"`
	// Ваш секретный ключ.
	Token string `json:"token"`
}

func (c Config) ValidateConfigFields() error {
	if len(c.ShopID) == 0 {
		return errors.New("missing ShopID value")
	}
	if len(c.Token) == 0 {
		return errors.New("missing Token value")
	}
	return nil
}
