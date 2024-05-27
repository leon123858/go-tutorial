# Multithreaded Queue and Stack in Go

This is an example implementation of a multithreaded queue and stack in Go without using channels. The code demonstrates
how to create thread-safe queue and stack data structures using `sync.Mutex` for synchronization.

## Queue

The `Queue` struct represents a queue and has the following methods:

- `Enqueue(item int)`: Adds an item to the rear of the queue.
- `Dequeue() (int, bool)`: Removes and returns the item from the front of the queue. Returns `false` if the queue is
  empty.

## Stack

The `Stack` struct represents a stack and has the following methods:

- `Push(item int)`: Pushes an item onto the top of the stack.
- `Pop() (int, bool)`: Removes and returns the item from the top of the stack. Returns `false` if the stack is empty.

## Usage

1. Create a new queue and stack:
   ```go
   queue := &Queue{}
   stack := &Stack{}
   ```

2. Perform multithreaded operations on the queue:
   ```go
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
   ```

3. Perform multithreaded operations on the stack:
   ```go
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
   ```

## Synchronization

The code uses `sync.Mutex` to ensure thread safety for the queue and stack operations. Each method of the `Queue`
and `Stack` structs acquires a lock using `lock.Lock()` before accessing the shared data and releases the lock
using `defer lock.Unlock()` to prevent data races.

## Concurrency

The code demonstrates concurrent operations on the queue and stack using goroutines. Multiple goroutines are spawned to
perform enqueue/push operations simultaneously. The main goroutine waits for all the worker goroutines to complete
using `sync.WaitGroup` before proceeding with dequeue/pop operations.

## Output

The code includes `fmt.Printf` statements to display the enqueued/pushed items and the dequeued/popped items. The output
will show the order in which the items were added to and removed from the queue and stack.

Feel free to use and modify this code as needed for your specific requirements.