package main

import "fmt"

func main(){
	fmt.Println("Hello world!")
	var x int
	var y int
	x=9604
	y=x/3600
	fmt.Printf("Полных часов прошло %d",y)
	fmt.Println()
	y=x-y*3600
	var s int
	s=y/60
	fmt.Printf("Полных минут прошло после очередного часа %d",s)
	fmt.Println()
	var a int
	a=y-s*60
	fmt.Printf("Полных секунд прошло после очередной минуты %d",a)

}