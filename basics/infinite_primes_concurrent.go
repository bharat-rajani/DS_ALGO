// A concurrent prime sieve

package main

import "fmt"

func Generate(src chan<- int) {
	for i := 2; ; i++ {
		// fmt.Printf(">> Gen (%d)\n", i)
		src <- i
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		// fmt.Printf("Filter(%d) received %d\n", prime, i)
		if i%prime != 0 {
			// fmt.Println(i, prime)
			out <- i
		} else {
			// fmt.Printf("Filter(%d) killed %d\n", prime, i)
		}
	}
}

func main() {

	// In practice this concurrent program is slower :)
	// good read: https://stackoverflow.com/questions/51006592/why-do-sequential-loops-run-faster-than-concurrent-methods-in-go

	src := make(chan int, 1)
	go Generate(src)
	for i := 0; ; i++ {
		prime := <-src
		fmt.Println(prime)
		dest := make(chan int, 1)
		// fmt.Println(src, dest, prime)
		go Filter(src, dest, prime)
		src = dest
	}
}
