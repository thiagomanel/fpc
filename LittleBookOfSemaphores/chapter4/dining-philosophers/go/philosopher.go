package main

import (
	"fmt"
	"sync"
)

type Philosopher struct {
	Id         int
	Amount     int
	MaxSleepMs int
	Counter    int
}

func (p *Philosopher) Run(forks *Forks, turnstile *Turnstile, quit chan bool, done *sync.WaitGroup) {
	for {
		select {
		case <-quit:
			done.Done()
			return
		default:
			p.dinner(forks, turnstile)
		}
	}
}

func (p *Philosopher) dinner(forks *Forks, turnstile *Turnstile) {
	turnstile.Pass(p.Id)

	p.think()
	fs, err := p.getForks(forks)
	if err != nil {
		panic(err)
	}
	p.eat()
	p.Counter++
	p.releaseForks(forks, fs)

	turnstile.Return()
}

func (p *Philosopher) think() {
	fmt.Printf("Philosopher (%d) is thinking\n", p.Id)
	RandomSleep(p.MaxSleepMs)
}

func (p *Philosopher) getForks(forks *Forks) ([]*Fork, error) {
	fstForkId := p.Id
	secForkId := (p.Id + 1) % p.Amount
	fmt.Printf("Philosopher (%d) is getting the forks\n", p.Id)
	fs, err := forks.GetForks([]*int{&fstForkId, &secForkId})
	fmt.Printf("Philosopher (%d) got the forks\n", p.Id)
	return fs, err
}

func (p *Philosopher) eat() {
	fmt.Printf("Philosopher (%d) is eating\n", p.Id)
	RandomSleep(p.MaxSleepMs)
}

func (p *Philosopher) releaseForks(forks *Forks, fs []*Fork) {
	fmt.Printf("Philosopher (%d) is releasing the forks\n", p.Id)
	forks.ReleaseForks(fs)
	fmt.Printf("Philosopher (%d) released the forks\n", p.Id)
}
