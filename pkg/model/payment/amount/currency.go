package amount

type Currency string

const RubCurrency = "RUB"

func (c Currency) String() string {
	return string(c)
}
