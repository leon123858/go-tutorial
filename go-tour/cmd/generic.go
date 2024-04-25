package main

import (
	"errors"
	"golang.org/x/exp/constraints"
	"strconv"
)

// [T comparable] is a type constraint, which means T must be comparable
func findItem[T comparable](s []T, item T) int {
	for i, v := range s {
		if v == item {
			return i
		}
	}
	return -1
}

// constraints library provides some predefined type constraints
func findItem2[T constraints.Ordered](s []T, item T) int {
	for i, v := range s {
		if v == item {
			return i
		}
	}
	return -1
}

// ListItem List represents a singly-linked list that holds
// values of any type.
type ListItem[T any] struct {
	next *ListItem[T]
	val  T
}

// NewListItem NewList creates a new list with the given value.
func NewListItem[T any](val T) *ListItem[T] {
	return &ListItem[T]{val: val}
}

// Append adds a new value to the end of the list.
func (l *ListItem[T]) Append(val T) {
	for cur := l; ; cur = cur.next {
		if cur.next == nil {
			cur.next = NewListItem(val)
			return
		}
	}
}

// printAll prints all values in the list and applies the given function to each value.
func (l *ListItem[T]) printAll(printFunc func(T) T) {
	for cur := l; cur != nil; cur = cur.next {
		println(printFunc(cur.val))
	}
}

// multiple type constraints
func testMultipleConstraints[T int | float64 | string](x T) {
	println(x)
}

// many type constraints
func testManyConstraints[T error, T2 constraints.Unsigned](x T, x2 T2) {
	println(x.Error())
	println(x2 * 2)
}

type customError2 struct {
	Msg  string
	Code uint
}

func (e *customError2) Error() string {
	return e.Msg + " " + strconv.Itoa(int(e.Code))
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	println(findItem(s, 3))
	println(findItem(s, 6))

	s2 := []string{"a", "b", "c", "d", "e"}
	println(findItem2(s2, "c"))
	println(findItem2(s2, "f"))

	l := NewListItem[int](1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.printAll(func(v int) int {
		return v * 2
	})

	testMultipleConstraints(10)
	testMultipleConstraints(10.5)
	testMultipleConstraints("hello")

	testManyConstraints(&customError2{"error", 400}, uint(10))
	testManyConstraints(errors.New("cool"), uint(10))
}
