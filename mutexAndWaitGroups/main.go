package main

import (
	"fmt"
	"sync"
)

func main() {

score := []int{0}

wg := &sync.WaitGroup{}
mx := &sync.Mutex{}

mr := &sync.RWMutex{}

wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("One")
		mx.Lock()
		score = append(score, 1)
		mx.Unlock()
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Two")
			mx.Lock()
		score = append(score, 2)
				mx.Unlock()
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Three")
		mr.RLock()
		fmt.Println(score)
		mr.RUnlock()
	}()

	wg.Wait()
	fmt.Println(score)

}