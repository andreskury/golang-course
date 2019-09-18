package main

import "fmt"

func bar(first, last string) string {
	return fmt.Sprintf("Hola %s %s", first, last)
}

func main() {
	fmt.Println(bar("Jonathan", "Leibiusky"))
}
