package main

import (
	"math/rand"
	"sync"
	"time"
)

func main() {
	n := 3
	sits := make(chan *Customer, n)

	barber := &Barber{MaxSleepMs: 10}
	go barber.Work(sits)

	lock := &sync.Mutex{}
	for i := 0; i < 20; i++ {
		customer := &Customer{
			Id:   i + 1,
			Lock: lock,
		}
		go customer.BarberArrival(sits, n)
		time.Sleep(time.Duration(2) * time.Millisecond)
	}
	time.Sleep(time.Duration(5) * time.Second)
}

func RandomSleep(maxMs int) {
	sleep := time.Duration(rand.Intn(maxMs)+1) * time.Millisecond
	time.Sleep(sleep)
}