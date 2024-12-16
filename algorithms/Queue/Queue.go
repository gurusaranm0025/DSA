package queue

type node struct {
	value any
	next  *node
}

type Queue struct {
	length int
	head   *node
	tail   *node
}

func NewQueue() *Queue {
	return &Queue{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (q *Queue) Peek() any {
	if q.head == nil {
		return nil
	}

	return q.head.value
}

func (q *Queue) Dequeue() any {
	if q.head == nil {
		return nil
	}

	q.length--

	head := q.head
	q.head = q.head.next

	if q.length == 0 {
		q.tail = nil
	}

	return head.value
}

func (q *Queue) Enqueue(item any) {
	node := &node{
		value: item,
		next:  nil,
	}

	if q.tail == nil {
		q.head = node
		q.tail = node
		return
	}

	q.tail.next = node
	q.tail = node
}
