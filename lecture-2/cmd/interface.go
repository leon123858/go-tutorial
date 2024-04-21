package main

import (
	"fmt"
	oimg "image"
	"image/color"
	"io"
	"math"
	"reflect"
	"strings"
)

type operation struct {
	First  int
	Second int
	OpCode string
}

// can not function overloading, but can use variadic parameters
// use *operation to pass by reference, or can not change the value of the struct
func (op *operation) exec(opcode ...string) int {
	if len(opcode) > 0 {
		// convert interface to string
		op.OpCode = opcode[0]
	}
	var ret int
	switch op.OpCode {
	case "+":
		ret = op.First + op.Second
	case "-":
		ret = op.First - op.Second
	case "*":
		ret = op.First * op.Second
	case "/":
		ret = op.First / op.Second
	default:
		panic("wrong op code")
	}
	return ret
}

func testMethod() {
	op := operation{2, 5, "+"}
	println(op.OpCode, " ", op.exec())
	op.OpCode = "-"
	println(op.OpCode, " ", op.exec())
	op.OpCode = "*"
	println(op.OpCode, " ", op.exec())
	op.OpCode = "/"
	println(op.OpCode, " ", op.exec())

	op = operation{2, 5, "+"}
	println(op.OpCode, " ", op.exec("-"))
	println(op.OpCode, " ", op.exec("*"))
	println(op.OpCode, " ", op.exec("/"))
	println(op.OpCode, " ", op.exec("+"))
}

type buffer struct {
	data int
}

func (b *buffer) writeWithRef(data int) {
	b.data = data
}

func (b buffer) writeWithValue(data int) {
	b.data = data
}

func readWithRef(b *buffer) int {
	return b.data
}

func testPointerReceivers() {
	b := buffer{10}
	println("before writeWithRef: ", readWithRef(&b))
	b.writeWithRef(20)
	println("after writeWithRef: ", readWithRef(&b))
	// while methods with pointer receivers take either a value or a pointer as the receiver when they are called
	b.writeWithValue(30)
	println("after writeWithValue: ", readWithRef(&b))
}

type desktop string
type laptop string

//func (d desktop) write(data int) {
//	println("desktop write: ", data)
//}
//
//func (l laptop) write(data int) {
//	println("laptop write: ", data)
//}

func (d *desktop) write(data int) {
	println("desktop write: ", data)
}

func (l *laptop) write(data int) {
	println("laptop write: ", data)
}

func testInterface() {
	// do not exist pointer receiver and value receiver at the same time
	type writer interface {
		write(data int)
	}

	// all methods of interface use value receiver
	//var w writer
	//w = desktop("desktop")
	//w.write(10)
	//w = laptop("laptop")
	//w.write(20)

	// all methods of interface use pointer receiver (suggest!)
	var w writer
	d := desktop("desktop")
	w = &d
	w.write(10)
	l := laptop("laptop")
	w = &l
	w.write(20)
}

func (b *buffer) write(data int) {
	b.data = data
}

func (b *buffer) read() int {
	return b.data
}

func testInterfaceValue() {
	// can let interface as an argument of function
	type bufferInterface interface {
		write(data int)
		read() int
	}

	save := func(b *bufferInterface, data int) {
		(*b).write(data)
	}

	read := func(b *bufferInterface) int {
		return (*b).read()
	}

	var bi bufferInterface
	b := buffer{10}
	bi = &b
	save(&bi, 20)
	println("read: ", read(&bi))
}

func testEmptyInterface() {
	type baby struct {
		data int
	}
	// empty interface can store any type
	var i interface{}
	i = 10
	fmt.Printf("(%v, %T)\n", i, i)
	i = "hello"
	fmt.Printf("(%v, %T)\n", i, i)
	i = true
	fmt.Printf("(%v, %T)\n", i, i)
	i = baby{10}
	fmt.Printf("(%v, %T)\n", i, i)

	// reflect struct name
	if reflect.TypeOf(i) != reflect.TypeOf(baby{}) {
		panic("can not get interface type")
	}
	// switch by struct name
	switch v := i.(type) {
	case int:
		panic("int")
	case string:
		panic("string")
	case bool:
		panic("bool")
	case baby:
		println("baby: ", v.data)
	}
	// type assertion
	b, ok := i.(baby)
	if ok {
		println("baby: ", b.data)
	}
	_, ok = i.(int)
	if !ok {
		println("not int")
	}
}

type IPAddr [4]byte

// implement Stringer interface
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func testExerciseStringers() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

func testExerciseErrors(x float64) (float64, error) {
	if x < 0 {
		return 0, &customError{"negative number", 400}
	}
	return math.Sqrt(x), nil
}

type customError struct {
	Msg  string
	Code uint
}

func (e *customError) Error() string {
	return fmt.Sprintf("code: %v, msg: %v", e.Code, e.Msg)
}

func testExerciseReader() {
	r := MyReader{}
	b := make([]byte, 8)
	n, e := r.Read(b)
	if e != nil {
		panic(e)
	}
	println(n, " ", string(b))
}

func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

type MyReader struct{}

func testExerciseRot13Reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	b := make([]byte, len("Lbh penpxrq gur pbqr!"))
	n, e := r.Read(b)
	if e != nil {
		panic(e)
	}
	println(n, " ", string(b))
}

// as slice is reference type, so no need to return the slice
func (r *rot13Reader) Read(b []byte) (int, error) {
	// original read method
	n, e := r.r.Read(b)
	if e != nil {
		return n, e
	}
	// update the read data
	for i := range b {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] = (b[i]-'A'+13)%26 + 'A'
		} else if b[i] >= 'a' && b[i] <= 'z' {
			b[i] = (b[i]-'a'+13)%26 + 'a'
		}
	}
	return n, nil
}

type rot13Reader struct {
	r io.Reader
}

func testExerciseImages() {
	m := Image{}
	println(m.Bounds().Max.X, " ", m.Bounds().Max.Y)
	println(m.At(0, 0).RGBA())
}

type Image struct{}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() oimg.Rectangle {
	return oimg.Rect(0, 0, 50, 50)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{R: uint8(x), G: uint8(y), B: 255, A: 255}
}

func main() {
	testMethod()
	testPointerReceivers()
	testInterface()
	testInterfaceValue()
	testEmptyInterface()
	testExerciseStringers()
	if _, e := testExerciseErrors(-1); e != nil {
		println(e.Error())
	} else {
		panic("should get error")
	}
	if r, e := testExerciseErrors(9); e != nil {
		panic("should not get error")
	} else {
		println(r)
	}
	testExerciseReader()
	testExerciseRot13Reader()
	testExerciseImages()
}
