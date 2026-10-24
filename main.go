package main

import "fmt"

// func compute(ch chan int, num int) {
// 	result := num * 2
// 	ch <- result
// }

func f(sl []string) []string {
	sl = append(sl, "hello")
	return sl

}

func main() {
	sl := make([]string, 0, 5)
	f(sl)
	fmt.Println("s", f(sl))
	// sl = append(sl, 5)
	// fmt.Println("s", sl[1:3])

	// ch := make(chan int)

	// go compute(ch, 2)
	// go compute(ch, 5)

	// result1 := <-ch
	// result2 := <-ch

	// fmt.Println("Results:", result1, "and", result2)
}
