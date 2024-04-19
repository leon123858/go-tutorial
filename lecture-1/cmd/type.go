package main

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
	type boo struct {
		// 大寫開頭, public (不同 pkg 可以 access)
		Foo int
		// 小寫開頭, private (同 pkg 可以 access)
		bar int
	}
	//foo := new(find.Boo)
	foo := new(boo)
	*foo = boo{Foo: 20, bar: 10}
	println(foo.Foo)
	println((*foo).bar)
}

func main() {
	testPointer()
	testStructPtr()
}
