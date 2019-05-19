package main

import (
	"fmt"
	"strings"
	"unicode"
)

func PrintMainScreen() {
	fmt.Println("Welcome to our pay terminal!")
	fmt.Println("Choose your operator:")
	fmt.Println("1 - Megafon")
	fmt.Println("2 - Babilon")
	fmt.Println("3 - Tcell")
	fmt.Println("4 - Beeline")
}

//Checking number operator
func CheckNumber(operatorID int, number string) bool {
	var checkPrefix bool
	operatorPreffix := map[int]string{
		1: "90,8888",
		2: "918,98",
		3: "93,55",
		4: "919,917",
	}
	//Is number digit or not
	for _, digit := range number {
		if !unicode.IsNumber(digit) {

			return false
		}
	}
	if len(number) != 9 {
		return false
	}

	prefix := operatorPreffix[operatorID]
	arrPrefix := strings.Split(prefix, ",")

	for _, _prefixValue := range arrPrefix {

		checkPrefix = strings.HasPrefix(number, _prefixValue)

		if checkPrefix {
			return true
		}
	}

	return checkPrefix
}
