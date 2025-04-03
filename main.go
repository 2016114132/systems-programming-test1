package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Modify the worker function to accept
func worker(wg *sync.WaitGroup, tasks chan string, dialer net.Dialer, openPorts *[]string, mu *sync.Mutex) {
	defer wg.Done()
	maxRetries := 3
	for addr := range tasks {
		var success bool
		for i := range maxRetries {
			conn, err := dialer.Dial("tcp", addr)
			if err == nil {
				conn.Close()
				fmt.Printf("Connection to %s was successful\n", addr)

				// Extract the port number part from the addr string
				portStr := strings.Split(addr, ":")[1]

				// Multiple workers might try to add to the array at the same time.
				// This can cause the application to crash or produce incorrect results.
				// To mitigate this, we use a concept called "mutual exclusion".
				// This lock the array while one goroutine is adding to it, then unlocks it after it's done.
				// This prevents other goroutines to update it while its being used by another goroutine.
				mu.Lock()
				*openPorts = append(*openPorts, portStr)
				mu.Unlock()

				success = true
				break
			}
			backoff := time.Duration(1<<i) * time.Second
			fmt.Printf("Attempt %d to %s failed. Waiting %v...\n", i+1, addr, backoff)
			time.Sleep(backoff)
		}
		if !success {
			fmt.Printf("Failed to connect to %s after %d attempts\n", addr, maxRetries)
		}
	}
}

func main() {
	// Keeps track of how the time to know how long the operation takes
	start := time.Now()

	var wg sync.WaitGroup
	tasks := make(chan string, 100)

	// Array to keep track of open ports
	var openPorts []string

	// A mutual exclusion needs to be used to make sure only one goroutine can access a variable at a time to avoid conflicts
	var mu sync.Mutex

	// Define flags with default values
	// Flags allow users to pass values when they run the program
	target := flag.String("target", "scanme.nmap.org", "Target IP Address Or Hostname")
	startPort := flag.Int("start-port", 1, "Starting Port Number")
	endPort := flag.Int("end-port", 1024, "Ending Port Number")
	workers := flag.Int("workers", 100, "Number Of Workers")
	timeout := flag.Int("timeout", 5, "Timeout In Seconds")

	// Parse the flags from command line
	flag.Parse()

	dialer := net.Dialer{
		// We timeout based on the amount of seconds specified by the user
		Timeout: time.Duration(*timeout) * time.Second,
	}

	// We create the number of workers specified by the user
	for i := 1; i <= *workers; i++ {
		wg.Add(1)
		go worker(&wg, tasks, dialer, &openPorts, &mu)
	}

	// We loop between ports specified by the user
	for p := *startPort; p <= *endPort; p++ {
		port := strconv.Itoa(p)
		address := net.JoinHostPort(*target, port)
		tasks <- address
	}
	close(tasks)
	wg.Wait()

	// Display the scan summary
	duration := time.Since(start)
	fmt.Println("\n---------------------------------")
	fmt.Println("Scan Summary:")
	fmt.Printf("\nOpen ports: %d\n", len(openPorts))
	fmt.Printf("Total ports scanned: %d\n", *endPort-*startPort+1)
	fmt.Printf("Time taken: %v\n", duration)
	fmt.Println("---------------------------------")
}
