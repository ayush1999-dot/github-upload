package main

import (
	"fmt"
)

func reciver(ch <-chan int) { //recieve
	fmt.Println(<-ch)

}
func send(ch chan<- int) { //send
	ch <- 12344

}
func main() {

	ch := make(chan int)
	go send(ch)
	go reciver(ch)
	ch <- 42

	fmt.Println(<-ch)
	//we cannot write the function recive below the declared  channel
	//beacuse decared channel blocks the flow of statement so it never reaches the func
}
