package queue

import (
	"errors"
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

func (q *Queue) Back() any {
	defer q.lock.Unlock()

	q.lock.Lock()
	if q.last == nil {
		return nil
	}
	return q.last.Value
}

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

func (q *Queue) Font() any {
	defer q.lock.Unlock()

	q.lock.Lock()
	if q.first == nil {
		return nil
	}
	return q.first.Value
}

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

func (q *Queue) Size() int {
	return q.length
}

func (q *Queue) value(node *Node) (any, error) {
	v := node.Value
	node = nil
	return v, nil
}
