// Exercise 3: Data race and sync.Mutex
//
// Run: go run ./exercises/03_mutex
// Run with race detector: go run -race ./exercises/03_mutex
//
// STEP 1: Run it a few times. Expected 1000, but you often get less.
//
//	Why? -> `counter++` is really 3 steps: read, add 1, write.
//	Two goroutines can read the same value at the same time, and
//	one of the +1 is lost. This is a DATA RACE.
//
// STEP 2: Run with -race. Go tells you exactly where the race is.
//
// STEP 3: Fix it with a sync.Mutex:
//
//	   mu.Lock()
//	   counter++
//	   mu.Unlock()
//	Only ONE goroutine can be between Lock and Unlock at a time.
//	Run again with -race: no warning, always 1000.
//
// This is exactly why the worker pool code has `mu` around `errs = append(...)`.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	counter := 0

	for range 1000 {

		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++ // TODO: protect with mu
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("counter =", counter, "(expected 1000)")
}
