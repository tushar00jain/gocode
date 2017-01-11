
package main

import (
	"fmt"
	"time"
)

func producer(ch chan int, d time.Duration, i int) {
	for {
		ch <- i
	  time.Sleep(d)
	}
}

func reader(out chan int) {
	for {
		fmt.Println(<-out)
	}
}

func main() {
	ch, out := make(chan int), make(chan int)
	go producer(ch, time.Millisecond * 100, 1)
	go producer(ch, time.Millisecond * 300, 2)
	go reader(out)
	for {
		out <- <-ch
	}
}
