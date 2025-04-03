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
func worker(wg *sync.WaitGroup, tasks chan string, dialer net.Dialer, openPorts *[]string, mu *sync.Mutex, endPort int) {
	defer wg.Done()
	maxRetries := 3
	for addr := range tasks {
		var success bool
		for i := range maxRetries {
			// Extract the port number part from the addr string
			portStr := strings.Split(addr, ":")[1]

			// Print a message to show which port is being scanned
			fmt.Printf("Scanning port %s/%s\n", portStr, strconv.Itoa(endPort))

			conn, err := dialer.Dial("tcp", addr)
			if err == nil {
				// This prevents hanging forever if no data is sent by the server
				conn.SetReadDeadline(time.Now().Add(2 * time.Second))

				// Create a space in memory to store the response
				buffer := make([]byte, 1024)

				// We attempt to read the data from the server
				n, err := conn.Read(buffer)
				var banner string
				if err == nil && n > 0 {
					// Convert the bytes data into string
					banner = string(buffer[:n])
					// Display the banner string
					fmt.Println("----------------------------------------------------------")
					fmt.Printf("Connection to %s was successful\n", addr)
					fmt.Printf("Banner: %s", banner)
					fmt.Println("----------------------------------------------------------")
				} else {
					fmt.Printf("Connection to %s was successful (no banner)\n", addr)
				}

				conn.Close()

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
	// target := flag.String("target", "scanme.nmap.org", "Target IP Address Or Hostname")
	targets := flag.String("targets", "scanme.nmap.org", "List Of Targets Separated By Commas (e.g., scanme.nmap.org,example.com)")
	startPort := flag.Int("start-port", 1, "Starting Port Number")
	endPort := flag.Int("end-port", 1024, "Ending Port Number")
	workers := flag.Int("workers", 100, "Number Of Workers")
	timeout := flag.Int("timeout", 5, "Timeout In Seconds")

	// Parse the flags from command line
	flag.Parse()

	// Split the targets strings
	targetList := strings.Split(*targets, ",")

	dialer := net.Dialer{
		// We timeout based on the amount of seconds specified by the user
		Timeout: time.Duration(*timeout) * time.Second,
	}

	// We create the number of workers specified by the user
	for i := 1; i <= *workers; i++ {
		wg.Add(1)
		go worker(&wg, tasks, dialer, &openPorts, &mu, *endPort)
	}

	// We loop through the list of targets provided by the user
	for _, target := range targetList {
		// We loop between ports specified by the user for the current target
		for p := *startPort; p <= *endPort; p++ {
			port := strconv.Itoa(p)
			address := net.JoinHostPort(target, port)
			tasks <- address
		}
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
