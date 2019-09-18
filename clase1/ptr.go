package main

import "fmt"

func main() {
	var a = 10

	b := &a
	// *b++

	fmt.Println(a, b)
}
