package main

import "fmt"

type Dog struct {
	Name  string
	angry bool
}

func (d Dog) Talk() {
	if d.angry {
		fmt.Printf("%s: grrrrrrr!\n", d.Name)
	} else {
		fmt.Printf("%s: wof wof!\n", d.Name)
	}
}
func (d *Dog) MakeAngry() {
	d.angry = true
}

func main() {
	ari := Dog{Name: "Ari"}
	ari.Talk()
	ari.MakeAngry()
	ari.Talk()
}
