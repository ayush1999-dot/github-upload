package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	//now := time.Now()
	/*go task1()
	go task2()
	go task3()
	go task4()
	time.Sleep(1 * time.Second)
	fmt.Println("time elaspsed", time.Since(now))*/
	wtgrp(wg)

}

/*
func task1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task 1")
}
func task2() {

	time.Sleep(200 * time.Millisecond)
	fmt.Println("task 2")
}
func task3() {

	fmt.Println("task 3")
}
func task4() {

	time.Sleep(400 * time.Millisecond)
	fmt.Println("task 4")
}
*/
