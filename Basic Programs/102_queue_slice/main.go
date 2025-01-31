package main

import (
	"errors"
	"fmt"
	"sync"
)

type queue struct {
	q []string
	l sync.RWMutex
}

func (cq *queue) Push(str string) {
	cq.l.Lock()
	defer cq.l.Unlock()
	cq.q = append(cq.q, str)
}

func (cq *queue) Pop() (string, error) {
	cq.l.Lock()
	defer cq.l.Unlock()
	len := len(cq.q)
	if len > 0 {
		val := cq.q[0]
		cq.q = cq.q[1:]
		return val, nil
	}
	return "", errors.New("queue is empty")
}

func (cq *queue) Size() int {
	cq.l.RLock()
	defer cq.l.RUnlock()
	return len(cq.q)
}

func (cq *queue) IsEmpty() bool {
	cq.l.RLock()
	defer cq.l.RUnlock()
	return len(cq.q) == 0
}

func (cq *queue) Front() (string, error) {
	cq.l.RLock()
	defer cq.l.RUnlock()
	len := len(cq.q)
	if len > 0 {
		return cq.q[0], nil
	}
	return "", errors.New("queue is empty")
}

func main() {

	queue := &queue{q: make([]string, 0)}

	fmt.Printf("Enqueue: A\n")
	queue.Push("A")
	fmt.Printf("Enqueue: B\n")
	queue.Push("B")
	fmt.Printf("Len: %d\n", queue.Size())

	for queue.Size() > 0 {
		frontVal, _ := queue.Front()
		fmt.Printf("Front: %s\n", frontVal)
		fmt.Printf("Dequeue: %s\n", frontVal)
		queue.Pop()
	}
	fmt.Printf("Length of queue is: %d\n", queue.Size())

}
