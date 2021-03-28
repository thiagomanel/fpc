package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	threads, n := 10, 5
	ch := make(chan bool, n)
	end := &sync.WaitGroup{}
	end.Add(threads)
	count := int32(0)
	for i := 0; i < n; i++ {
		ch <- true
	}
	for i := 0; i < threads; i++ {
		go thread(&count, ch, end)
	}
	end.Wait()
	fmt.Printf("Count: %d\n", atomic.LoadInt32(&count))
}

func thread(count *int32, ch chan bool, end *sync.WaitGroup) {
	for i := 0; i < 100000; i++ {
		<- ch
		*count = *count + 1
		ch <- true
	}
	end.Done()
}
