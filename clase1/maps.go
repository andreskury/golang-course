package main

import "fmt"

func main() {
	a := make(map[string]string)
	a["foo"] = "a"
	a["bar"] = "b"
	fmt.Printf("%#v\n", a)

	b := map[string]int{"a": 1, "b": 2}
	fmt.Printf("%#v\n", b)

	fmt.Println(len(a))
	delete(a, "foo")
	fmt.Printf("%#v\n", a)
}
