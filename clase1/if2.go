package main

import (
	"log"
)

func main() {
	a := 0

	if c := a % 2; c == 0 {
		log.Println("Es par")
	} else {
		log.Println("Es impar")
	}
}
