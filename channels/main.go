package main


import (
	"fmt"
	"sync"
)

func main(){
// ch := make(chan int)

// ch <- 5

// ch := make(chan int, 2)

// ch <- 5
// ch <- 10
// ch <- 15

// fmt.Println(<- ch)
// fmt.Println(<- ch)
// fmt.Println(<- ch)



ch := make(chan int)

wg := &sync.WaitGroup{}

wg.Add(2)

go func(ch <-chan int){
	defer wg.Done()
	val, isChannelOpen := <- ch
	fmt.Println(val)
	fmt.Println(isChannelOpen)
}(ch)

go func(ch chan<- int){
	defer wg.Done()
	ch <- 5
	close(ch)
	
}(ch)


wg.Wait()

}