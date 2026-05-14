package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Linkedlist struct {
	head *Node
	tail *Node
}

func (ll *Linkedlist) Addnode(data int) {
	NewNode := &Node{data, nil}

	if ll.head == nil {
		ll.head = NewNode
	} else {
		Currentnode := ll.head

		for Currentnode.next != nil {
			Currentnode = Currentnode.next
		}

		Currentnode.next = NewNode
	}
}

func (ll *Linkedlist) Prinlist() {
	if ll.head == nil {
		fmt.Println("Empty list")
		return
	} else {
		Currentnode := ll.head
		for Currentnode != nil {
			fmt.Printf("%d'\n %d\n", Currentnode.data, Currentnode.next)
			Currentnode = Currentnode.next
		}

	}
	fmt.Println()
}

// Reverse function

func (ll *Linkedlist) Reverse() {
	if ll.head == nil {
		return
	}

	prev := (*Node)(nil)
	Currentnode := ll.head
	ll.tail = ll.head

	for Currentnode != nil {
		next := Currentnode.next
		Currentnode.next = prev
		prev = Currentnode
		Currentnode = next
	}
	ll.head = prev

}
func main() {

	var List Linkedlist
	List.Addnode(1)
	List.Addnode(2)
	List.Addnode(3)
	List.Addnode(4)
	List.Addnode(5)
	List.Prinlist()
	fmt.Println("Reversed singly")
	List.Reverse()
	List.Prinlist()

}
