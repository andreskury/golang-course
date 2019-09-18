package main

import "fmt"

type Circle struct {
	X, Y   int
	Radius float32
}

func main() {
	c := Circle{X: 10, Y: 10}
	c.Radius = 3.4

	fmt.Printf("%#v\n", c)
}
