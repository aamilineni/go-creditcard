package main

import "strconv"

// CCNumber is the credit card number
type CCNumber string

// length gives the total credit card number
func (c CCNumber) length() int {
	return len(string(c))
}

// Digits used for accessing the cc number through digits
type Digits [4]int

// AtLocation returns the digits from the start to the given length
func (d *Digits) AtLocation(i int) int {
	return d[i-1]
}

// Till gives the digits of cc number for the given location
func (c CCNumber) Till(till int) (Digits, error) {
	var err error
	ccLength := c.length()
	ccDigits := Digits{}
	for i := 0; i < till; i++ {
		if i < ccLength {
			ccDigits[i], err = strconv.Atoi(string(c)[:i+1])
			if err != nil {
				return Digits{}, err
			}
		}
	}
	return ccDigits, nil
}

// IsAmexCard validates for Amex card
func (c *CCNumber) IsAmexCard(ccDigits Digits) bool {
	return (ccDigits.AtLocation(2) == 34 || ccDigits.AtLocation(2) == 37) && c.length() == 15
}

// IsJCBCard validates for JCB card
func (c *CCNumber) IsJCBCard(ccDigits Digits) bool {
	return inBetween(ccDigits.AtLocation(4), 3528, 3589) && inBetween(c.length(), 16, 19)
}

// IsMasterCard validates for Master card
func (c *CCNumber) IsMasterCard(ccDigits Digits) bool {
	return (inBetween(ccDigits.AtLocation(2), 51, 55) || inBetween(ccDigits.AtLocation(4), 2221, 2720)) && c.length() == 16
}

// IsVisaCard validates for VISA card
func (c *CCNumber) IsVisaCard(ccDigits Digits) bool {
	return ccDigits.AtLocation(1) == 4 && (c.length() == 13 || c.length() == 16 || c.length() == 19)
}

// IsMeastroCard validates for Meastro card
func (c *CCNumber) IsMeastroCard(ccDigits Digits) bool {
	return (ccDigits.AtLocation(2) == 50 || inBetween(ccDigits.AtLocation(2), 56, 58) || ccDigits.AtLocation(1) == 6) && inBetween(c.length(), 12, 19)
}

func inBetween(n, min, max int) bool {
	return n >= min && n <= max
}
