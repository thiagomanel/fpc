package main

import "fmt"

type Barber struct {
	MaxSleepMs int
}

func (b *Barber) Work(sits chan *Customer) {
	for c := range sits {
		c.Ready.Done()
		b.cutHair(c)
	}
}

func (b *Barber) cutHair(c *Customer) {
	fmt.Printf("Barber is cutting the hair of Customer (%d)\n", c.Id)
	RandomSleep(b.MaxSleepMs)
	fmt.Printf("Barber finished the hair cut of Customer (%d)\n", c.Id)
}
