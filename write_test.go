package main

import "testing"

func BenchmarkWrite(b *testing.B) {
	for i := 0 ; i < b.N; i++ {
		initWrite()
	}
}
