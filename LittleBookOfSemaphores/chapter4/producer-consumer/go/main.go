package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	count := 0
	genEvent := func() *Event {
		sleep := time.Duration(rand.Intn(999) + 1) * time.Millisecond
		time.Sleep(sleep)
		count++
		fmt.Printf("Producing event %d\n", count)
		return &Event{Id: count, Delay: sleep}
	}
	bufferSize := 5
	buffer := make(chan *Event, bufferSize) // The secret is here, it ensures all synchronization constraints of the problem.

	go Producer(buffer, genEvent)
	go Consumer(buffer)

	time.Sleep(time.Duration(15) * time.Second)
}

type Event struct {
	Id    int
	Delay time.Duration
}

func (e *Event) Process() {
	time.Sleep(2 * e.Delay) // Event consumption is 2x slower than production.
	fmt.Printf("Processing event %d\n", e.Id)
}

func Producer(buffer chan *Event, waitForEvent func() *Event) {
	for {
		event := waitForEvent()
		buffer <- event
	}
}

func Consumer(buffer chan *Event) {
	for {
		event := <-buffer
		event.Process()
	}
}
