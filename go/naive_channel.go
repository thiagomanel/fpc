package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	join_ch := make(chan int)

	//below goroutines communicate through the ch channel
	go prod(ch, join_ch)
	go cons(ch)

	//using a channel to join
	<-join_ch
}

func prod(my_chan chan int, j_ch chan int) {
	var n int = 0
	for {
		fmt.Printf("going to produce (%d) \n", n)
		my_chan <- n
		n = n + 1
		//produce up to 100 items, then signal to the join channel it is done
		if n == 100 {
			j_ch <- n
		}
	}
}

func cons(my_chan chan int) {
	for {
		//consume the items produced by the prod goroutine
		y := <-my_chan
		fmt.Printf("consumed (%d) \n", y)
	}
}
