// Exercise 1: What is a goroutine?
//
// Run: go run ./exercises/01_goroutine
//
// STEP 1: Run it. You will see "main finished" but almost no "hello" lines.
//
//	Why? -> `go` starts the function in the background and main does
//	NOT wait. When main returns, the whole program exits and kills
//	all goroutines.
//
// STEP 2: Add `time.Sleep(time.Second)` before the last Println and run again.
//
//	Now you see all 5 "hello" lines. Is the order always 0,1,2,3,4?
//	Run it several times and check.
//
// STEP 3: Think: why is time.Sleep a BAD way to wait?
//
//	(Hint: what if one goroutine takes 2 seconds?)
//	-> Exercise 2 shows the right way.
package main

import (
	"fmt"
	"sync"
)

func sayHello(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello from goroutine", i)

}

func main() {
	wg := sync.WaitGroup{}
	for i := range 5 {
		wg.Add(1)
		go sayHello(i, &wg)

	}

	wg.Wait()

	// time.Sleep(1 * time.Second)
	// TODO (step 2): add time.Sleep here (and import "time")

	fmt.Println("main finished")
}
