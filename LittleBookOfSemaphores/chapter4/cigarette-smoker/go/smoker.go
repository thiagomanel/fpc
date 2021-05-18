package main

import (
	"fmt"
)

func SmokerA(matcherSem *Semaphore, agent *Semaphore) {
	id := "SmokerA"
	for {
		matcherSem.Wait()
		makeCigarette(id)
		agent.Signal()
		smoke(id)
	}
}

func SmokerB(paperSem *Semaphore, agent *Semaphore) {
	id := "SmokerB"
	for {
		paperSem.Wait()
		makeCigarette(id)
		agent.Signal()
		smoke(id)
	}
}

func SmokerC(tobaccoSem *Semaphore, agent *Semaphore) {
	id := "SmokerC"
	for {
		tobaccoSem.Wait()
		makeCigarette(id)
		agent.Signal()
		smoke(id)
	}
}

func makeCigarette(id string) {
	fmt.Printf("Smoker (%s) is making a cigarette\n", id)
}

func smoke(id string) {
	fmt.Printf("Smoker (%s) is making a smoking\n", id)
}
