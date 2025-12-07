# Practice Questions

## 1 Basic (Filtering Data)

### Scenario

You are processing a list of customer transaction amounts. You need to create a new slice containing only the transactions that exceed a specifice threashold.

### Task

Write a Go function named ```FilterHighValueTransactions``` that accepts

1. A slice of floating-point numbers (```[]float64```) representing transaction amounts.
2. A single ```float64``` representing the minimum threshold.

The function should return a new slice containing only the amounts greater than the threshold. The original slice must not be modified.

### Contraints

* Use the append function to build the new slice.
* The function signature must be:
  
  ```func FilterHighValueTransactions(transactions []float64, threshold float64) []float64```

## Example Input/Output

Input: transactions = [10.50, 99.99, 500.00, 25.00, 1000.00]
threshold = 100.00
Output: [500.00, 1000.00]

## 2 Intermediate (In-Place Modification)

### Scenario
You are maintaining a queue of pending tasks represented by a slice of strings. When a task is completed, it must be removed from the slice. You need to do this efficiently without allocating a new, smaller slice every time.

### Task
Write a Go function named ```RemoveTaskAtIndex``` that accepts:

1. A slice of strings (```[]string```) representing the task queue.
2. An integer (```int```) representing the index of the task to remove.

## Constraints

* The removal must be done by copying the elements after the removed item over the gap, then slicing off the last element. This is the common, efficient way to remove from the middle of a slice in Go.
* The function signature must be:
  ```func RemoveTaskAtIndex(tasks []string, index int) []string```

## Example Input/Output

Input: tasks = ["A", "B", "C", "D", "E"]
index = 2 (Remove "C")
Output: ["A", "B", "D", "E"]