package main

import (
	"time"
)

func main() {
	agent := NewSemaphore(1)

	tobaccoSem := NewSemaphore(0)
	paperSem := NewSemaphore(0)
	matcherSem := NewSemaphore(0)

	go SmokerA(matcherSem, agent)
	go SmokerB(paperSem, agent)
	go SmokerC(tobaccoSem, agent)

	tobacco := NewSemaphore(0)
	paper := NewSemaphore(0)
	matcher := NewSemaphore(0)

	pusher := NewPusher(tobacco, paper, matcher)

	go pusher.PusherA(matcherSem, paperSem)
	go pusher.PusherB(paperSem, tobaccoSem)
	go pusher.PusherC(tobaccoSem, matcherSem)

	go AgentA(agent, tobacco, paper)
	go AgentB(agent, paper, matcher)
	go AgentC(agent, tobacco, matcher)

	time.Sleep(time.Duration(15) * time.Second)
}





