package linkedList

import (
	"fmt"
)

type Node struct {
	Value any
	prev *Node `default:"nil"`
	next *Node `default:"nil"`
}

type DoublyLinkedList struct {
	head *Node `default:"nil"`
	tail *Node `default:"nil"`
	length int `default:"0"`
}

type LinkedListInterface interface {
	GetLength() int
	InsertAt() error
	Remove() error
	RemoveAt() error
	Append()
	Prepend()
	Get() any
}

func (dLinkedList *DoublyLinkedList) GetLength() int {
	return dLinkedList.length
}

func (dLinkedList *DoublyLinkedList) Prepend(value any){
	node := Node{
		Value: value,
	}

	dLinkedList.length++
	if dLinkedList.head == nil {
		dLinkedList.head, dLinkedList.tail = &node, &node
		return
	}

	node.next = dLinkedList.head
	dLinkedList.head.prev = &node
	dLinkedList.head = &node 
}

func (dLinkedList *DoublyLinkedList) Append(value any){
	dLinkedList.length++
	node := Node{
		Value: value,
	}
	if dLinkedList.tail == nil {
		dLinkedList.head , dLinkedList.tail = &node, &node
		return
	}

	node.prev = dLinkedList.tail
	dLinkedList.tail.next = &node
	dLinkedList.tail = &node
}

func(dLinkedList *DoublyLinkedList) InsertAt(value any, index int) error {
	if index > dLinkedList.length {
		return fmt.Errorf("index (\"%d\") out of range", index)
	} else if index == dLinkedList.length {
		dLinkedList.Append(value)
		return nil
	} else if index == 0 {
		dLinkedList.Prepend(value)
		return nil
	}

	dLinkedList.length++
	var curr *Node = dLinkedList.getAt(index)

	var node Node = Node{
		Value: value,
		next: curr,
		prev: curr.prev,
	}

	curr.prev = &node
	node.prev.next = &node

	return nil
}

func (dLinkedList *DoublyLinkedList) Remove(value any) error {
	curr := dLinkedList.head
	for i:=0; i<dLinkedList.length; i++ {
		if curr.Value == value {
			break
		}
		curr = curr.next
	}

	if curr == nil {
		return fmt.Errorf("no node with the value \"%v\" was found", value)
	}

	dLinkedList.removeNode(curr)

	return nil
}

func (dLinkedList *DoublyLinkedList) RemoveAt(index int) error {
	var node *Node = dLinkedList.getAt(index) 

	if node == nil {
		return fmt.Errorf("node not found at index %d", index)
	}
	dLinkedList.removeNode(node)
	return nil
}

func (dLinkedList *DoublyLinkedList) removeNode(node *Node) {
	dLinkedList.length--
	
	if dLinkedList.length == 0 {
		dLinkedList.head, dLinkedList.tail = nil, nil
		return
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if node == dLinkedList.head {
		dLinkedList.head = node.next
	}

	if node == dLinkedList.tail {
		dLinkedList.tail = node.prev
	}

	node.prev, node.next = nil, nil
}

func (dLinkedList *DoublyLinkedList) Get(index int) any {
	return dLinkedList.getAt(index).Value
}

func (dLinkedList *DoublyLinkedList) getAt(index int) *Node {
	var curr *Node = dLinkedList.head
	for i:=0; curr != nil && i<index; i++ {
		curr = curr.next
	}
	return curr
}