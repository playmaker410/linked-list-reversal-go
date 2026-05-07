//Without pointers you cannot just change addresses. You have to copy everything into slice

package main

import "fmt"

type Node struct {
	curr *Node
	next *Node
	tail *Node
	val  any
}

func (n *Node) Display() {
	listall := n.curr
	for listall != nil {
		fmt.Println("->", listall.val)
		listall = listall.next
	}
	fmt.Println("======================")
}

func (n *Node) Reverse() {
	current := n.curr
	var previous *Node
	n.tail = n.curr

	for current != nil {
		next := current.next
		current.next = previous
		previous = current
		current = next
	}
	n.curr = previous

}

func main() {
	singly := &Node{}

	singly.curr = &Node{val: "ada", next: &Node{val: "obi", next: &Node{val: "john", next: &Node{val: "ude"}}}}

	singly.Display()
	singly.Reverse()
	singly.Display()
}
