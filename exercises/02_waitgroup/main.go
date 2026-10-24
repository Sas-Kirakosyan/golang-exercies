// Exercise 2: Waiting for goroutines with sync.WaitGroup
//
// Run: go run ./exercises/02_waitgroup
//
// A WaitGroup is just a counter:
//
//	wg.Add(1)  -> counter + 1   (call BEFORE starting the goroutine)
//	wg.Done()  -> counter - 1   (call when the goroutine finishes)
//	wg.Wait()  -> block until counter == 0
//
// TASK: Use a WaitGroup so main prints "all done" only AFTER all 5 workers
//
//	have printed "finished". No time.Sleep in main!
//
// BONUS: What happens if you forget wg.Done()? Try it. (You will see a
//
//	"deadlock" error. Understand why.)
package main

import (
	"fmt"
	"sync"
	"time"
)

func work(id int) {
	fmt.Println("worker", id, "started")
	time.Sleep(time.Duration(id*100) * time.Millisecond) // pretend to work
	fmt.Println("worker", id, "finished")
}

func main() {
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work(i)
		}()
	}

	wg.Wait()

	fmt.Println("all done")
}
