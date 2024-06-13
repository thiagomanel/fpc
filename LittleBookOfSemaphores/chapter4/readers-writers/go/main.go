package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	db := &Database{}
	lock := &sync.RWMutex{}

	go Reader(lock, db)
	go Reader(lock, db)
	go Reader(lock, db)

	go Writer(lock, db)
	go Writer(lock, db)

	time.Sleep(time.Duration(10) * time.Second)
}

type Object struct {
	Value int
}

type Database struct {
	Obj Object
}

func (d *Database) Read() Object {
	randomSleep() // simulates the latency of the operation.
	return d.Obj
}

func (d *Database) Write(obj Object) {
	randomSleep() // simulates the latency of the operation.
	d.Obj = obj
}

func Reader(lock *sync.RWMutex, db *Database) {
	for {
		lock.RLock()
		obj := db.Read()
		fmt.Printf("Reading: %v\n", obj)
		lock.RUnlock()
	}
}

func Writer(lock *sync.RWMutex, db *Database) {
	for {
		lock.Lock()
		obj := Object{Value: rand.Intn(10000)}
		fmt.Printf("Writing: %v\n", obj)
		db.Write(obj)
		lock.Unlock()
	}
}

func randomSleep() {
	sleep := time.Duration(rand.Intn(999)+1) * time.Millisecond
	time.Sleep(sleep)
}
