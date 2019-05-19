package main

import (
	"fmt"
	"time"
)

func main() {
	var exitTerminal bool
	var operator int
	var number string
	var sum float32

	for !exitTerminal {
		var operatorCheck bool
		var numberCheck bool
		var sumCheck bool
		var commandCheck bool
		var command string
		var operatorName string

		PrintMainScreen()

		for !operatorCheck {
			fmt.Scan(&operator)
			if operator >= 1 && operator <= 4 {
				if operator == 1 {
					operatorName = "Megafon"
				} else if operator == 2 {
					operatorName = "Babilon"
				} else if operator == 3 {
					operatorName = "Tcell"
				} else if operator == 4 {
					operatorName = "Beeline"
				}
				operatorCheck = true
			} else {
				fmt.Println("Please, choose right operator")
			}
		}

		for !numberCheck {
			fmt.Print("Enter your number (Operator): ")
			fmt.Scan(&number)
			if CheckNumber(operator, number) {

				numberCheck = true
			} else {
				fmt.Println("Please, enter right number")
			}

		}

		for !sumCheck {
			fmt.Println("Enter sum: ")
			fmt.Scan(&sum)

			if sum > 0 {
				sumCheck = true
			} else {
				fmt.Println("Please, enter right sum")
			}
		}
		fmt.Println("----------------------------------------------")
		fmt.Println("Operator: " + operatorName)
		fmt.Println("Number: " + number)
		fmt.Printf("Sum: %.2f\n", sum)
		dat := time.Now()
		fmt.Printf("Data: %02d.%02d.%d %02d:%02d:%02d\n", dat.Day(), dat.Month(), dat.Year(), dat.Hour(), dat.Minute(), dat.Second())
		fmt.Println("OperationStatus: Successfull")
		fmt.Println("----------------------------------------------")

		fmt.Println("esc- to go exit, back - to go to terminal")

		for !commandCheck {
			fmt.Scan(&command)

			if command == "esc" {
				commandCheck = true
				exitTerminal = true
			} else if command == "back" {
				commandCheck = true
				continue
			} else {
				fmt.Println("command not recongize")
			}
		}
	}

	fmt.Println("application terminated")
}
