package linkedlistloop

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func CheckLoop(head *Node) *Node {
	fmt.Println("Check Loop: In")
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			fmt.Println("Loop Exist")
			break
		}
	}

	if fast == nil || fast.Next == nil {
		fmt.Println("No Loop")
		return nil
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	fmt.Println("Check Loop: Out")
	return fast
}