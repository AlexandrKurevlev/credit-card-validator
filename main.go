package main

import (
	"fmt"
	"strings"
)

type Bank struct {
	Name   string
	Prefix string
}

func DetectBank(cardNumber string, banks []Bank) *Bank {
	if len(cardNumber) == 0 {
		return nil
	}

	for i := range banks {
		if strings.HasPrefix(cardNumber, banks[i].Prefix) {
			return &banks[i]
		}
	}

	return nil
}

func LuhnCheck(cardNumber string) bool {
	if len(cardNumber) == 0 {
		return false
	}

	total := 0
	isSecond := false
	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')
		if digit < 0 || digit > 9 {
			return false
		}

		if isSecond {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		total += digit
		isSecond = !isSecond
	}

	return total%10 == 0
}

func main() {
	banks := []Bank{
		{Name: "Lunar Bank", Prefix: "4000"},
		{Name: "Mars Credit Union", Prefix: "5000"},
		{Name: "Venus Express", Prefix: "6000"},
		{Name: "Saturn Ring", Prefix: "7000"},
		{Name: "Jupiter Trust", Prefix: "8000"},
	}

	testCard := "4000123456789017"
	detectedBank := DetectBank(testCard, banks)
	bankName := "не определен"
	if detectedBank != nil {
		bankName = detectedBank.Name
	}
	fmt.Println("Валиден по Луне:", LuhnCheck(testCard))
	fmt.Println("Банк:", bankName)
}
