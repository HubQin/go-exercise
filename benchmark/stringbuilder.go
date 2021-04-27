package main

import (
	"bytes"
	"math/rand"
	"time"
)

const (
	data = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890,.-=/"
)

func init() {
	rand.Seed(time.Now().Unix()) // 设置随机种子
}

// GenRandString 生成n个随机字符的string
func GenRandString(n int) string {
	max := len(data)
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteByte(data[rand.Intn(max)])
	}

	return buf.String()
}

// GenRandBytes 生成n个随机字符的[]byte
func GenRandBytes(n int) []byte {
	max := len(data)
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = data[rand.Intn(max)]
	}

	return buf
}
