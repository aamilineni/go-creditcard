package main

import (
	"testing"
)

func TestCardValid(t *testing.T) {

	fakeInputs := map[string]bool{
		"5237251624778133": true,
		"378282246310005":  true,
		"378282246310001":  false,
		"6759649826438453": true,
	}

	for number, val := range fakeInputs {
		card := NewCard(number)
		isValid, err := card.IsValid()
		if err != nil {
			t.Errorf("err : %+v : %s", err, number)
		}
		if isValid != val {
			t.Errorf("error validating card : %s", number)
		}
	}

}
