package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	mutex := &sync.RWMutex{}
	end := &sync.WaitGroup{}
	end.Add(2)
	count := int32(0)
	go threadA(&count, mutex, end)
	go threadB(&count, mutex, end)
	end.Wait()
	fmt.Printf("Count: %d\n", atomic.LoadInt32(&count))
}

func threadA(count *int32, mutex *sync.RWMutex, end *sync.WaitGroup) {
	for i := 0; i < 1000000; i++ {
		mutex.Lock()
		*count = *count + 1
		mutex.Unlock()
	}
	end.Done()
}

func threadB(count *int32, mutex *sync.RWMutex, end *sync.WaitGroup) {
	for i := 0; i < 1000000; i++ {
		mutex.Lock()
		*count = *count + 1
		mutex.Unlock()
	}
	end.Done()
}
