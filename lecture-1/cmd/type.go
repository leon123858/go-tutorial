package main

import (
	"fmt"
	"lec1/pkg/boo"
	"strings"
	"sync"
	"time"
)

func testPointer() {
	arr := [3]int{0, 1, 2}
	println(arr[0])
	// go array not base on pointer
	// println(arr)
	// Go has no pointer arithmetic
	// x++
	// create a int pointer
	var a *int
	// case 1 set ptr by *
	a = new(int) // create space for *int
	*a = 5
	// case 2 set ptr by &
	//b := 5
	//a = &b
	println(a, ":", *a)
}

func testStructPtr() {
	//type boo struct {
	//	// 大寫開頭, public (不同 pkg 可以 access)
	//	Foo int
	//	// 小寫開頭, private (同 pkg 可以 access)
	//	bar int
	//}
	foo := new(boo.Boo)
	*foo = boo.Boo{Foo: 20}
	//*foo = boo.Boo{Foo: 20, bar: 10}
	println(foo.Foo)
	//println((*foo).bar)
}

func testStructPtr2() {
	type bar struct {
		// 大寫開頭, public (不同 pkg 可以 access)
		Foo int
		// 小寫開頭, private (同 pkg 可以 access)
		bar int
	}
	foo := new(bar)
	*foo = bar{Foo: 20, bar: 10}
	println(foo.Foo)
	println((*foo).bar)
}

func testStructLiteral() {
	type Vertex struct {
		X, Y int
		Z    string
	}

	var (
		v1 = Vertex{1, 2, "aaa"}  // has type Vertex
		v2 = Vertex{X: 1}         // Y:0 is implicit
		v3 = Vertex{}             // X:0 and Y:0
		p  = &Vertex{1, 2, "bbb"} // has type *Vertex
		p2 = new(Vertex)
	)

	fmt.Println(v1, p, v2, v3, p2)
}

func testSliceRefArray() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// ref array and create slice
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// slice do not store data, just ref to array
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func testSliceRefArray2() {
	// call by ref: slice, map, channel
	// call by value: array, struct, int, string, .....
	// call by ptr: any type use ptr
	// note: we can use copy to copy slice

	// slice
	mySlice := []int{1, 2, 3, 4, 5}
	func(s []int) {
		s[0] = 100
	}(mySlice)
	fmt.Println(mySlice)

	// array
	myArray := [5]int{1, 2, 3, 4, 5}
	func(a [5]int) {
		a[0] = 100
	}(myArray)
	fmt.Println(myArray)

	// channel
	myChan := make(chan int)
	go func(c chan int) {
		c <- 100
	}(myChan)
	fmt.Println(<-myChan)

	// struct
	type myStruct struct {
		Foo int
	}
	myStructObj := myStruct{Foo: 10}
	func(s myStruct) {
		s.Foo = 100
	}(myStructObj)
	fmt.Println(myStructObj)

	// ptr struct
	myStructPtr := &myStruct{Foo: 10}
	func(s *myStruct) {
		s.Foo = 100
	}(myStructPtr)
	fmt.Println(*myStructPtr)

	// copy slice
	mySlice2 := []int{1, 2, 3, 4, 5}
	mySlice3 := make([]int, len(mySlice2))
	copy(mySlice3, mySlice2)
	func(s []int) {
		s = append(s, 6)
	}(mySlice3)
	fmt.Println(mySlice3)
}

func testSliceCapacity() {
	a := new([5]int)
	a = &([5]int{1, 2, 3, 4, 5})
	s := a[:]
	ss := a[:3]
	fmt.Println(len(s), cap(s))
	fmt.Println(len(ss), cap(ss)) // cap is 5, but len is 3

	// slice can be resized
	ss = append(ss, 6)
	fmt.Println(len(ss), cap(ss)) // cap is 5, but len is 4
	fmt.Println(ss, s, a)         // a[3] is changed!
	orginSS := &ss[0]
	fmt.Printf("address of slice %p add of Arr %p \n", &ss[0], &a)
	ss = append(ss, 7)
	fmt.Println(len(ss), cap(ss)) // cap is 5, and len is 5
	fmt.Println(ss, s, a)         // a[4] is changed
	fmt.Printf("address of slice %p add of Arr %p \n", &ss[0], &a)
	if orginSS == &ss[0] {
		fmt.Println("same address")
	} else {
		panic("diff address")
	}
	ss = append(ss, 8)
	fmt.Println(len(ss), cap(ss)) // cap is 10, and len is 6
	fmt.Printf("address of slice %p add of Arr %p \n", &ss[0], &a)
	if orginSS == &ss[0] {
		panic("same address")
	} else {
		// because slice is ref to array, when over capacity,
		// go will create new array and copy data
		fmt.Println("diff address")
	}
}

func testMakeSlice() {
	// make can create dynamic type variable like slice, map, channel
	// new return ptr, make return type
	a := make([]int, 5)
	fmt.Println(a)
	b := make(map[string]int)
	fmt.Println(b)
	c := make(chan int)
	fmt.Println(c)
}

func testSliceInSlice() {
	// slice in slice
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", board[i])
	}
	// array in array
	board2 := [3][3]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	board2[0][0] = "X"
	board2[2][2] = "O"
	board2[1][2] = "X"
	board2[1][0] = "O"
	board2[0][2] = "X"
	for i := 0; i < len(board2); i++ {
		fmt.Printf("%s\n", board2[i])
	}
}

func testRange() {
	// range can return index and value
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	// if you only want the value
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
	// range for map return key and value
	pow2 := map[int]int{1: 1, 2: 4, 3: 9, 4: 16}
	for k, v := range pow2 {
		fmt.Printf("%d**2 = %d\n", k, v)
	}
	// create a inverse map
	pow3 := make(map[int]int)
	for k, v := range pow2 {
		pow3[v] = k
	}
	fmt.Println(pow3)
}

func testExerciseSlice() {
	result1 := func(dx, dy int) [][]uint8 {
		result := make([][]uint8, dy)
		for i := 0; i < dy; i++ {
			for j := 0; j < dx; j++ {
				result[i] = append(result[i], 0)
			}
		}
		return result
	}(5, 10)
	fmt.Println(result1)

	result2 := func(dx, dy int) [][]uint8 {
		result := make([][]uint8, dy)
		for i := range result {
			result[i] = make([]uint8, dx)
		}
		return result
	}
	fmt.Println(result2(5, 10))
}

func testExerciseMap() {
	result1 := func(s string) map[string]int {
		var newStringSlice []string
		for i, word := 0, ""; i < len(s); i++ {
			if s[i] == ' ' {
				newStringSlice = append(newStringSlice, word)
				word = ""
			} else {
				word += string(s[i])
				if i == len(s)-1 {
					newStringSlice = append(newStringSlice, word)
				}
			}
		}
		cache := make(map[string]int)
		for _, v := range newStringSlice {
			value, ok := cache[string(v)]
			if ok {
				cache[string(v)] = value + 1
			} else {
				cache[string(v)] = 1
			}
		}
		return cache
	}("I am learning Go Go !")
	fmt.Println(result1)

	result2 := func(s string) map[string]int {
		newStringSlice := strings.Split(s, " ")
		cache := make(map[string]int)
		for _, v := range newStringSlice {
			cache[v]++
		}
		return cache
	}("I am learning Go Go !")
	fmt.Println(result2)
}

func testFuncValues() {
	// function values
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))

	// immediate function
	result := func(x, y int) int {
		return x + y
	}(1, 1)
	fmt.Println(result)

	// closure, 可以用來: 1. 保持變數狀態 2. 保持變數私有 3. 保持變數不被污染
	addClosure := func(x int) func(int) int {
		return func(y int) int {
			return add(x, y)
		}
	}
	add1 := addClosure(1)
	fmt.Println(add1(1))
	add2 := addClosure(2)
	fmt.Println(add2(1))

	// closure 2
	counter := func(initValue int) func() int {
		i := initValue
		return func() int {
			i++
			return i
		}
	}
	count := counter(0)
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())

	// closure 3 [doc handle]
	type DocType int
	const (
		PDF DocType = iota
		WORD
	)
	// low level function
	docHandle := func(include string, t DocType) string {
		switch t {
		case PDF:
			return include + ".pdf"
		case WORD:
			return include + ".doc"
		default:
			return include
		}
	}
	// high level function, wrap low level function to high level function
	// 直接綁定固定參數或持續狀態, 並且可以隨時更換低階函數: func(a,b) => func(a) func(b)
	wrapDocHandle := func(t DocType) func(string) string {
		return func(s string) string {
			return docHandle(s, t)
		}
	}
	pdfHandler := wrapDocHandle(PDF)
	fmt.Println(pdfHandler("pdf_include"))
	wordHandler := wrapDocHandle(WORD)
	fmt.Println(wordHandler("word_include"))

	// closure 4 [體積], [面積], [長度] 共用底層函式, 降低重複程式碼
	volume := func(x, y, z int) int {
		return x * y * z
	}
	_area := func(z int) func(x, y int) int {
		return func(x, y int) int {
			return volume(x, y, z)
		}
	}
	area := _area(1)
	_length := func(y int) func(x int) int {
		return func(x int) int {
			return area(x, y)
		}
	}
	length := _length(1)
	fmt.Println(volume(1, 2, 3))
	fmt.Println(area(1, 2))
	fmt.Println(length(3))
}

func testFibonacci() {
	// 高技巧寫法
	fibGenerator := func() func() int {
		arr2 := [2]int{0, 1}
		count := 0x00
		return func() int {
			// update next 2 value, a, b, a+b, a+b+b
			ret := arr2[count]
			arr2[count] = arr2[0] + arr2[1]
			count ^= 0x01 // go can support low level bit operation
			return ret
		}
	}
	fib := fibGenerator()
	fmt.Println("Fibonacci:")
	for i := 0; i < 10; i++ {
		fmt.Print(fib(), " ")
	}
	fmt.Println()
	fmt.Println("---end---")
	// 低技巧寫法
	fibGenerator2 := func() func() int {
		arr := [2]int{0, 1}
		index := -1
		return func() int {
			index++ // index = index + 1
			if index < 2 {
				return arr[index]
			}
			index = 0
			arr[0] = arr[0] + arr[1]
			arr[1] = arr[0] + arr[1]
			return arr[index]
		}
	}
	fib2 := fibGenerator2()
	fmt.Println("Fibonacci:")
	for i := 0; i < 10; i++ {
		fmt.Print(fib2(), " ")
	}
	fmt.Println()
	fmt.Println("---end---")
}

func testFibSpeed() {
	fibGenerator := func() func() int {
		arr2 := [2]int{0, 1}
		count := 0x00
		return func() int {
			// update next 2 value, a, b, a+b, a+b+b
			ret := arr2[count]
			arr2[count] = arr2[0] + arr2[1]
			//count = (count + 1) % 2
			count ^= 0x01 // go can support low level bit operation
			return ret
		}
	}
	fibGenerator2 := func() func() int {
		arr := [2]int{0, 1}
		index := -1
		return func() int {
			index++ // index = index + 1
			if index < 2 {
				return arr[index]
			}
			index = 0
			arr[0] = arr[0] + arr[1]
			arr[1] = arr[0] + arr[1]
			return arr[index]
		}
	}
	fib := fibGenerator()
	fib2 := fibGenerator2()
	timeTrack := func(start time.Time, name string) {
		elapsed := time.Since(start)
		fmt.Printf("%s took %s\n", name, elapsed)
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	const bigInt = 100000000
	go func() {
		defer wg.Done()
		defer timeTrack(time.Now(), "fib2")
		for i := 0; i < bigInt; i++ {
			fib2()
		}
	}()
	go func() {
		defer wg.Done()
		defer timeTrack(time.Now(), "fib")
		for i := 0; i < bigInt; i++ {
			fib()
		}
	}()
	wg.Wait()
}

func main() {
	testPointer()
	testStructPtr()
	testStructPtr2()
	testStructLiteral()
	testSliceRefArray()
	testSliceRefArray2()
	testSliceCapacity()
	testMakeSlice()
	testSliceInSlice()
	testRange()
	testExerciseSlice()
	testExerciseMap()
	testFuncValues()
	testFibonacci()
	testFibSpeed()
}
