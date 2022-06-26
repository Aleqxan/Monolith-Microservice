package price

import (
	"errors"
)

var (
	ErrPriceToLow      = errors.New("Price must be greater than 0")
	ErrInvalidCurrency = errors.New("Invalid currency")
)

type Price struct {
	cents    uint
	currency string
}

func NewPrice(cents uint, currency string) (Price, error) {

	if cents <= 0 {
		return Price{}, ErrPriceToLow
	}

	if len(currency) != 3 {
		return Price{}, ErrInvalidCurrency
	}

	return Price{cents, currency}, nil
}

func NewPriceP(cents uint, currency string) (Price, error) {
	p, err := NewPrice(cents, currency)

	if err != nil {
		panic(err)
	}

	return p
}

func (p Price) Cents() uint {
	return p.cents
}

func (p Price) Currency() string {
	return p.currency
}
