package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func calc(c chan int, temp int) {
	for i := 0; i < 4; i++ {
		c <- temp + i
	}
	wg.Done()
}

func printCalc(c chan int) {
	for i := 0; i < 4; i++ {
		fmt.Println(<-c)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	chann := make(chan int)
	defer close(chann)
	go calc(chann, 5)
	go printCalc(chann)
	wg.Wait()
}
