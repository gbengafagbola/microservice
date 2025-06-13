package main

import (
	"fmt"        // Package for formatted I/O (e.g., printing to console, writing to HTTP response)
	"net/http"   // Package for building HTTP servers and clients
	"time"       // Package for time-related functions (e.g., simulating delays)
)

// fetchData simulates fetching data from a backend service.
// It takes a serviceName (string) and a channel (ch chan string) to send results.
// This function runs as a goroutine to simulate concurrent operations.
func fetchData(serviceName string, ch chan string) {
	// Simulate network delay based on the length of the service name.
	// This helps demonstrate asynchronous behavior.
	time.Sleep(time.Duration(len(serviceName)) * 100 * time.Millisecond)

	// Send the simulated data string back through the channel.
	// The 'ch <-' syntax sends a value into the channel.
	ch <- fmt.Sprintf("Data from %s service", serviceName)
}

// aggregateHandler is the HTTP handler for the "/aggregate" endpoint.
// It demonstrates how to concurrently fetch data from multiple simulated services.
func aggregateHandler(w http.ResponseWriter, r *http.Request) {
	// Create a buffered channel of strings with a capacity of 3.
	// A buffered channel allows up to 3 values to be sent without a receiver being ready,
	// which is perfect for collecting results from a known number of goroutines.
	ch := make(chan string, 3)

	// Launch three goroutines, each simulating a call to a different backend service.
	// 'go' keyword starts a new concurrent execution (goroutine).
	go fetchData("UserService", ch)
	go fetchData("ProductService", ch)
	go fetchData("OrderService", ch)

	// Collect results from the channel.
	// We expect 3 responses, so we loop 3 times, receiving one value per iteration.
	// The '<-ch' syntax receives a value from the channel.
	var responses []string // Slice to store the collected responses
	for i := 0; i < 3; i++ {
		responses = append(responses, <-ch) // Append received data to the slice
	}

	// Format and send the aggregated data back as the HTTP response.
	fmt.Fprintf(w, "Aggregated Data: %v\n", responses)
}

// main is the entry point of the application.
func main() {
	// Register the aggregateHandler function to handle requests for the "/aggregate" path.
	http.HandleFunc("/aggregate", aggregateHandler)

	// Start the HTTP server on port 8080.
	// This function blocks, meaning it keeps the server running until it's explicitly stopped.
	fmt.Println("Server listening on :8080")
	err := http.ListenAndServe(":8080", nil) // 'nil' means use default HTTP server multiplexer
	if err != nil {
		// If there's an error starting the server (e.g., port already in use), print it.
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
