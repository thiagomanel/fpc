package main

type node struct {
	value *chan bool
	next  *node
}

type Queue struct {
	head *node
	tail *node
}

func (q *Queue) Push(value *chan bool) {
	if q == nil {
		panic("Queue is nil")
	}
	newNode := &node{value: value}
	if q.Empty() {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}

func (q *Queue) ExtractFront() (*chan bool, bool) {
	if q == nil {
		panic("Queue is nil")
	}
	if q.Empty() {
		return nil, false
	}
	front := q.head.value
	if q.head == q.tail {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}
	return front, true
}

func (q *Queue) Empty() bool {
	if q == nil {
		panic("Queue is nil")
	}
	if q.head == nil && q.tail == nil {
		return true
	}
	return false
}
