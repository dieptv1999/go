package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	array := []int{}
	chunk1 := 0
	chunk2 := 0
	for i := 0; i < 20; i++ {
		array = append(array, i)
	}
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Printf("%v\n", array)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			chunk1 += array[i]
		}
		fmt.Printf("chunk1= %v\n", chunk1)
	}()
	go func() {
		defer wg.Done()
		for i := 10; i < 20; i++ {
			chunk2 += array[i]
		}
		fmt.Printf("chunk2= %v\n", chunk2)
	}()
	wg.Wait()
	fmt.Printf("chunk1 + chunk2= %v\n", chunk1+chunk2)
	time.Sleep(2 * time.Second)
}
//// Output
//[0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
//chunk2= 145
//chunk1= 45
//chunk1 + chunk2= 190