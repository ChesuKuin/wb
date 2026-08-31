package main

import (
	"fmt"
	"sync"
)

// result — пара "число / его квадрат", чтобы не потерять соответствие
// при конкурентном выполнении (горутины завершаются в произвольном порядке).
type result struct {
	number int
	square int
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	results := make(chan result, len(numbers))

	for _, n := range numbers {
		wg.Add(1)

		// каждое число обрабатывается в своей горутине
		go func(n int) {
			defer wg.Done()
			results <- result{number: n, square: n * n}
		}(n)
	}

	// отдельная горутина закрывает канал, когда все воркеры завершились —
	// это позволяет использовать range по каналу без ручного подсчёта
	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Printf("%d^2 = %d\n", r.number, r.square)
	}
}
