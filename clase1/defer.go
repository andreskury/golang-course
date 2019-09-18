package main

import (
	"log"
)

func listo() {
	log.Println("Se terminó la cosa")
}

func main() {
	defer listo()

	log.Println("Esto va antes")
}
