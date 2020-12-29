package stacks

import "fmt"

// for stack type definition see Stack Set

// Queue is a FIFO data structure implemented using two stacks
type Queue struct {
	Front Stack
	Back  Stack
}

// ToString provides a string representation for a queue
func (q *Queue) ToString() string {
	return fmt.Sprintf("Front: %s\nBack: %s\n", q.Front.ToString(), q.Back.ToString())
}

// O(n)
func (q *Queue) transfer(source, destination *Stack) {
	for source.Top != nil {
		destination.Push(source.Pop())
	}
}

// Add is O(1) without a transfer and O(n) with a transfer
func (q *Queue) Add(data int) {
	if q.Back.Top == nil {
		q.transfer(&q.Front, &q.Back)
	}
	q.Back.Push(data)
}

// Remove is O(1) without a transfer and O(n) with a transfer
func (q *Queue) Remove() int {
	if q.Front.Top == nil {
		q.transfer(&q.Back, &q.Front)
		if q.Front.Top == nil {
			panic("queue is empty")
		}
	}

	return q.Front.Pop()
}
