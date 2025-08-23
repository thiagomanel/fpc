package main

import (
	"fmt"
)

func AgentA(agent, tobacco, paper *Semaphore) {
	for {
		agent.Wait()
		fmt.Printf("Agent A is producing Tobacco and Paper\n")
		tobacco.Signal()
		paper.Signal()
	}
}

func AgentB(agent, paper, matcher *Semaphore) {
	for {
		agent.Wait()
		fmt.Printf("Agent B is producing Paper and Matcher\n")
		paper.Signal()
		matcher.Signal()
	}
}

func AgentC(agent, tobacco, matcher *Semaphore) {
	for {
		agent.Wait()
		fmt.Printf("Agent C is producing Tobacco and Matcher\n")
		tobacco.Signal()
		matcher.Signal()
	}
}
