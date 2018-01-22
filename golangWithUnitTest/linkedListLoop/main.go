package main

import (
	 "./linkedlistloop"
	"fmt"
)

func main () {
	n6 := &linkedlistloop.Node{6, nil}
	n5 := &linkedlistloop.Node{5, n6}
	n4 := &linkedlistloop.Node{4, n5}
	n3 := &linkedlistloop.Node{3, n4}
	n2 := &linkedlistloop.Node{2, n3}
	n1 := &linkedlistloop.Node{1, n2}
	head := &linkedlistloop.Node{0, n1}

	//make loop
	n6.Next = n3

	result := linkedlistloop.CheckLoop(head)
	if result != nil {
		fmt.Println("Loop starts at node:", result.Data)
	}
}