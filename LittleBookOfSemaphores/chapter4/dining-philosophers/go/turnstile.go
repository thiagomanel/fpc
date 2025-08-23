package main

import "fmt"

type Turnstile struct {
	tableSits int
	sitsCh    chan bool
}

func NewTurnstile(sits int) *Turnstile {
	ch := make(chan bool, sits)
	for i := 0; i < sits; i++ {
		ch <- true
	}
	return &Turnstile{tableSits: sits, sitsCh: ch}
}

func (t *Turnstile) Pass(fid int) {
	fmt.Printf("Philosopher (%d) going through the turnstile\n", fid)
	<-t.sitsCh
	fmt.Printf("Philosopher (%d) released from the turnstile\n", fid)
}

func (t *Turnstile) Return() {
	t.sitsCh <- true
}
