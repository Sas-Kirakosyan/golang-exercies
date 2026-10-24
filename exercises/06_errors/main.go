// Exercise 6: Collect errors from goroutines
//
// Run: go run -race ./exercises/06_errors
//
// A goroutine CANNOT return a value to the caller:
//
//	go check(x)   // the returned error is thrown away!
//
// So we save errors into a shared slice (protected by a mutex, Exercise 3)
// and return them all at the end with errors.Join.
//
//	errors.Join(nil...)      -> nil   (no errors = success)
//	errors.Join(e1, e2, e3)  -> one error that contains all three
//
// TASK: Finish CheckAll so it:
//  1. runs check(n) for every number in its own goroutine
//  2. waits for all of them (WaitGroup)
//  3. collects every error into errs (Mutex!)
//  4. returns errors.Join(errs...)
//
// Expected output: errors for 3, 7 and 9 (any order).
//
// FINAL STEP: after this, open workerpool/main.go. It is Exercises
// 2 + 3 + 5 + 6 put together. You should now understand every line.
package main

import (
	"errors"
	"fmt"
	"sync"
)

func check(n int) error {
	if n == 3 || n == 7 || n == 9 {
		return fmt.Errorf("number %d is bad", n)
	}
	return nil
}

func CheckAll(numbers []int) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errs []error
	_, _ = &wg, &mu // remove this line once you use wg and mu

	for _, n := range numbers {
		// TODO: start a goroutine that calls check(n)
		//       and appends the error to errs if it is not nil
		_ = n
	}

	// TODO: wait

	return errors.Join(errs...)
}

func main() {
	err := CheckAll([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if err != nil {
		fmt.Println("errors:\n" + err.Error())
		return
	}
	fmt.Println("no errors")
}
