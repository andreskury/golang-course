package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	r.GET("/foo", FooHandler)
	r.GET("/bar", BarHandler)

	http.ListenAndServe(":8080", r)
}

func FooHandler(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(rw, "foo")
}
func BarHandler(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(rw, "bar")
}
