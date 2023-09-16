package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, result chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			result <- 0
			continue
		}
		conn.Close()
		result <- p
	}
}

func main() {
	ports := make(chan int, 50)
	result := make(chan int)
	var openPorts []int
	for i := 1; i < cap(ports); i++ {
		go worker(ports, result)
	}
	//We have 50 goroutines waiting to read
	go func() {
		for i := 1; i < 100; i++ {
			ports <- i
		}
	}()
	//We send a 100 values that get consumed by the workers really quickly
	//so the ports channel does not get filled up (because a part of it gets emptied before reaching capacity)

	for i := 1; i < 100; i++ {
		open := <-result
		if open != 0 {
			openPorts = append(openPorts, open)
		}
	}
	close(ports)
	close(result)
	sort.Ints(openPorts)

	for _, openPort := range openPorts {
		fmt.Printf("%d open \n", openPort)
	}
}
