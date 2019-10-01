package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	req, _ := http.NewRequest("GET", "https://api.github.com/users/xetorthio", nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, res.Body)
}
