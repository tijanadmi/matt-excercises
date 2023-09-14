package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Order struct {
	ID     int
	Amount float64
}

func validateOrders(done <-chan struct{}, input <-chan Order, output chan<- Order) {
	defer close(output) // Close the output channel when the function finishes

	for {
		select {
		case <-done:
			return
		case order, ok := <-input:
			if !ok {
				return
			}

			// Simulating validation
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

			// Validate order
			if order.Amount > 0 {
				output <- order
			}
		}
	}
}

func enrichOrders(done <-chan struct{}, input <-chan Order, output chan<- Order) {
	defer close(output) // Close the output channel when the function finishes

	for {
		select {
		case <-done:
			return
		case order, ok := <-input:
			if !ok {
				return
			}

			// Simulating enrichment
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

			// Enrich order
			order.Amount *= 1.1
			output <- order
		}
	}
}

func calculateOrderValues(done <-chan struct{}, input <-chan Order, output chan<- Order) {
	defer close(output) // Signal the wait group that the processing is done for this goroutine

	for {
		select {
		case <-done:
			return
		case order, ok := <-input:
			if !ok {
				return
			}

			// Simulating calculation
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

			// Perform calculation
			order.Amount += 5.0
			output <- order
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create done channel to signal termination
	done := make(chan struct{})
	defer close(done)

	// Create channels to connect the stages
	orders := make(chan Order)
	validOrders := make(chan Order)
	enrichedOrders := make(chan Order)
	calculatedOrders := make(chan Order)

	// Start the stages concurrently
	go validateOrders(done, orders, validOrders)
	go enrichOrders(done, validOrders, enrichedOrders)
	go calculateOrderValues(done, enrichedOrders, calculatedOrders)

	// Generate sample orders
	go func() {
		for i := 1; i <= 10; i++ {
			order := Order{
				ID:     i,
				Amount: float64(i * 10),
			}
			orders <- order
		}
		close(orders)
	}()

	// Receive the final output from the pipeline
	for order := range calculatedOrders {
		fmt.Printf("Processed order ID: %d, Final amount: %.2f\n", order.ID, order.Amount)
	}
}
