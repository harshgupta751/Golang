package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var signals = []string{
	"test",
}
var mx sync.Mutex
func main() {

websites := []string{
	"https://google.com",
	"https://codehurdle.com",
	"https://cpbytekiet.netlify.app/",
	"https://cpbytestudentportal.netlify.app/login",
}

var wg sync.WaitGroup

for _, website := range websites {
		wg.Add(1)
	go getStatus(website, &wg)
} 

wg.Wait()
fmt.Println(signals)


// go greet("Hello")
// greet("World")


}

// func greet(s string) {
	// for i := 0; i <= 5; i++ {
	// 	time.Sleep(3*time.Second)
	// 	fmt.Println(s)
	// }

// }

func getStatus(website string, wg *sync.WaitGroup){

defer wg.Done()
res, err := http.Get(website)
if err!=nil {
	log.Fatal(err)
}

mx.Lock()
signals= append(signals, website)
mx.Unlock()
fmt.Printf("StatusCode by pinging %s is : %v\n", website, res.StatusCode)


}