package main

import (
	"fmt"
	"log"
)

type PasoAlgoMalo int

func (m PasoAlgoMalo) Error() string {
	return fmt.Sprintf("Pasó algo muuuuuuy malo. Y el código de error es: %d", m)
}

func Algo() (string, error) {
	return "", PasoAlgoMalo(10)
}

func main() {
	s, err := Algo()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s)
}
