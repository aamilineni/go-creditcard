package main

import "log"

func main() {

	card := NewCard("5105105105105100")

	isValidCard, err := card.IsValid()
	if err != nil {
		log.Fatalf("INVALID CARD")
	}
	log.Println("Is Card valid:", isValidCard)

	company := card.GetCompany()
	log.Println("Company Found as:", company.Name)
}
