package main

import (
	"errors"
	"log"
)

func Algo() (string, error) {
	return "", errors.New("foobar")
}

func main() {
	s, err := Algo()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s)
}
