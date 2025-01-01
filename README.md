# Overview

A Queue for Go package, Achieve FIFO LIFO

# API
```
func NewQueue() *Queue
```
NewQueue return a new queue pointer

```
func (q *Queue) Back() any
```
Back return the value of the first node in the queue

```
func (q *Queue) Contains(v any) bool
```
Check whether there are nodes in the queue whose value is v

```
func (q *Queue) Each(exec func(node *Node) bool)
```
Go through all nodes in the queue and pass in the current node to the callback function, and the callback function returns false to exit the loop

```
func (q *Queue) Font() any
```
Font return the value of the last node in the queue

```
func (q *Queue) Pop() (any, error)
```
Retrieve the value of the first node in the queue and remove it

```
func (q *Queue) Popend() (any, error)
```
Retrieve the value of the last node in the queue and remove it

```
func (q *Queue) Put(v any)
```
Join the queue last

```
func (q *Queue) Remove(v any) (bool, error)
```
Delete a node whose queue value is v

```
func (q *Queue) Size() int
```
Returns the queue length

# Example
```
package main

import (
	"fmt"
	"qq2383/queue"
	"sync"
	"time"
)

func main() {
	q := queue.NewQueue()

	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		i := 0
		for i < 100 {
			q.Put(i)
			// fmt.Printf("put to queue, value: %d\n", i)
			i++
			time.Sleep(time.Millisecond * 50)
		}

	}(&wg)

	time.Sleep(time.Millisecond * 100)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		for q.Size() !=0 {
			first := q.Font()
			if first != nil {
				v := q.Pop()
				fmt.Printf("get first: %d, size: %d\n", v, q.Size())
			}
			time.Sleep(time.Millisecond * 200)
		}

	}(&wg)

	wg.Wait()
}

```

## License

> Apache License 2.0 license