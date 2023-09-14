package main

import (
	"fmt"
	"time"
)

func tee(source <-chan int, destinations ...chan<- int) {
	go func() {
		for value := range source {
			for _, dest := range destinations {
				dest <- value
			}
		}
		for _, dest := range destinations {
			close(dest)
		}
	}()
}

func main() {
	sensorData := make(chan int)
	realTimeAnalysis := make(chan int)
	storage := make(chan int)

	// Use tee function to split the source channel into two destination channels
	tee(sensorData, realTimeAnalysis, storage)

	// Send values to the source channel
	go func() {
		for i := 1; i <= 10; i++ {
			sensorData <- i
		}
		close(sensorData)
	}()

	// Consume values from the destination channels
	go func() {
		for value := range realTimeAnalysis {
			fmt.Println("Performing real-time analysis on data:", value)
		}
	}()
	go func() {
		for value := range storage {
			fmt.Println("Storing data:", value)
		}
	}()

	// Wait for goroutines to finish
	time.Sleep(time.Second)
}
