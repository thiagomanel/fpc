package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	sem := make(chan bool, 1)
	sem <- true
	end := &sync.WaitGroup{}
	end.Add(2)
	count := int32(0)
	go threadA(&count, sem, end)
	go threadB(&count, sem, end)
	end.Wait()
	fmt.Printf("Count: %d\n", atomic.LoadInt32(&count))
}

func threadA(count *int32, sem chan bool, end *sync.WaitGroup) {
	for i := 0; i < 1000000; i++ {
		<- sem
		atomic.AddInt32(count, 1)
		sem <- true
	}
	end.Done()
}

func threadB(count *int32, sem chan bool, end *sync.WaitGroup) {
	for i := 0; i < 1000000; i++ {
		<- sem
		atomic.AddInt32(count, 1)
		sem <- true
	}
	end.Done()
}
