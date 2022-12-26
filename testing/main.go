package main

import (
	"errors"
	"log"
)

func main() {
	ans, err := divide(10.0, 0)

	if err != nil {
		log.Printf("Error: %e", err)
		return
	}
	log.Printf("Answer: %f", ans)
}

func divide(x, y float32) (float32, error) {
	if y == 0 {
		return 0.0, errors.New("cannot divide by zero")
	}

	return x / y, nil
}
