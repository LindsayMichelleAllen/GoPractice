package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup
var n = 25

func calculate(channel chan int, start int) {
	defer wg.Done()
	channel <- start + 5
}

func main() {
	wg.Add(n)
	intChannel := make(chan int, n)
	for i := 0; i <= n; i++ {
		go calculate(intChannel, rand.Intn(100))
	}
	wg.Wait()
	wg.Add(1)
	go func() {
		close(intChannel)
		for i := range intChannel {
			fmt.Println(i)
		}
		wg.Done()
	}()
	wg.Wait()
}
