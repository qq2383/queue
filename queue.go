package queue

import (
	"sync"
)

type Node struct {
	Value any
	Next  *Node
	Prev  *Node
}

type Queue struct {
	lock   sync.Mutex
	first  *Node
	last   *Node
	length int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Put(v any) {
	defer q.lock.Unlock()

	q.lock.Lock()
	node := &Node{Value: v}
	if q.first == nil {
		q.first = node
	} else {
		if q.last == nil {
			first := q.first
			first.Next = node
			node.Prev = first
		} else {
			last := q.last
			last.Next = node
			node.Prev = last
		}
		q.last = node
	}
	q.length++
}

func (q *Queue) Pop() any {
	defer q.lock.Unlock()

	q.lock.Lock()
	node := q.first
	if node == nil {
		return nil
	}
	q.length--
	next := node.Next
	if next == nil {
		q.first = nil
	} else {
		next.Prev = nil
		q.first = next
		if next == q.last {
			q.last = nil
		}
	}
	return node.Value
}

func (q *Queue) Size() int {
	return q.length
}

func (q *Queue) Font() any {
	defer q.lock.Unlock()

	q.lock.Lock()
	if q.first == nil {
		return nil
	}
	return q.first.Value
}

func (q *Queue) Back() any {
	defer q.lock.Unlock()

	q.lock.Lock()
	if q.last == nil {
		return nil
	}
	return q.last.Value
}
