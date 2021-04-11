package main

import (
	"fmt"
	"sync"
	"time"
)

// https://gobyexample.com/waitgroups
func main() {
	fmt.Println("normal example...")
	var wg sync.WaitGroup

	for i := 1; i<=5;i++  {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("Worker %d starting\n", i)
			time.Sleep(time.Second)
			fmt.Printf("Worker %d done\n", i)
		}(i)
	}
	wg.Wait()

	fmt.Println("closure buggy example...")

	for i := 1; i<=5;i++  {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("***Worker %d starting\n", i)
			time.Sleep(time.Second)
			fmt.Printf("***Worker %d done\n", i)
		}()
	}
	wg.Wait()
}