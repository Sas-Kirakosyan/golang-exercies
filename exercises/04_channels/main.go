// Exercise 4: Channels basics
//
// Run: go run ./exercises/04_channels
//
// A channel is a pipe between goroutines:
//
//	ch := make(chan int)      // unbuffered: send WAITS until someone receives
//	ch := make(chan int, 3)   // buffered: can hold 3 values without a receiver
//	ch <- 5                   // send
//	v := <-ch                 // receive
//
// PART A (predict BEFORE running, then uncomment and run each one):
//
//	A1: will it work or deadlock?
//	A2: will it work or deadlock?
//	A3: will it work or deadlock?
//
// PART B (TASK): Start 3 goroutines. Each one squares its number
// (2, 3, 4) and SENDS the result into `results`. main receives 3 values
// and prints the sum. Expected: 4 + 9 + 16 = 29.
package main

import "fmt"

func main() {
	// ---- PART A ----

	// A1:
	// ch := make(chan int)
	// ch <- 1

	// fmt.Println(<-ch)

	// A2:
	// ch := make(chan int, 1)
	// ch <- 1
	// fmt.Println(<-ch)

	// A3:
	// ch := make(chan int, 1)
	// ch <- 1
	// ch <- 2
	// fmt.Println(<-ch)

	// ---- PART B ----
	numbers := []int{2, 3, 4}
	results := make(chan int)

	for _, n := range numbers {
		go func() {
			results <- n * n
			fmt.Println("send", n*n)
			_ = n
		}()
	}

	sum := 0

	for range numbers {
		sum += <-results
		fmt.Println("resive", sum)
	}

	fmt.Println("sum =", sum, "(expected 29)")
}
