package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	n := 5
	barrier := &sync.WaitGroup{}
	barrier.Add(n)
	end := &sync.WaitGroup{}
	end.Add(n)
	for i := 0; i < n; i++ {
		go thread(i, barrier, end)
	}
	end.Wait()
}

func thread(id int, barrier, end *sync.WaitGroup) {
	rendezvous(id)
	barrier.Done()
	barrier.Wait()
	// Critical Section
	end.Done()
}

func rendezvous(id int) {
	sleep := rand.Intn(5) + 1
	fmt.Printf("Thread (%d): sleeping during (%d) seconds\n", id, sleep)
	time.Sleep(time.Duration(sleep) * time.Second)
}

