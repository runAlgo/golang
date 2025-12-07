package main

import "fmt"

// Practice Question 1: Basic (Filtering Data)

// FilterHighValueTransactions returns a new slice containing only transactions
// that exceed the specified threshold.
func FilterHighValueTransactions(transactions []float64, threshold float64) []float64 {
	// 1. Create a new slice with zero length but with an initial capacity
	// based on the input size. This is a common performance optimization
	// to minimize reallocations.

	filtered := make([]float64, 0, len(transactions))

	// 2. Iterate through the input slice.
	for _, amount := range transactions {
		// 3. Check the condition.
		if amount > threshold {
			// 4. Append the high-value transaction to the new slice.
			filtered = append(filtered, amount)
		}
	}
	return filtered
}


func main() {
	var transactions = []float64{10.50, 99.99, 500.00, 25.00, 1000.00} 
	threshold := 100.00
	highValue := FilterHighValueTransactions(transactions, threshold)
	fmt.Println(highValue)
}