package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("normal example...")
	for i := 1; i <= 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(1 * time.Second)


	fmt.Println("normal example 2...")
	for i := 1; i <= 5; i++ {
		i := i
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(1 * time.Second)


	fmt.Println("closure buggy example...")
	for i := 1; i <= 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(1 * time.Second)
}
