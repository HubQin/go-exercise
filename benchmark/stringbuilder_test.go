package main

import (
	"bytes"
	"strings"
	"testing"
)

// go test -bench=Concat -benchmem .
// BenchmarkSpliceAddString10 测试使用 += 拼接N次长度为10的字符串
func BenchmarkConcatSpliceAddString10(b *testing.B) {
	s := ""
	for i := 0; i < b.N; i++ {
		s += GenRandString(10)
	}
}

// BenchmarkSpliceBuilderString10 测试使用strings.Builder拼接N次长度为10的字符串
func BenchmarkConcatSpliceBuilderString10(b *testing.B) {
	var builder strings.Builder
	for i := 0; i < b.N; i++ {
		builder.WriteString(GenRandString(10))
	}
}

// BenchmarkSpliceBufferString10 测试使用bytes.Buffer拼接N次长度为10的字符串
func BenchmarkConcatSpliceBufferString10(b *testing.B) {
	var buff bytes.Buffer
	for i := 0; i < b.N; i++ {
		buff.WriteString(GenRandString(10))
	}
}

// BenchmarkSpliceBufferByte10 测试使用bytes.Buffer拼接N次长度为10的[]byte
func BenchmarkConcatSpliceBufferByte10(b *testing.B) {
	var buff bytes.Buffer
	for i := 0; i < b.N; i++ {
		buff.Write(GenRandBytes(10))
	}
}

// BenchmarkSpliceBuilderByte10 测试使用string.Builder拼接N次长度为10的[]byte
func BenchmarkConcatSpliceBuilderByte10(b *testing.B) {
	var builder strings.Builder
	for i := 0; i < b.N; i++ {
		builder.Write(GenRandBytes(10))
	}
}
