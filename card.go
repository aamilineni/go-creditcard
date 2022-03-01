package main

import (
	"fmt"
	"strconv"
)

type Card struct {
	Number string
}

func NewCard(no string) *Card {
	return &Card{
		Number: no,
	}
}

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
