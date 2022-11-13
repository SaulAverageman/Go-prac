package main

import (
	"fmt"

	"./localpackage"
)

var s = 1

var PublicVar = 911   //available to outside of package
var privateVar = 2611 //available inside package

func main() {
	s := 2
	fmt.Println(s)
	sayWhat()
	localpackage.PublicFunc()

}

func sayWhat() {
	fmt.Println(s)
}
