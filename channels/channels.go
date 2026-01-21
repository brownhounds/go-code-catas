package channels

func Run() []int {
	nums := make(chan int)
	results := make(chan int)
	squares := []int{}

	go func() {
		for i := 1; i <= 100; i++ {
			nums <- i
		}

		close(nums)
	}()

	go func() {
		for v := range nums {
			results <- v * v
		}
		close(results)
	}()

	for r := range results {
		squares = append(squares, r)
	}

	return squares
}
