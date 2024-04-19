package main

import "math"

func testIfShort1() {
	// if statement
	if x := 10; x < 20 {
		println("x is less than 20")
	}
}

func testIfShort2(x *int) error {
	*x = 10
	return nil
}

func testIfElse() {
	// if else statement
	if x := 11; x > 20 {
		println("x is bigger than 20")
	} else if x > 10 {
		println("x is bigger than 10")
	} else {
		println("x is less than 10")
	}
}

func testMyLoopAndFunc(x float64) float64 {
	if x < 0 {
		panic("x should bigger than 0")
	}
	dist := x
	z := float64(1)
	for i := 0; i < 10; i++ {
		if inDist := math.Abs(x - z*z); inDist < dist {
			dist = inDist
		}
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func testMySwitch() {
	// can use init in switch
	switch os := "linux"; os {
	case "linux":
		println("Linux")
		// fallthrough can execute next case
		//fallthrough
	case "windows":
		println("Windows")
		//fallthrough
	default:
		println("Unknown")
	}
}

func testMySwitchCondition(x int) {
	// switch evaluation order
	// switch case will be evaluated from top to bottom
	switch {
	case x > 30:
		println("x is bigger than 30")
	case x > 20:
		println("x is bigger than 20")
	case x > 10:
		println("x is bigger than 10")
	default:
		println("x is less than 10")
	}
}

func main() {
	testIfShort1()
	x := 0
	if err := testIfShort2(&x); err != nil {
		panic("Error")
	}
	println(x, "should be 10")
	testIfElse()
	println(testMyLoopAndFunc(9), "should be", math.Sqrt(9))
	testMySwitch()
	testMySwitchCondition(15)
}
