package main

import "testing"

func BenchmarkQuote(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := getQuote()
		if x == "" {
			panic("no quotes, please add some quotes first.")
		}
	}
}
