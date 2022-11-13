package main

import (
	"fmt"
	"time"
)

type Bdata struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	DOB        time.Time
}

func main() {
	bigman := Bdata{
		FirstName: "Hesus",
		LastName:  "McSanta",
		Age:       2022,
	}

	fmt.Println(bigman.FirstName, bigman.MiddleName, " born on:", bigman.DOB)
}
