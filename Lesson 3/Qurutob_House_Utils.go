package main

import (
	"fmt"
	"strings"
	"unicode"
)

func PrintMainScreen() {
	fmt.Println("Welcome to our Qurutob House")

	fmt.Println("What kind of food do you eat")
	fmt.Println("1 - Qurutob")
	fmt.Println("2 - Qurutob with meet")
	fmt.Println("3 - Shakarob")
	fmt.Println("4 - FatirShurbo")
	fmt.Println("5 - Shurbo")
}

func CheckFood(foodID int, portion float64) bool {
	var checkPortion bool

	foodPortion := map[int]string{
		1: "0.5, 1.0, 1.5, 2.0, 2.5, 3.0",
		2: "0.5, 1.0, 1.5, 2.0, 2.5, 3.0",
		3: "0.5, 1.0, 1.5, 2.0, 2.5, 3.0",
		4: "0.5, 1.0, 1.5, 2.0, 2.5, 3.0",
		5: "1.0, 2.0, 3.0",
	}

	a := fmt.Sprintf("%.1f", portion)
	prePortion := foodPortion[foodID]
	arrPortion := strings.Split(prePortion, ", ")

	for _, _portionValue := range arrPortion {

		if a == _portionValue {
			return true
		}

	}
	return checkPortion
}

func cashlessPaymantCardNumber(cardNumber string) bool {
	var checkCardNumberF bool
	for _, digitCard := range cardNumber {
		if !unicode.IsNumber(digitCard) {
			return false
		}
	}

	if len(cardNumber) == 16 {
		return true
	}

	return checkCardNumberF

}
func cashlessPaymantPinCode(pinCode string) bool {
	var checkPinCodeF bool
	for _, digitPin := range pinCode {
		if !unicode.IsNumber(digitPin) {
			return false
		}
	}
	if len(pinCode) == 4 {
		return true
	}
	return checkPinCodeF
}
