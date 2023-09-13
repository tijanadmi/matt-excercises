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

func validateOrders(input <-chan Order, output chan<- Order) {
	defer close(output) // Close the output channel when the function finishes

	for order := range input {
		// Simulating validation
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		// Validate order
		if order.Amount > 0 {
			fmt.Printf("Processed order ID: %d, validate amount: %.2f\n", order.ID, order.Amount)
			output <- order
		}
	}
}

func enrichOrders(input <-chan Order, output chan<- Order) {
	defer close(output) // Close the output channel when the function finishes

	for order := range input {
		// Simulating enrichment
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		// Enrich order
		order.Amount *= 1.1
		fmt.Printf("Processed order ID: %d, enrich amount: %.2f\n", order.ID, order.Amount)
		output <- order
	}
}

func calculateOrderValues(input <-chan Order, output chan<- Order) {
	defer close(output) // Close the output channel when the function finishes

	for order := range input {
		// Simulating calculation
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		// Perform calculation
		order.Amount += 5.0
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
