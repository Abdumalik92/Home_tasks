package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var j int
	var i string
	var operator, number string
	for {
		fmt.Println("Wecome to our pay terminal")
		fmt.Println("Choose your operator")
		fmt.Println("1 - Megafon")
		fmt.Println("2 - Babilon")
		fmt.Println("3 - Tcell")
		fmt.Println("4 - Beeline")
		fmt.Scan(&i)
		for j = 10; j > 0; j-- {
			if i == "1" {
				operator = "Megafon"
			} else if i == "2" {
				operator = "Babilon"
			} else if i == "3" {
				operator = "Tcell"
			} else if i == "4" {
				operator = "Beeline"
			} else if i == "Esc" {
				return
			} else {
				fmt.Println("Please, choose right operator")
				fmt.Println(fmt.Sprint("You have ", j, " attempts"))
				fmt.Scan(&i)
			}
		}
		fmt.Print("Enter your number (Operator): ")
		fmt.Scan(&number)
		fmt.Print("\n")
		for j = 10; j > 0; j-- {
			if operator == "Megafon" {
				if (strings.HasPrefix(number, "900") || strings.HasPrefix(number, "905")) && len(number) == 9 {
					break
				} else if number == "Esc" {
					return
				} else {
					fmt.Println("Please, enter right operator number")
					fmt.Println(fmt.Sprint("You have ", j, " attempts"))
					fmt.Scan(&number)
				}
			} else if operator == "Babilon" {
				if (strings.HasPrefix(number, "988") || strings.HasPrefix(number, "918")) && len(number) == 9 {
					break
				} else if number == "Esc" {
					return
				} else {
					fmt.Println("Please, enter right operator number")
					fmt.Println(fmt.Sprint("You have ", j, " attempts"))
					fmt.Scan(&number)
				}
			} else if operator == "Tcell" {
				if (strings.HasPrefix(number, "93") || strings.HasPrefix(number, "777")) && len(number) == 9 {
					break
				} else if number == "Esc" {
					return
				} else {
					fmt.Println("Please, enter right operator number")
					fmt.Println(fmt.Sprint("You have ", j, " attempts"))
					fmt.Scan(&number)
				}
			} else if operator == "Beeline" {
				if (strings.HasPrefix(number, "919") || strings.HasPrefix(number, "917")) && len(number) == 9 {
					break
				} else if number == "Esc" {
					return
				} else {
					fmt.Println("Please, enter right operator number")
					fmt.Println(fmt.Sprint("You have ", j, " attempts"))
					fmt.Scan(&number)
				}
			}
		}
		fmt.Println("Enter sum: ")
		var sum float64

		fmt.Scan(&sum)
		//s: = strconv.FormatInt(sum)
		for j = 10; j > 0; j-- {
			if sum > 0 {
				break
			} else {
				fmt.Println("Please, enter right sum > 0")
				fmt.Println(fmt.Sprint("You have ", j, " attempts"))
				fmt.Scan(&sum)
			}
		}
		fmt.Println("----------------------------------------------")
		fmt.Println("Operator: " + operator)
		fmt.Println("Number: " + number)
		fmt.Printf("Sum: %.2f\n", sum)
		dat := time.Now()
		fmt.Printf("Data: %02d.%02d.%d %02d:%02d\n", dat.Day(), dat.Month(), dat.Year(), dat.Hour(), dat.Minute())
		fmt.Println("OperationStatus: Successfull")
		fmt.Println("----------------------------------------------")

	}
}
