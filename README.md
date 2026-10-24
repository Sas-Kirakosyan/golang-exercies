# Go Concurrency Practice

My notes and exercises for learning Go concurrency. It started from a real interview question: improve a function that starts one goroutine per item.

## Project structure

```
go_todo/
├── main.go                # slice / append experiments
├── workerpool/main.go     # final solution to the interview task
└── exercises/
    ├── 01_goroutine/      # what `go` does, main doesn't wait
    ├── 02_waitgroup/      # waiting for goroutines
    ├── 03_mutex/          # data races and sync.Mutex
    ├── 04_channels/       # send / receive, buffered vs unbuffered
    ├── 05_limit/          # buffered channel as a concurrency limit
    └── 06_errors/         # collecting errors from goroutines
```

Run any program from the project root:

```bash
go run ./workerpool
go run ./exercises/01_goroutine
go run -race ./exercises/03_mutex   # with the race detector
```

---

## The interview task

```go
func ImportGames(games []Game) error {
	for _, game := range games {
		go func() {
			syncGame(game)
		}()
	}
	return nil
}
```

### What's wrong

| # | Problem | Fix |
|---|---|---|
| 1 | Doesn't wait: returns before the syncs finish | `sync.WaitGroup` |
| 2 | One goroutine per game: 1M games means 1M syncs at once | Buffered channel as a limit |
| 3 | Errors are lost and it always returns `nil` | Collect errors and return `errors.Join(errs...)` |
| 4 | Before Go 1.22, every closure shared the same `game` variable | Fixed in Go 1.22+ (each iteration gets its own copy) |
| 5 | No `context`: can't cancel or time out | Pass `ctx context.Context` (nice to have) |
| 6 | A panic in one goroutine crashes the whole program | `defer recover()` inside the goroutine (nice to have) |

### Simple solution: WaitGroup + buffered channel

```go
func ImportGames(games []Game) error {
	var wg sync.WaitGroup            // 1. wait for all goroutines
	limit := make(chan struct{}, 10) // 2. max 10 running at the same time

	var mu sync.Mutex // 3. protects errs
	var errs []error

	for _, game := range games {
		wg.Add(1)
		limit <- struct{}{} // take a slot; blocks if 10 are already running

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

	wg.Wait()
	return errors.Join(errs...) // nil if there were no errors
}
```

**Interview answer in one line:** "The code is missing three things: waiting, a limit, and error handling. I fix them with a WaitGroup, a buffered channel, and collecting the errors."

### Alternative: `errgroup`

```go
g, ctx := errgroup.WithContext(ctx)
g.SetLimit(10)
for _, game := range games {
	g.Go(func() error { return syncGame(ctx, game) })
}
return g.Wait()
```

The difference: `errgroup` cancels the rest at the **first** error. The WaitGroup version keeps going and reports **all** errors, which usually suits an import better.

---

## Key concepts

### I/O-bound work still needs a limit

Downloading or syncing is I/O-bound: goroutines spend their time waiting, and they're cheap (a few KB each). But each one holds real resources: a socket (file descriptor), memory buffers, a DB connection, and load on the remote server. Unlimited concurrency fails with `too many open files`, OOM, or rate limits.

- **CPU-bound:** limit ≈ `runtime.GOMAXPROCS(0)`
- **I/O-bound:** higher limit (tens to hundreds), tuned by measuring, but **still limited**

### `runtime.NumCPU()` vs `runtime.GOMAXPROCS()`

| | `NumCPU()` | `GOMAXPROCS(n)` |
|---|---|---|
| What | Logical CPUs the process can use | Max OS threads running Go code at the same moment |
| Changeable | No, read once at startup | Yes. `GOMAXPROCS(0)` only reads it |
| Default | From the machine | Same as `NumCPU()` |

- `GOMAXPROCS` is a function and an environment variable, not a constant.
- It doesn't limit the number of goroutines, only how many run in parallel.
- Since Go 1.25, `GOMAXPROCS` follows container CPU limits automatically.

### `atomic` vs `sync.Mutex`

- **`atomic`**: a single counter or flag. Faster and simple (`counter.Add(1)`).
- **`Mutex`**: anything else (a slice, a map, or several variables that must change together).
- Not sure? Use a Mutex.

### WaitGroup rules

```go
var wg sync.WaitGroup // zero value is ready; `var` is the usual style

wg.Add(1)            // BEFORE `go`
go func() {
	defer wg.Done()  // FIRST line, with `defer`, so it runs AFTER the work
	work()
}()
wg.Wait()
```

Mistakes I made:
- Calling `wg.Done()` in `main` right after `go`: the counter drops to 0 immediately, so `Wait()` doesn't wait.
- Calling `wg.Done()` without `defer` before the work: same problem.
- Passing `sync.WaitGroup` by value instead of `*sync.WaitGroup`: `Done()` changes a copy, causing a deadlock. `go vet` catches this.

Prefer the closure style (`go func() { defer wg.Done(); work(i) }()`). It keeps `work` a normal function.

### `var` vs `:=`

- `var x T`: zero value (`var wg sync.WaitGroup`, `var mu sync.Mutex`).
- `x := value`: a real starting value (`ch := make(chan int)`).
- `:=` only works inside functions.

### `range` forms

```go
for i, v := range numbers {} // index and value
for i := range numbers {}    // index only
for _, v := range numbers {} // value only
for range numbers {}         // neither, just repeat (NOT `for _ := range`)
```

### Channels

- **Unbuffered** (`make(chan int)`): a send waits until **another goroutine** receives.
- **Buffered** (`make(chan int, 3)`): a send waits only when the buffer is full, and a receive waits only when it's empty.

Why this deadlocks:

```go
ch := make(chan int)
ch <- 1           // main waits here for a receiver...
fmt.Println(<-ch) // ...but this line, also in main, is never reached
```

Both lines run one after another in the same goroutine. Nobody else can receive, so Go stops with `all goroutines are asleep - deadlock!`. The fix is to send from a different goroutine, or to use a buffered channel.

### Race detector

```bash
go run -race ./exercises/03_mutex
go test -race ./...
```

It reports when two goroutines access the same variable at the same time and at least one writes. It only finds races that actually happen during that run, so run it several times.

---

## Progress

| # | Exercise | Status |
|---|---|---|
| 1 | Goroutines | ✅ Done |
| 2 | WaitGroup | ✅ Done |
| 3 | Mutex | ✅ Done |
| 4 | Channels | ✅ Done |
| 5 | Limit (buffered channel) | ⬜ Next |
| 6 | Collecting errors | ⬜ |

After exercise 6: re-read [workerpool/main.go](workerpool/main.go). It combines exercises 2, 3, 5 and 6.

### Next topics

1. `context`: cancellation and timeouts
2. `select`: waiting on several channels
3. Worker pool with a fixed number of workers and a jobs channel
4. `errgroup` from `golang.org/x/sync`

Extra practice: [Tour of Go: Concurrency](https://go.dev/tour/concurrency/1)
