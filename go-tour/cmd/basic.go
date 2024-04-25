package main

import (
	"fmt"
	baby "math"
	"math/rand"
	"strconv"
)

func packageTest() {
	fmt.Println("My favorite number is", rand.Intn(5))
	fmt.Println("My favorite number is", baby.Max(2, 10))
}

func functionsTest1(x int, y int) int {
	return x + y
}

func functionsTest2(x, y int) int {
	return x + y
}

func functionsTest3(x int, y ...int) int {
	total := x
	for _, num := range y {
		total += num
	}
	return total
}

func multipleResultsTest1() (int, int) {
	return 3, 4
}

func multipleResultsTest2() (x, y int) {
	x = 3
	y = 4
	return
}

func zeroTest() {
	var (
		i int
		f float64
		b bool
		s string
	)

	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// condition only works with boolean type
	if b {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func typeConversionsTest() {
	var x, y int = 3, 4
	var f float64 = baby.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	s := "556"
	n, e := strconv.Atoi(s)
	if e != nil {
		panic("convertion error")
	}

	fmt.Println(n)
}

// define const variables and enum
const (
	Big   = 1 << 100
	Small = Big >> 99
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func printWeekday() {
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}

func main() {
	packageTest()
	fmt.Println(functionsTest1(1, 2))
	fmt.Println(functionsTest2(1, 2))
	fmt.Println(functionsTest3(1, 2, 3, 4, 5))
	fmt.Println(multipleResultsTest1())
	fmt.Println(multipleResultsTest2())
	zeroTest()
	typeConversionsTest()
	//fmt.Println(Big, Small)
	fmt.Println(float64(Big), Small)
	printWeekday()
}
