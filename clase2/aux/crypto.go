package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func MineBlock(block string, difficulty int) (string, int) {
	leftPad := strings.Repeat("0", difficulty)
	p := 0
	for {
		hash := checksum(getPayload(block, p))
		if strings.HasPrefix(hash, leftPad) {
			return hash, p
		}
		p++
	}
}

func getPayload(block string, pivot int) string {
	return fmt.Sprintf("%s|%d", block, pivot)
}

func checksum(payload string) string {
	h := sha256.New()
	h.Write([]byte(payload))
	return fmt.Sprintf("%x", string(h.Sum(nil)))
}

func ValidateBlock(block string, difficulty int, pivot int) bool {
	leftPad := strings.Repeat("0", difficulty)
	hash := checksum(getPayload(block, pivot))
	if strings.HasPrefix(hash, leftPad) {
		return true
	} else {
		return false
	}
}

/*
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	hash, pivot := MineBlock("aaabbbc", 6)
	valid := ValidateBlock("aaabbbc", 6, pivot)
	if valid {
		log.Println("Valid", hash, pivot)
	} else {
		log.Println("Not valid", hash, pivot)
	}
}
*/
