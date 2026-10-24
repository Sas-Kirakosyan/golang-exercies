// Exercise 5: Limit how many goroutines run at the same time
//
// Run: go run ./exercises/05_limit
//
// Right now all 20 jobs run at once ("running now" goes up to 20).
//
// TRICK: a buffered channel with capacity N can hold only N items.
//
//	limit := make(chan struct{}, 3)
//	limit <- struct{}{}   // take a slot  (BLOCKS if 3 slots are already taken)
//	<-limit               // free a slot
//
// struct{}{} is just an empty value that uses 0 bytes; we only care about
// the COUNT of items in the channel, not the values.
//
// TASK: Use `limit` so that "running now" never goes above 3.
//  1. Take a slot in the loop, BEFORE `go func()`.
//  2. Free the slot inside the goroutine when the job is finished (defer).
//
// QUESTION: why do we take the slot in the loop and not inside the goroutine?
//
//	(Hint: if we take it inside, how many goroutines exist at once?)
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var running atomic.Int32 // safe counter, works like a counter + mutex

func job(id int) {
	now := running.Add(1)
	fmt.Printf("job %2d started, running now: %d\n", id, now)
	time.Sleep(200 * time.Millisecond)
	running.Add(-1)
}

func main() {
	var wg sync.WaitGroup
	limit := make(chan struct{}, 3)
	_ = limit // remove this line once you use limit

	for i := range 20 {
		wg.Add(1)
		// TODO: take a slot

		go func() {
			defer wg.Done()
			// TODO: free the slot when finished (defer)
			job(i)
		}()
	}

	wg.Wait()
	fmt.Println("all jobs done")
}
