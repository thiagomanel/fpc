package main

import (
	"fmt"
	"sync"
)

func main() {
	awg := &sync.WaitGroup{}
	bwg := &sync.WaitGroup{}
	awg.Add(1)
	bwg.Add(1)
	end := &sync.WaitGroup{}
	end.Add(2)
	go threadA(awg, bwg, end)
	go threadB(awg, bwg, end)
	end.Wait()
}

func threadA(aArrived, bArrived, end *sync.WaitGroup) {
	//Beginning of Statement A1
	fmt.Printf("Statement A1\n")
	//End of Statement A1
	aArrived.Done()
	bArrived.Wait()
	//Beginning of Statement A2
	fmt.Printf("Statement A2\n")
	//End of Statement A2
	end.Done()
}

func threadB(aArrived, bArrived, end *sync.WaitGroup) {
	//Beginning of Statement B1
	fmt.Printf("Statement B1\n")
	//End of Statement B1
	bArrived.Done()
	aArrived.Wait()
	//Beginning of Statement B2
	fmt.Printf("Statement B2\n")
	//End of Statement B2
	end.Done()
}
