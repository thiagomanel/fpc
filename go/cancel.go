package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func msg(id int, done chan interface{}) {
	for {
		select {
		case <-done:
			fmt.Println("I'm going to die")
			return
		default:
			fmt.Printf("default <%d>\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func killer(done chan interface{}) {
	//how to kill the msg goroutines?

	//read from input stream
	os.Stdin.Read(make([]byte, 1))
	close(done)
}

func main() {
	done := make(chan interface{})

	go killer(done)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var ranN = r.Intn(35)
	fmt.Printf("%d gourotines\n", ranN)

	for i := 0; i < ranN; i++ {
		go msg(i, done)
	}

	time.Sleep(1 * time.Minute)
}
