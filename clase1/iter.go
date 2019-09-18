package main

import "fmt"

func main() {
	a := []int{10, 50, 80}

	for i, v := range a {
		fmt.Printf("[%d]: %d\n", i, v)
	}

	m := map[string]int{"foo": 10, "bar": 50}

	for k, v := range m {
		fmt.Printf("%s: %d\n", k, v)
	}
}
