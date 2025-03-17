package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Savage struct {
	Id         int
	MaxSleepMs int
	Lock       *sync.Mutex
}

func (s *Savage) Run(emptyPot, fullPot *sync.WaitGroup, servings *int32) {
	for {
		s.Lock.Lock()
		if atomic.LoadInt32(servings) == 0 {
			fullPot.Add(1)
			emptyPot.Done()
			fullPot.Wait()
		}
		atomic.AddInt32(servings, -1)
		s.Lock.Unlock()
		s.eat()
	}
}

func (s *Savage) eat() {
	fmt.Printf("Savage (%d) is eating\n", s.Id)
	RandomSleep(s.MaxSleepMs)
	fmt.Printf("Savage (%d) finished eating\n", s.Id)
}
