package main

import "fmt"

func add_mul(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	a, m := add_mul(3, 3)
	fmt.Printf("Addition: %d - Multiplication: %d\n", a, m)
}
