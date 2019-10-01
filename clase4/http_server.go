package main

import "net/http"

func createUser(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/users", createUser)
}
