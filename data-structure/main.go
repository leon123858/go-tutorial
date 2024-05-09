package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	items []int
	lock  sync.Mutex
}

func (q *Queue) Enqueue(item int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (int, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.items) == 0 {
		return 0, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

type Stack struct {
	items []int
	lock  sync.Mutex
}

func (s *Stack) Push(item int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.items) == 0 {
		return 0, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func main() {
	queue := &Queue{}
	stack := &Stack{}

	var wgQueue sync.WaitGroup
	for i := 0; i < 5; i++ {
		wgQueue.Add(1)
		go func(item int) {
			defer wgQueue.Done()
			queue.Enqueue(item)
			fmt.Printf("Enqueued item: %d\n", item)
		}(i)
	}
	wgQueue.Wait()

	for i := 0; i < 5; i++ {
		if item, ok := queue.Dequeue(); ok {
			fmt.Printf("Dequeued item: %d\n", item)
		}
	}

	var wgStack sync.WaitGroup
	for i := 0; i < 5; i++ {
		wgStack.Add(1)
		go func(item int) {
			defer wgStack.Done()
			stack.Push(item)
			fmt.Printf("Pushed item: %d\n", item)
		}(i)
	}
	wgStack.Wait()

	for i := 0; i < 5; i++ {
		if item, ok := stack.Pop(); ok {
			fmt.Printf("Popped item: %d\n", item)
		}
	}
}
