// Interview task: improve this code.
//
//	func ImportGames(games []Game) error {
//		for _, game := range games {
//			go func() {
//				syncGame(game)
//			}()
//		}
//		return nil
//	}
//
// Problems:
//  1. Doesn't wait for goroutines      -> fix: sync.WaitGroup
//  2. Unlimited goroutines             -> fix: buffered channel as a limit
//  3. Errors lost, always returns nil  -> fix: collect errors and return them
//
// Run: go run ./workerpool
package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Game struct {
	ID int
}

func ImportGames(games []Game) error {
	var wg sync.WaitGroup            // 1. to wait for all goroutines
	limit := make(chan struct{}, 10) // 2. max 10 goroutines running at the same time

	var mu sync.Mutex // 3. protects errs (many goroutines write to it)
	var errs []error

	for _, game := range games {
		wg.Add(1)
		limit <- struct{}{} // take a slot; BLOCKS if 10 are already running

		go func() {
			defer wg.Done()
			defer func() { <-limit }() // free the slot when finished

			if err := syncGame(game); err != nil {
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()                   // wait until ALL games are synced
	return errors.Join(errs...) // nil if no errors
}

func syncGame(game Game) error {
	time.Sleep(10 * time.Millisecond) // pretend this is an HTTP/DB call
	if game.ID%25 == 0 {
		return fmt.Errorf("game %d: sync failed", game.ID)
	}
	return nil
}

func main() {
	games := make([]Game, 100)
	for i := range games {
		games[i] = Game{ID: i}
	}

	if err := ImportGames(games); err != nil {
		fmt.Println("errors:\n" + err.Error())
		return
	}
	fmt.Println("all games synced")
}
