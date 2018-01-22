package linkedlistloop

import (
	"testing"
)

func TestList_CheckLoopOneNode(t *testing.T) {

	head := &Node{1, nil}

	result := CheckLoop(head)

	if result != nil {
		t.Fatalf("Check One Node Failed")
	}
}

func TestList_CheckLoopTwoNodeNoLoop(t *testing.T) {

	n1 := &Node{2, nil}
	head := &Node{1, n1}

	result := CheckLoop(head)

	if result != nil {
		t.Fatalf("Check One Node Failed")
	}
}

func TestList_CheckLoopTwoNodeLoop(t *testing.T) {

	n1 := &Node{2, nil}
	head := &Node{1, n1}
	n1.Next = head

	result := CheckLoop(head)

	if result == nil {
		t.Fatalf("Check One Node Failed")
	}
}

func TestList_CheckLoopThreeNodeNoLoop(t *testing.T) {

	n2 := &Node{2, nil}
	n1 := &Node{1, n2}
	head := &Node{0, n1}

	result := CheckLoop(head)

	if result != nil {
		t.Fatalf("Check One Node Failed")
	}
}

func TestList_CheckLoopThreeNodeLoop(t *testing.T) {

	n2 := &Node{2, nil}
	n1 := &Node{1, n2}
	head := &Node{0, n1}

	n2.Next = head
	result := CheckLoop(head)

	if result == nil {
		t.Fatalf("Check One Node Failed")
	}
}

func TestList_CheckLoopStartPoint(t *testing.T) {
	n6 := &Node{6, nil}
	n5 := &Node{5, n6}
	n4 := &Node{4, n5}
	n3 := &Node{3, n4}
	n2 := &Node{2, n3}
	n1 := &Node{1, n2}
	head := &Node{0, n1}

	//make loop
	n6.Next = n3

	result := CheckLoop(head)

	if result != n3 {
		t.Fatalf("Check loop start Failed")
	}
}