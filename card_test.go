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
		"675964":           false,
		"675964982A438453": false,
	}

	for number, val := range fakeInputs {
		card := NewCard(number)
		isValid, _ := card.IsValid()
		if isValid != val {
			t.Errorf("error validating card : %s", number)
		}
	}

}

func TestCardCompany(t *testing.T) {

	fakeInputs := map[string]string{
		"378282246310005":  AMERICAN_EXPRESS_CARD,
		"3530111333300000": JCB_CARD,
		"6759649826438453": MAESTRO_CARD,
		"4012888888881881": VISA_CARD,
		"5105105105105100": MASTER_CARD,
		"5105105105105101": MASTER_CARD,
		"1105105105105101": UNKNOWN_CARD,
		"A759649821438453": UNKNOWN_CARD,
	}

	for number, val := range fakeInputs {
		card := NewCard(number)
		company := card.GetCompany()
		if company.Name != val {
			t.Errorf("error validating card : %s", number)
		}
	}

}
