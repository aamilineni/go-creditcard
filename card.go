package main

import (
	"fmt"
	"strconv"
)

const (
	AMERICAN_EXPRESS_CARD = "American Express"
	JCB_CARD              = "JCB"
	MAESTRO_CARD          = "Maestro"
	MASTER_CARD           = "MasterCard"
	VISA_CARD             = "Visa"
	UNKNOWN_CARD          = "UNKNOWN_CARD"
)

type Card struct {
	Number CCNumber
}

type Company struct {
	Name string
}

func NewCard(no string) *Card {
	return &Card{
		Number: CCNumber(no),
	}
}

// IsValid checks whether the cc number is valid
func (c *Card) IsValid() (bool, error) {

	var sum int
	var alternate bool

	numberLen := len(c.Number)

	if numberLen < 13 || numberLen > 19 {
		return false, fmt.Errorf("INVALID CARD NUMBER")
	}

	for i := numberLen - 1; i > -1; i-- {
		mod, err := strconv.Atoi(string(c.Number[i]))
		if err != nil {
			return false, fmt.Errorf("INVALID CARD NUMBER")
		}
		if alternate {
			mod *= 2
			if mod > 9 {
				mod = (mod % 10) + 1
			}
		}

		alternate = !alternate

		sum += mod
	}

	return sum%10 == 0, nil

}

// GetCompany decides the credit card company based on the cc number
func (c *Card) GetCompany() *Company {
	ccDigits, err := c.Number.Till(4)
	if err != nil {
		return &Company{UNKNOWN_CARD}
	}

	switch {
	default:
		return &Company{UNKNOWN_CARD}
	case c.Number.IsAmexCard(ccDigits):
		return &Company{AMERICAN_EXPRESS_CARD}
	case c.Number.IsJCBCard(ccDigits):
		return &Company{JCB_CARD}
	case c.Number.IsMeastroCard(ccDigits):
		return &Company{MAESTRO_CARD}
	case c.Number.IsMasterCard(ccDigits):
		return &Company{MASTER_CARD}
	case c.Number.IsVisaCard(ccDigits):
		return &Company{VISA_CARD}
	}
}
