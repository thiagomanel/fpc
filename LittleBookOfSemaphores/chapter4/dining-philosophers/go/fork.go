package main

import (
	"fmt"
)

type Forks struct {
	forks map[int]*Fork
}

type Fork struct {
	id int
	ch chan bool
}

func NewForks(amount int) *Forks {
	fs := &Forks{forks: map[int]*Fork{}}
	for i := 0; i < amount; i++ {
		fid := i
		ch := make(chan bool, 1)
		ch <- true
		fs.forks[fid] = &Fork{id: fid, ch: ch}
	}
	return fs
}

func (fs *Forks) GetForks(ids []*int) ([]*Fork, error) {
	forks := make([]*Fork, len(ids))
	for i, id := range ids {
		f, ok := fs.forks[*id]
		if !ok {
			return nil, fmt.Errorf("unable to find fork with id (%d)", *id)
		}
		<-f.ch
		forks[i] = f
	}
	return forks, nil
}

func (fs *Forks) ReleaseForks(forks []*Fork) {
	for _, fork := range forks {
		fork.ch <- true
	}
}
