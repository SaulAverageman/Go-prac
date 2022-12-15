package main

import "fmt"

func main() {
	mymap := make(map[string]int)
	mymap["one"] = 1
	fmt.Println(mymap["one"])

	change(mymap, "one", 11)
	fmt.Println(mymap["one"])
}

func change(mymap map[string]int, key string, value int) {
	mymap[key] = value
}
