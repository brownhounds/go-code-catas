package goroutines

import (
	"log"
	"sync"
)

func Run() []int {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	squares := []int{}

	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, v := range nums {
		wg.Add(1)

		go func(v int) {
			defer wg.Done()

			sq := v * v

			mu.Lock()
			squares = append(squares, sq)
			mu.Unlock()
		}(v)
	}

	wg.Wait()
	log.Println(squares)
	return squares
}
