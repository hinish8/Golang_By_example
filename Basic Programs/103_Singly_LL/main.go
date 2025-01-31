package main

import "fmt"

type node struct {
	val  int
	next *node
}

var ll *node   // Head pointer
var curr *node // Last node pointer

func InsertEnd(num int) {
	nn := &node{val: num, next: nil}
	if ll == nil {
		ll = nn
		curr = nn
		return
	}
	curr.next = nn
	curr = nn
}

func InsertBegin(num int) {
	nn := &node{val: num, next: nil}
	if ll == nil {
		ll = nn
		curr = nn
		return
	}
	nn.next = ll
	ll = nn
}

func DeleteEnd() *node {
	if ll == nil {
		fmt.Println("Linked list is empty")
		return nil
	}
	if ll.next == nil { // Only one node
		ll = nil
		curr = nil
		return nil
	}
	temp := ll
	for temp.next.next != nil {
		temp = temp.next
	}
	temp.next = nil
	curr = temp
	return curr
}

func DeleteFront() {
	if ll == nil {
		fmt.Println("Linked list is empty")
		return
	}
	if ll.next == nil { // Only one node
		ll = nil
		curr = nil
		return
	}
	ll = ll.next
}

func Front() *node {
	return ll
}

func Size() int {
	count := 0
	temp := ll
	for temp != nil {
		count++
		temp = temp.next
	}
	return count
}

func Display() {
	temp := ll
	if temp == nil {
		fmt.Println("List is empty")
		return
	}
	for temp != nil {
		fmt.Printf("%d -> ", temp.val)
		temp = temp.next
	}
	fmt.Println("nil")
}

func main() {
	InsertEnd(10)
	InsertEnd(20)
	InsertBegin(5)
	Display()
	fmt.Println("Size:", Size()) // 3
	DeleteEnd()
	Display()
	DeleteFront()
	Display()
	fmt.Println("Size:", Size()) // 1
}
