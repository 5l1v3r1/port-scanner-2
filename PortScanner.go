package main

import (
	"fmt"
	"net"
	"sort"
)

var input string

// Function that holds host to be scanned.
func host() {
	fmt.Println("Enter the IP/Host you want scanned: ")
	fmt.Scanln(&input)
}

/* Function that creates two channels and loops through
host and port numbers. */
func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf(input+":%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	/* You can increase the buffer > 100 which will improve speed, but
	will reduce the reliability of the results. */
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	host()
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
