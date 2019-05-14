package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	var k int
	var n int
	k = 1
	n = k % 7
	//Задание а)
	if n == 1 {
		fmt.Printf("Понедельник")
	}
	//Задание б)
	if n == 2 {
		fmt.Printf("Вторник")
	}
	//Задание в)
	if n == 3 {
		fmt.Printf("Среда d=3")
	}
	if n == 4 {
		fmt.Printf("Четвег d=4")
	}
	if n == 5 {
		fmt.Printf("Пятница d=5")
	}
	if n == 6 {
		fmt.Printf("Субота d=6")
	}
	if n == 0 {
		fmt.Printf("Воскресение d=7")
	}

}
