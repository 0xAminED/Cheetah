package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

// List of HTTP methods to test
var methods = []string{
	"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD",
}

func checkURL(target, path string, method string, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()

	// Prepare the full URL by concatenating the target URL and the path
	url := fmt.Sprintf("%s/%s", target, path)

	// Create an HTTP client to make requests with custom methods
	client := &http.Client{
		Timeout: 10 * time.Second, // Set a timeout for each request
	}

	// Prepare the HTTP request
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		ch <- fmt.Sprintf("Error: %v | Method: %s | URL: %s", err, method, url)
		return
	}

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		ch <- fmt.Sprintf("Error: %v | Method: %s | URL: %s", err, method, url)
		return
	}
	defer resp.Body.Close()

	// Send the status code, headers, and method to the channel
	ch <- fmt.Sprintf("Method: %s | URL: %s\nStatus Code: %d\nHeaders: %v\n", method, url, resp.StatusCode, resp.Header)
}

func processFile(filePath string) []string {
	var paths []string
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Add each line (path) from the file to the paths slice
		paths = append(paths, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}

	return paths
}

func main() {
	// Command-line flags for target URL and input file
	target := flag.String("u", "", "Target URL (e.g., https://example.com)")
	inputFile := flag.String("i", "", "File containing the list of directories to test")

	// Parse the command-line flags
	flag.Parse()

	// Check if the URL and input file are provided
	if *target == "" || *inputFile == "" {
		log.Fatal("Both target URL (-u) and input file (-i) are required.")
	}

	// Read the list of paths from the input file
	paths := processFile(*inputFile)

	// Create a wait group and a channel to handle concurrency
	var wg sync.WaitGroup
	ch := make(chan string, len(paths))

	// Set the maximum number of concurrent requests (for throttling if needed)
	maxConcurrency := 50
	sem := make(chan struct{}, maxConcurrency)

	// Start the directory discovery for each path with method tampering
	startTime := time.Now()
	for _, path := range paths {
		wg.Add(1)
		sem <- struct{}{} // acquire a slot in the semaphore
		for _, method := range methods {
			go func(p, m string) {
				defer func() { <-sem }() // release the slot in the semaphore
				checkURL(*target, p, m, &wg, ch)
			}(path, method)
		}
	}

	// Close the channel when all work is done
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Collect and display the results
	for result := range ch {
		fmt.Println(result)
	}

	// Display the total time taken for the operation
	duration := time.Since(startTime)
	fmt.Printf("\nDirectory discovery completed in %v\n", duration)
}
