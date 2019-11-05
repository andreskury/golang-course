package main

import "net/http"

type handler struct {
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func main() {
	h := handler{}
	http.ListenAndServe(":8080", h)
}
