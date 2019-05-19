package main

import (
	"fmt"
	"log"
	"os/user"
	"strings"
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
	arrPortion := strings.Split(prePortion, ",")

	for _, _portionValue := range arrPortion {

		fmt.Println(_portionValue)
		if a == _portionValue {
			return true
		}
		fmt.Println(a)
	}
	return checkPortion
}

func cashlessPaymantCardNumber(cardNumber string) bool {
	var checkCardNumber bool
	/*for _, digitCard := range cardNumber {
		if !unicode.IsNumber(digitCard) {
			return false
		}
	}*/
	fmt.Println(checkCardNumber)
	if len(cardNumber) != 16 {
		return false
	}
	fmt.Println(checkCardNumber == true)
	return checkCardNumber
}
func cashlessPaymantPinCode(pinCode string) bool {
	var checkPinCode bool
	/*for _, digitPin := range pinCode {
		if !unicode.IsNumber(digitPin) {
			checkPinCode = false
		}
	}*/
	if len(pinCode) != 4 {
		checkPinCode = false
	}
	return checkPinCode
}
