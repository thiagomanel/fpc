package main

import (
	"sync"
)

type Pusher struct {
	lock      *sync.Mutex
	Tobacco   *Semaphore
	Paper     *Semaphore
	Matcher   *Semaphore
	isPaper   bool
	isMatch   bool
	isTobacco bool
}

func NewPusher(tobacco, paper, matcher *Semaphore) *Pusher {
	return &Pusher{
		lock: &sync.Mutex{},
		isTobacco: false,
		isMatch: false,
		isPaper: false,
		Tobacco: tobacco,
		Paper: paper,
		Matcher: matcher,
	}
}

func (p *Pusher) PusherA(matcherSem *Semaphore, paperSem *Semaphore) {
	for {
		p.Tobacco.Wait()
		p.lock.Lock()
		if p.isPaper {
			p.isPaper = false
			matcherSem.Signal()
		} else if p.isMatch {
			p.isMatch = false
			paperSem.Signal()
		} else {
			p.isTobacco = true
		}
		p.lock.Unlock()
	}
}

func (p *Pusher) PusherB(paperSem *Semaphore, tobaccoSem *Semaphore) {
	for {
		p.Matcher.Wait()
		p.lock.Lock()
		if p.isTobacco {
			p.isTobacco = false
			paperSem.Signal()
		} else if p.isPaper {
			p.isPaper = false
			tobaccoSem.Signal()
		} else {
			p.isMatch = true
		}
		p.lock.Unlock()
	}
}

func (p *Pusher) PusherC(tobaccoSem *Semaphore, matcherSem *Semaphore) {
	for {
		p.Paper.Wait()
		p.lock.Lock()
		if p.isMatch {
			p.isMatch = false
			tobaccoSem.Signal()
		} else if p.isTobacco {
			p.isTobacco = false
			matcherSem.Signal()
		} else {
			p.isPaper = true
		}
		p.lock.Unlock()
	}
}
