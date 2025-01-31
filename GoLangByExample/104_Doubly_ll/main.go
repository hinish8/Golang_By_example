package main

type node struct {
	val  int
	next *node
	prv  *node
}

var head *node = nil
var curr *node = nil
