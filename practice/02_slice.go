package main

import (
	"fmt"
)

// RemoveTaskAtIndex modifies the slice in-place by removing the element
// at the specified index and return the truncated slice.
func RemoveTaskAtIndex(tasks []string, index int) []string {
	// 1. Handle out-of-bounds index (important safety check).
	if index < 0 || index >= len(tasks) {
		return  tasks
	}
	// 2. The core removal technique:
	// - copy(tasks[index:], tasks[index+1:])
	// - This copies the elements from index+1 onwards
	// 	 (the segment we want to keep) over the element at 'index'
	//   (the one we want to remove).
	// - The first arguent it the destination, the second is the source.
	copy(tasks[index:], tasks[index+1:])
	// Before: {"A", "B", "C", "D", "E"}
	// After: {"A", "B", "D", "E", "E"}
	fmt.Println("After Copy:",tasks)
	// 3. Return a slice that is one element shorter.
	// - We slice off the last element, which is now a duplicate
	//   of the one immediately precending it (due to the copy operation).
	// - tasks[:len(tasks)-1] return a slice from the beginning up to the 
	//   second-to-last element.
	return tasks[:len(tasks)-1]
}

// func main() {
// 	tasks := []string{"A", "B", "C", "D", "E"}
// 	indexToRemove := 2 // Index of "C"

// 	fmt.Println("Before Removel:", tasks) // Output: [A B C D E]

// 	tasks = RemoveTaskAtIndex(tasks, indexToRemove)

// 	fmt.Println("After Removal:", tasks) // Output: [A B D E]

// }
