package main

import (
	"fmt"
	"math/rand"
	"time"
)

type leader struct{ id int }
type follower struct{ id int }

func main() {
	lq := make(chan *leader)
	fq := make(chan *follower)

	dance := func(l *leader, f *follower) {
		fmt.Printf("Leader (%d) is dancing with Follower (%d)\n", l.id, f.id)
	}

	triggerLeader := func(id int) {
		leaderArrival(id, lq, fq, dance)
	}
	triggerFollower := func(id int) {
		followerArrival(id, lq, fq, dance)
	}

	go trigger("Leader", triggerLeader)
	go trigger("Follower", triggerFollower)

	time.Sleep(time.Duration(15) * time.Second)
}

func trigger(name string, event func(int)) {
	for i := 0; ; i++ {
		sleep := rand.Intn(999) + 1
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Triggering (%s)\n", name)
		go event(i)
	}
}

func leaderArrival(id int, lq chan *leader, fq chan *follower, dance func(*leader, *follower)) {
	l := &leader{id: id}
	select {
	case f := <-fq:
		dance(l, f)
	default:
		lq <- l
	}
}

func followerArrival(id int, lq chan *leader, fq chan *follower, dance func(*leader, *follower)) {
	f := &follower{id: id}
	select {
	case l := <-lq:
		dance(l, f)
	default:
		fq <- f
	}
}