package queue

type Node struct {
	Value any
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
	length int
}

func(queue *Queue) Peek() any {
	if queue.head == nil {
		return nil
	}

	return queue.head.Value
}

func(queue *Queue) Dequeue() any {
	if queue.head == nil {
		return nil
	}
	queue.length--

	var head *Node= queue.head
	queue.head = head.next
	head.next = nil
	return head.Value
}

func (queue *Queue) Enqueue(value any) {
	var node Node = Node{
		Value: value,
	}
	queue.length++
	if queue.tail == nil {
		queue.tail , queue.head = &node, &node
		return
	}

	queue.tail.next = &node
	queue.tail = &node

}