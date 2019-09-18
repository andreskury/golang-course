package main

import "fmt"

type Dog struct {
	Name string
}

func (d Dog) Talk() {
	fmt.Printf("%s: wof wof!\n", d.Name)
}

func main() {
	ari := Dog{Name: "Ari"}
	ari.Talk()
}
