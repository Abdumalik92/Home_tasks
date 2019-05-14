package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	var x int
	var z, a, s, p int
	x = 35
	a = x / 10
	z = x % 10
	fmt.Printf("Число десятков = %d", a)
	fmt.Println()
	fmt.Printf("Число единиц = %d", z)
	fmt.Println()
	s = a + z
	fmt.Printf("Сумма цифр = %d", s)
	fmt.Println()
	p = a * z
	fmt.Printf("Произведение цифр = %d", p)
	fmt.Println()
}
