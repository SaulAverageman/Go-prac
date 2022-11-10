package main

import "fmt"

//package level var(need not be used)
var waste string

func main() {
	fmt.Println("ah siht here we go again!")

	age := "seven"
	fmt.Println(age)

	name, age2 := say2things()
	fmt.Println("the thing is", name, "and its ancient", age2)
}

//returns 2 things
func say2things() (string, int) {
	return "golang", 10
}
