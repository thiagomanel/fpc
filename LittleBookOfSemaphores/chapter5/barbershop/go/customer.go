package main

import (
	"fmt"
	"sync"
)

type Customer struct {
	Id    int
	Lock  *sync.Mutex
	Ready *sync.WaitGroup
}

func (c *Customer) BarberArrival(sits chan *Customer, cap int) {
	c.Lock.Lock()
	if len(sits) < cap {
		fmt.Printf("Customer (%d) will wait in the chairs\n", c.Id)
		c.Ready = &sync.WaitGroup{}
		c.Ready.Add(1)
		sits <- c
		c.Lock.Unlock()

		c.Ready.Wait()
		c.getHairCut()
	} else {
		c.Lock.Unlock()
		c.balk()
	}
}

func (c *Customer) balk() {
	fmt.Printf("Customer (%d) balked the wait for the hair cut\n", c.Id)
}

func (c *Customer) getHairCut() {
	fmt.Printf("Customer (%d) is getting a hair cut\n", c.Id)
}