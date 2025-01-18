package queue

import (
	"errors"
	"sync"
)

// Node struct
// Value is the value of the node,
// Next points to the next node
type Node struct {
	Value any
	Next  *Node
	Prev  *Node
}

// Queue struct
// first The first node in the queue,
// last The last node in the queue,
// length The number of nodes in the queue
type Queue struct {
	lock   sync.Mutex
	first  *Node
	last   *Node
	length int
}

// Return a new queue pointer
func NewQueue() *Queue {
	return &Queue{}
}

// Returns the value of the last node in the queue
func (q *Queue) Back() any {
	defer q.lock.Unlock()

	q.lock.Lock()
	if q.last == nil {
		return nil
	}
	return q.last.Value
}

// Check whether there are nodes in the queue whose value is v
func (q *Queue) Contains(v any) bool {
	defer q.lock.Unlock()

	q.lock.Lock()
	re := false
	q.Each(func(n *Node) bool {
		re = n.Value == v
		return !re
	})
	return re
}

// Go through all nodes in the queue and pass in the current node to the callback function, 
// and the callback function returns false to exit the loop
func (q *Queue) Each(exec func(node *Node) bool) {
	node := q.first
	for {
		if !exec(node) {
			break
		}
		if node == nil || node.Next == nil {
			break
		}
		node = node.Next
	}
}

// Return the value of the first node in the queue
func (q *Queue) Font() any {
	defer q.lock.Unlock()

	q.lock.Lock()
	if q.first == nil {
		return nil
	}
	return q.first.Value
}

// Returns the value of the first or last node in the queue and deletes it
func (q *Queue) pop(node *Node) error {
	if node == nil {
		return errors.New("node not nil")
	}

	if node == q.first {
		next := node.Next
		if next != nil {
			next.Prev = nil
		}
		q.first = next
	} else {
		prev := node.Prev
		if node == q.last {
			prev.Next = nil
			q.last = prev
		} else {
			prev.Next = node.Next
		}
	}

	q.length--
	return nil
}

// Return the value of the first node in the queue and remove it
func (q *Queue) Pop() (any, error) {
	defer q.lock.Unlock()

	q.lock.Lock()
	node := q.first
	if node == nil {
		return nil, errors.New("queue is empty")
	}

	err := q.pop(node)
	if err != nil {
		return nil, err
	}
	return q.value(node)
}

// Returns the value of the last node in the queue and deletes it
func (q *Queue) Popend() (any, error) {
	defer q.lock.Unlock()

	q.lock.Lock()
	node := q.last
	if node == nil {
		if q.first == nil {
			return nil, errors.New("queue is empty")
		}
		node = q.first
	}

	err := q.pop(node)
	if err != nil {
		return nil, err
	}
	return q.value(node)
}

// Join the queue last
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

// Delete a node whose queue value is v
func (q *Queue) Remove(v any) (bool, error) {
	defer q.lock.Unlock()

	q.lock.Lock()
	if q.Size() == 0 || q.first == nil {
		return false, errors.New("queue is empty")
	}

	var node *Node
	q.Each(func(n *Node) bool {
		if n.Value == v {
			node = n
			return false
		}
		return true
	})

	if node == nil {
		return false, errors.New("not find v in queue")
	}

	err := q.pop(node)
	return err == nil, err
}

// Returns the queue length
func (q *Queue) Size() int {
	return q.length
}

// Returns the node value
func (q *Queue) value(node *Node) (any, error) {
	v := node.Value
	node = nil
	return v, nil
}
