package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	storage := make(map[int]int, 5)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			storage[i] = i * 2
			wg.Done()
		}(i)

	}
	wg.Wait()
	fmt.Println(storage)
}
