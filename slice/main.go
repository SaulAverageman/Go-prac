package main

import (
	"fmt"
	"sort"
)

func main() {
	var name []string
	name = append(name, "a")
	name = append(name, "b")
	fmt.Println(name)

	name2 := []string{"c", "d"}
	fmt.Println(name2)

	nums := []int{4, 6, 1, 4, 0}
	sort.Ints(nums)
	fmt.Println(nums)
	fmt.Println(nums[:])

}
