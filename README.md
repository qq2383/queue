# Overview

A Queue for Go package

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