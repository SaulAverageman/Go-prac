package main

import (
	"fmt"
	"time"
)

type Bdata struct {
	FirstName  string
	MiddleName string
	LastName   string
	age        int
	DOB        time.Time
}

func (dataInstance Bdata) changeFName() {
	dataInstance.FirstName = "Fred"
}

func main() {
	//struct assignment
	bigman := Bdata{
		FirstName: "Hesus",
		LastName:  "McSanta",
		age:       2022,
	}

	var alsobigman Bdata
	alsobigman.age = bigman.age - 600

	fmt.Println(bigman.FirstName, bigman.MiddleName, " born on:", bigman.DOB)
	fmt.Println(alsobigman.age)
	bigman.changeFName()

	//function associated with data type
	person1 := data{
		Name: "eddy",
		Age:  23,
	}
	person2 := data{
		Name: "Reddy",
		Age:  17,
	}

	checkAssociatedFunc(person1)
	checkAssociatedFunc(person2)
}

type data struct {
	Name string
	Age  int
}

func (dataInstance *data) isAdult() bool {
	return dataInstance.Age >= 18
}

func checkAssociatedFunc(person data) {
	fmt.Println(person.Name, "is 18+?", person.isAdult())
}
