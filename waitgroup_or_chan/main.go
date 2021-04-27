package main

import (
	"fmt"
	"sync"
	"time"
)

// https://gobyexample.com/waitgroups
func main() {
	fmt.Println("waitgroup example...")
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("Worker %d starting\n", i)
			time.Sleep(time.Second)
			fmt.Printf("Worker %d done\n", i)
		}(i)
	}
	wg.Wait()

	// not good
	fmt.Println("chan example...")
	c := make(chan bool, 6)
	for i := 1; i <= 5; i++ {
		i := i
		go func() {
			fmt.Printf("Worker %d starting\n", i)
			time.Sleep(time.Second)
			fmt.Printf("Worker %d done\n", i)
			c <- true
		}()
	}
	for i := 1; i <= 5; i++ {
		<-c
	}
}
