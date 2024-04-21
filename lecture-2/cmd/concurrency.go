package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"time"
)

func testChannel[T constraints.Integer](arr []T) T {
	ch := make(chan T)
	mid := len(arr) / 2
	go func() {
		var sum T
		for i := 0; i < mid; i++ {
			sum += arr[i]
		}
		ch <- sum
	}()
	go func() {
		var sum T
		for i := mid; i < len(arr); i++ {
			sum += arr[i]
		}
		ch <- sum
	}()
	return <-ch + <-ch
}

func testCloseChannel() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// send 0 value to channel, and set the channel to close
		close(ch)
	}()

	for i := 0; ; i++ {
		v, ok := <-ch
		if !ok {
			fmt.Printf("v: %d, ok: %v\n", v, ok)
			break
		}
		fmt.Printf("v: %d, ok: %v\n", v, ok)
	}
	fmt.Println("channel is closed")
}

func testCloseChannel2() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Printf("v: %d\n", v)
	}
	fmt.Println("channel is closed")
}

func testSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		sum := 0
		for i := 0; i < 100000; i++ {
			sum += i
		}
		//time.Sleep(1 * time.Second)
		ch1 <- sum
	}()
	go func() {
		sum := 0
		for i := 0; i < 100000; i++ {
			sum += i
		}
		//time.Sleep(1 * time.Second)
		ch2 <- sum
	}()

	select {
	case v := <-ch1:
		println("ch1 is selected ", v)
	case v := <-ch2:
		println("ch2 is selected ", v)
	}
}

func testDefaultSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		sum := 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		//time.Sleep(1 * time.Second)
		ch1 <- sum
	}()
	go func() {
		sum := 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		//time.Sleep(1 * time.Second)
		ch2 <- sum
	}()

	for {
		select {
		case v := <-ch1:
			println("ch1 is selected ", v)
			// should use goto to break the loop, since break will only break the select
			goto end
		case v := <-ch2:
			println("ch2 is selected ", v)
			// should use goto to break the loop, since break will only break the select
			goto end
		default: // for...select...default is polling IO model, for...select is blocking IO model
			println("default is selected")
		}
	}
end:
	println("end")
}

func main() {
	// array size
	//const size = 1000000000 // when size is large, parallel sum is faster than sequential sum
	const size = 1000 // when size is small, parallel sum is slower than sequential sum

	// create a large array
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i % 100
	}

	// parallel sum by channel
	t := time.Now()
	println(testChannel(arr))
	println(time.Since(t))

	// sequential sum
	t = time.Now()
	sum := 0
	for _, v := range arr {
		sum += v
	}
	println(sum)
	println(time.Since(t))

	// test close channel
	testCloseChannel()
	testCloseChannel2()

	// test select
	testSelect()
	testDefaultSelect()
}
