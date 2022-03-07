package main

import (
	"fmt"
	"sync"
)

func wtgrp(wg sync.WaitGroup) {

	wg.Add(1)
	go func() {
		tp()
		wg.Done()
	}()
	wg.Wait()
}
func tp() {
	fmt.Println()
	fmt.Println(" wait grp executed")
	fmt.Println()
}
