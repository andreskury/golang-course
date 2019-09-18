package main

import "fmt"

func foo(name string) string {
	return fmt.Sprintf("Hola %s", name)
}

func main() {
	fmt.Println(foo("Jonathan"))
}
