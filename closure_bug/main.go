package main

import (
	"fmt"
	"time"
)

// https://yourbasic.org/golang/gotcha-data-race-closure/
func main() {
	// 错误示例：
	fmt.Println("closure buggy example...")
	for i := 1; i <= 5; i++ {
		// 每个goroutine共享一个变量，goroutine还没开始的时候，i已经变成了6
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(1 * time.Second)

	// 正确示例1：
	fmt.Println("normal example...")
	for i := 1; i <= 5; i++ {
		go func(i int) { // 使用局部变量
			fmt.Println(i)
		}(i)
	}
	time.Sleep(1 * time.Second)

	// 正确示例2：
	fmt.Println("normal example 2...")
	for i := 1; i <= 5; i++ {
		i := i // 为每个闭包创建一个变量
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(1 * time.Second)

}
