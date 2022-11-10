package main

import "fmt"

func main() {
	say := "bless your heart"
	fmt.Println(say)

	youMean(&say)
	fmt.Println("meaning: ", say)
}

func youMean(s *string) {
	fmt.Println("address: ", s, " contains: ", *s)
	*s = "f*** you"
}
