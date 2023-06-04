package main

import "fmt"

//node represents a node in the linked list

type node struct {
	data interface{}
	next *node
}

// LinkedList represents a linked list
type LinkedList struct {
	head *node
	size int
}

// AddNode adds a new node to the linked list
func (ll *LinkedList) AddNode(data interface{}) {
	newNode := &node{data: data, next: nil}

	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
		ll.size++
	}
}

// DisplayList displays the elements of the linked list
func (ll *LinkedList) DisplayList() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (ll *LinkedList) InsertAtBeginning(data interface{}) {
	newNode := &node{data: data, next: ll.head}
	ll.head = newNode
	ll.size++
}

func (ll *LinkedList) InsertAtPosition(data interface{}, position int) {
	if position <= 0 {
		ll.InsertAtBeginning(data)
		return
	}
	newNode := &node{data: data}
	current := ll.head
	for count := 0; current.next != nil && count < position; count++ {
		if count == position-1 {
			newNode.next = current.next
			current.next = newNode
		}
		current = current.next
		ll.size++
	}
}

func (ll *LinkedList) Remove(data interface{}) {
	//newNodeToRemove := &node{data: data}
	if ll.head == nil {
		return
	}
	if ll.head.data == data {
		ll.head = ll.head.next
		ll.size--
		return
	}
	current := ll.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			return
		}
		current = current.next
		ll.size--
	}

}
func (ll *LinkedList) SearchSuccess(data interface{}) bool {
	current := ll.head
	for current.next != nil {
		if current.next.data == data {
			return true
		}
		current = current.next
	}
	return false
}

func main____________________() {
	ll := LinkedList{}
	//for i := 0; i <= 6; i++ {
	//	ll.AddNode(i)
	//	//ll.Remove(i)
	//}
	ll.InsertAtBeginning(1000)
	ll.Remove(1000)
	fmt.Println(ll.size)
	ll.DisplayList()
}
