package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"os/user"
	"time"
)

func main() {
	var exitQurutobHouse bool
	var foodID int
	var portion float64
	var paymant int
	var cardNumber string
	var pinCode string
	var total float64
	var s float64

	for !exitQurutobHouse {
		var foodIDCheck bool
		var portionCheck bool
		var typeOfPayment bool
		var commandCheck bool
		var command string
		var foodName string
		var cardNumberCheck bool
		var pinCodeCheck bool
		cur, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Our waiter ", cur.Username, " will be serving you")
		PrintMainScreen()

		for !foodIDCheck {
			fmt.Scan(&foodID)

			if foodID >= 1 && foodID <= 5 {
				if foodID == 1 {
					foodName = "Qurutob 13som"
				} else if foodID == 2 {
					foodName = "Qurutob with meet 18som"
				} else if foodID == 3 {
					foodName = "Shakarob 15som"
				} else if foodID == 4 {
					foodName = "FatirShurbo 15som"
				} else if foodID == 5 {
					foodName = "Surbo 10som"
				}
				foodIDCheck = true
			} else {
				fmt.Println("Please, choose right foodID")
			}
		}

		for !portionCheck {
			fmt.Print("Choose your portion (foodID): ")
			fmt.Scan(&portion)
			if CheckFood(foodID, portion) {

				portionCheck = true
			} else {
				fmt.Println("Please, enter right portion")
			}
		}
		s = portion

		for !typeOfPayment {
			fmt.Println("Choose type of payment: ")
			fmt.Println("1 - Cashless ")
			fmt.Println("2 - Cash ")
			fmt.Scan(&paymant)
			if paymant == 1 {

				for !cardNumberCheck {
					fmt.Printf("Enter your card nubmer: ")
					fmt.Scan(&cardNumber)

					if cashlessPaymantCardNumber(cardNumber) {
						cardNumberCheck = true
					} else {
						fmt.Println("Please, enter right card number")
					}
				}
				fmt.Println()

				for !pinCodeCheck {
					fmt.Printf("Enter your pin card: ")
					fmt.Scan(&pinCode)
					if cashlessPaymantPinCode(pinCode) {
						pinCodeCheck = true
					} else {
						fmt.Println("Please, enter right pin")
					}
				}
				fmt.Println()

				if foodID == 1 {

					total = 13 * s
				} else if foodID == 2 {

					total = 18 * s
				} else if foodID == 3 {

					total = 15 * s
				} else if foodID == 4 {

					total = 15 * s
				} else if foodID == 5 {

					total = 10 * s

				}
				typeOfPayment = true
			} else if paymant == 2 {
				if foodID == 1 {

					total = 13 * s
				} else if foodID == 2 {

					total = 18 * s
				} else if foodID == 3 {

					total = 15 * s
				} else if foodID == 4 {

					total = 15 * s
				} else if foodID == 5 {

					total = 10 * s
				}
				typeOfPayment = true
			} else {
				fmt.Println("Please, enter right type of paymant")
			}
		}
		fmt.Println("----------------------------------------------")
		var h string
		h = fmt.Sprint("Total : ", foodName, " x ", portion, " = ", total)
		fmt.Println(h)
		dat := time.Now()
		fmt.Printf("Data: %02d.%02d.%d %02d:%02d:%02d\n", dat.Day(), dat.Month(), dat.Year(), dat.Hour(), dat.Minute(), dat.Second())
		fmt.Println("Casher ", cur.Username, " served for you")
		fmt.Println("OperationStatus: Successfull")
		fmt.Println("----------------------------------------------")
		hash := md5.New()
		b := []byte(h)
		fmt.Printf("%x\n", hash.Sum(b))
		hash.Write(b)
		fmt.Printf("%x\n", hash.Sum(nil))
		fmt.Println("esc- to go exit, back - to go to terminal")

		for !commandCheck {
			fmt.Scan(&command)

			if command == "esc" {
				commandCheck = true
				exitQurutobHouse = true
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
