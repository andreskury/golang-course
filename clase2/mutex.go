package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	for r := 0; r < 100; r++ {
		go func() {
			for {

				key := rand.Intn(5)
				mutex.Lock()
				state[key] += 1
				mutex.Unlock()

				time.Sleep(time.Millisecond * 100)
			}
		}()
	}

	time.Sleep(time.Second)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
