package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	page := 0
	sources := []int{1, 2, 3, 4, 5}
	response := make([]int, 5)
	// semPattern := make(chan int, 5)
	round := 10
	for i := 0; i < round; i++ {
		for index, data := range sources {
			go func(index int, data int) {
				response[index] = doSomethings(index, data)
				page++
				// semPattern <- response[index]
			}(index, data)

		}
		fmt.Println("Round", i, "resul", response[0])
		time.Sleep(2 * time.Second)
	}
}

func doSomethings(index, data int) int {
	log.Printf("index: %d data: %d", index, data)
	return index + data
}
