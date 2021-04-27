package main

import (
	"math/rand"
	"testing"
	"time"
)

//go test -bench=Generate -benchmem .
//goos: windows
//goarch: amd64
//pkg: benmark
//BenchmarkGenerateWithCap-4            49          24674882 ns/op         8003592 B/op          1 allocs/op
//BenchmarkGenerate-4                   32          36220822 ns/op        45188358 B/op         40 allocs/op
//PASS
//ok      benmark 2.670s
// 预先设置好cap，内存消耗大大减少

func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func BenchmarkGenerateWithCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateWithCap(1000000)
	}
}

func BenchmarkGenerate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generate(1000000)
	}
}
