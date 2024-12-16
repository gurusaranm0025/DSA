package stack

import "math"

type Node struct {
	Value any
	prev  *Node
}

type Stack struct {
	head   *Node
	length int
}

func NewStack() *Stack {
	return &Stack{
		length: 0,
		head:   nil,
	}
}

func (stack *Stack) Peek() any {
	if stack.head == nil {
		return nil
	}

	return stack.head.Value
}

func (stack *Stack) Push(value any) {
	var node Node = Node{
		Value: value,
	}

	stack.length++
	if stack.head == nil {
		stack.head = &node
		return
	}

	node.prev = stack.head
	stack.head = &node
}

func (stack *Stack) Pop() any {
	stack.length = int(math.Max(0, float64(stack.length)))

	if stack.length == 0 {
		node := stack.head
		stack.head = nil
		return node.Value
	}

	node := stack.head
	stack.head = stack.head.prev
	return node.Value
}
