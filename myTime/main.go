package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate:= time.Date(2010, time.June, 25, 15, 38, 45, 2, time.UTC);

	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}
