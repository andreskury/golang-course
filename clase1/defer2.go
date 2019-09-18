package main

import (
	"log"
)

func listo(msg string) {
	log.Printf("Se terminó la cosa %s\n", msg)
}

func main() {
	defer listo("1")
	log.Println("Esto va antes")
	defer listo("2")
	log.Println("Esto también va antes")
	defer listo("3")
}
