package main

import (
	"encoding/json"
	"os"
)

type User struct {
	Id        string `json:"id"`
	Login     string `json:"login"`
	AvatarUrl string `json:"avatar_url"`
}

func main() {
	u := User{Id: "1234", Login: "xetorthio", AvatarUrl: "foo.bar"}
	json.NewEncoder(os.Stdout).Encode(u)
}
