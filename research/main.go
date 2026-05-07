package main

import "fmt"

// Using pointer on(type) is implemented here
// curr is pointing to type Node. same with prev and key
// key is assigned to an empty interface so it can accept any datatype
type Node struct {
	next *Node
	prev *Node
	key  any
}

type linkedlist struct {
	head *Node
	tail *Node
}

func (ll *linkedlist) Push(key any) {
	list := &Node{
		next: ll.head,
		key:  key,
	}

	if ll.head != nil {
		ll.head.prev = list
	}

	ll.head = list

	l := ll.head

	for l.next != nil {

		l = l.next
	}

	ll.tail = l

}

func (ll *linkedlist) display() {
	list := ll.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

// Normal dispaly function
// func dispaly(list *Node) {
// 	for list != nil {
// 		fmt.Println("%+v ->", list.key)

// 		list = list.next
// 	}

// 	fmt.Println()
// }

// reverse the linked stingy

func (ll *linkedlist) reverse() {
	currentnode := ll.head
	var next *Node
	var Previous *Node
	ll.tail = ll.head

	for currentnode != nil {
		next, currentnode.next = currentnode.next, Previous
		Previous, currentnode = currentnode, next
	}
	ll.head = Previous
	ll.display()
}

func main() {
	link := linkedlist{}
	link.Push(9)
	link.Push(12)
	link.Push(15)
	link.Push(8)
	link.Push(1)
	link.Push(3)

	fmt.Println("==============================")
	fmt.Printf("Head: %v\n", link.head.key)
	fmt.Printf("Tail: %v\n", link.tail.key)
	link.display()
	fmt.Println("==============================")
	link.reverse()
	fmt.Printf("head: %v\n", link.head.key)
	fmt.Printf("tail: %v\n", link.tail.key)
	fmt.Println("==============================")
}
