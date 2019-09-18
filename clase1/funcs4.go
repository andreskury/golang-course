package main

import "fmt"

func add_mul2(a, b int) (addition int, multiplication int) {
	addition = a + b
	multiplication = a * b
	return
}

func main() {
	a, m := add_mul2(3, 3)
	fmt.Printf("Addition: %d - Multiplication: %d\n", a, m)
}
