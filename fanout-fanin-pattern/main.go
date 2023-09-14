package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Amount float64
}

func validateOrders(input <-chan Order, output chan<- Order) {
	defer close(output) // Close the output channel when the function finishes

	// Create a channel to receive validated orders
	validatedOrders := make(chan Order)

	// Specify the number of worker goroutines to use for fan-out
	workerCount := 3

	// Fan-out: Start multiple worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(input <-chan Order) {
			defer wg.Done()
			for order := range input {
				// Simulating validation
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

				// Validate order
				if order.Amount > 0 {
					validatedOrders <- order
				}
			}
		}(input)
	}

	// Fan-in: Collect the validated orders from the workers
	go func() {
		wg.Wait()
		close(validatedOrders)
	}()

	for order := range validatedOrders {
		output <- order
	}
}

func enrichOrders(input <-chan Order, output chan<- Order) {
	defer close(output) // Close the output channel when the function finishes

	// Create a channel to receive enriched orders
	enrichedOrders := make(chan Order)

	// Specify the number of worker goroutines to use for fan-out
	workerCount := 3

	// Fan-out: Start multiple worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(input <-chan Order) {
			defer wg.Done()

			for order := range input {
				// Simulating calculation
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

				// Enrich order
				order.Amount *= 1.1
				enrichedOrders <- order
			}
		}(input)
	}
	// Fan-in: Collect the enriched orders from the workers
	go func() {
		wg.Wait()
		close(enrichedOrders)
	}()

	for order := range enrichedOrders {
		fmt.Printf("Processed order ID: %d, enrich amount: %.2f\n", order.ID, order.Amount)
		output <- order
	}

}

func calculateOrderValues(input <-chan Order, output chan<- Order) {
	defer close(output) // Close the output channel when the function finishes

	// Create a channel to receive validated orders
	calculatedOrders := make(chan Order)

	// Specify the number of worker goroutines to use for fan-out
	workerCount := 3

	// Fan-out: Start multiple worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(input <-chan Order) {
			defer wg.Done()
			for order := range input {
				// Simulating calculation
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

				// Perform calculation
				order.Amount += 5.0
				calculatedOrders <- order
			}
		}(input)
	}

	// Fan-in: Collect the validated orders from the workers
	go func() {
		wg.Wait()
		close(calculatedOrders)
	}()
	for order := range calculatedOrders {
		fmt.Printf("Processed order ID: %d, calculate amount: %.2f\n", order.ID, order.Amount)
		output <- order
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create channels to connect the stages
	orders := make(chan Order)
	validOrders := make(chan Order)
	enrichedOrders := make(chan Order)
	calculatedOrders := make(chan Order)

	// Start the stages concurrently
	go validateOrders(orders, validOrders)
	go enrichOrders(validOrders, enrichedOrders)
	go calculateOrderValues(enrichedOrders, calculatedOrders)

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
