package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Cooker struct {
	M     int
	Start chan bool
}

func (c *Cooker) Run(emptyPot, fullPot *sync.WaitGroup, servings *int32) {
	c.Start <- true
	for {
		emptyPot.Add(1)
		emptyPot.Wait()
		c.serve(servings)
		fullPot.Done()
	}
}

func (c *Cooker) serve(servings *int32) {
	fmt.Printf("Cooker serving (%d) meals\n", c.M)
	atomic.SwapInt32(servings, int32(c.M))
	fmt.Printf("Cooker served (%d) meals\n", c.M)
}
