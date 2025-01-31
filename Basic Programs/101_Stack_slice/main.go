package main

import (
	"fmt"
	"sync"
)

type customStack struct {
	stack []string
	lock  sync.RWMutex
}

func (c *customStack) Push(name string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.stack = append(c.stack, name)
}

func (c *customStack) Pop() (string, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	length := len(c.stack)
	if length == 0 {
		return "", fmt.Errorf("pop error: Stack is empty")
	}

	// Get last element before removing it
	val := c.stack[length-1]
	c.stack = c.stack[:length-1] // Remove last element

	return val, nil
}

func (c *customStack) Top() (string, error) {
	c.lock.RLock() // Use read lock
	defer c.lock.RUnlock()

	length := len(c.stack)
	if length == 0 {
		return "", fmt.Errorf("top error: Stack is empty")
	}
	return c.stack[length-1], nil
}

func (c *customStack) Size() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return len(c.stack)
}

func (c *customStack) Empty() bool {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return len(c.stack) == 0
}

func main() {
	customStack := &customStack{stack: make([]string, 0)}
	fmt.Printf("Push: A\n")
	customStack.Push("A")
	fmt.Printf("Push: B\n")
	customStack.Push("B")
	fmt.Printf("Size: %d\n", customStack.Size())

	for customStack.Size() > 0 {
		frontVal, _ := customStack.Top()
		fmt.Printf("Front: %s\n", frontVal)
		fmt.Printf("Pop: %s\n", frontVal)
		customStack.Pop()
	}

	fmt.Printf("Size: %d\n", customStack.Size())
}
