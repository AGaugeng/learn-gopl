package main

import (
	"testing"
	"time"
)

// 15976             71770 ns/op             432 B/op          6 allocs/op
func BenchmarkEcho1(t *testing.B) {
	now := time.Now()
	for i := 0; i < t.N; i++ {
		echo1()
	}
	t.Log(time.Since(now))
}

// 15380             85977 ns/op             432 B/op          6 allocs/op
func BenchmarkEcho2(t *testing.B) {
	now := time.Now()
	for i := 0; i < t.N; i++ {
		echo2()
	}
	t.Log(time.Since(now))
}
