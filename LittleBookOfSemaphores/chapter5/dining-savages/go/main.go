package main

import (
	"math/rand"
	"sync"
	"time"
)

func main() {
	m := 5  // serving meals
	s := 10 // savages

	ready := make(chan bool, 1)
	cooker := &Cooker{M: m, Start: ready}
	savages := make([]*Savage, s)
	lock := &sync.Mutex{}
	for i := range savages {
		savages[i] = &Savage{
			Id:         i + 1,
			MaxSleepMs: 1000,
			Lock:       lock,
		}
	}

	servings := int32(0) // initially zero servings
	emptyPot := &sync.WaitGroup{}
	fullPot := &sync.WaitGroup{}

	go cooker.Run(emptyPot, fullPot, &servings)

	<-ready
	for _, savage := range savages {
		go savage.Run(emptyPot, fullPot, &servings)
	}

	time.Sleep(time.Duration(15) * time.Second)
}

func RandomSleep(maxMs int) {
	sleep := time.Duration(rand.Intn(maxMs)+1) * time.Millisecond
	time.Sleep(sleep)
}