package main

import (
	"sync"
)

type Semaphore struct {
	count int
	lock  *sync.Mutex
	chs   *Queue
}

func NewSemaphore(count int) *Semaphore {
	return &Semaphore{count: count, lock: &sync.Mutex{}, chs: &Queue{}}
}

func (s *Semaphore) Wait() {
	s.lock.Lock()
	s.count--
	if s.count < 0 {
		ch := make(chan bool, 1)
		s.chs.Push(&ch)
		s.lock.Unlock()
		<-ch
		return
	}
	s.lock.Unlock()
}

func (s *Semaphore) Signal() {
	s.lock.Lock()
	s.count++
	ch, ok := s.chs.ExtractFront()
	s.lock.Unlock()
	if ok {
		*ch <- true
	}
}
