package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ps := make([]*Philosopher, 5)
	amount := len(ps)
	for i := range ps {
		ps[i] = &Philosopher{Id: i, Amount: amount, MaxSleepMs: 10, Counter: 0}
	}
	turnstile := NewTurnstile(amount - 1)
	forks := NewForks(amount + 1)

	quit := make(chan bool, amount)
	done := &sync.WaitGroup{}
	done.Add(amount)
	for _, p := range ps {
		go p.Run(forks, turnstile, quit, done)
	}

	time.Sleep(time.Duration(15) * time.Second)
	for range ps {
		quit <- true
	}
	done.Wait()
	for _, p := range ps {
		fmt.Printf("Philosopher (%d) ate %d times\n", p.Id, p.Counter)
	}
}

func RandomSleep(maxMs int) {
	sleep := time.Duration(rand.Intn(maxMs)+1) * time.Millisecond
	time.Sleep(sleep)
}
