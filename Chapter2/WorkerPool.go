package main

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}

func main() {
	//var wg sync.WaitGroup
	//for i := 1; i <= 100; i++ {
	//wg.Add(1)
	//go func(j int) {
	//defer wg.Done()
	//fmt.Println(j)
	//}(i)
	//}
	//wg.Wait()

	ports := make(chan int, 100)
	var wg sync.WaitGroup
	//The workers
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}
	//Sender of information
	for i := 1; i < 80; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}
