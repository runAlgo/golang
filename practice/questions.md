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

## 3 Advanced (Underlying Array and Sub-Slicing)

## Scenario
You are implementing a component that manages a large, pre-allocated buffer (an underlying array) to store the log entries. Your component exposes several smaller "views" (slices) of this buffer to different processing units. You need to understand how modifications in one slice affect others due to the shared underlying array.

## Task
Write a Go function named ```SharedBufferTest``` that demonstrates the relationship between slices and their underlying array:

1. Create an initial slice of integers, ```buffer```, with a length and capacity of 5, initialized with values [1, 2, 3, 4, 5]
2. Create a new slice, ```viewA```, by taking a sub-slice of ```buffer``` from index 1(inclusive) to index 4(exclusive).
3. Modify the first element of ```viewA``` to ```99```.
4. Create a third slice, ```viewB```, by appending a new element (```100```) to the original ```buffer```.

The function should print the final state of all three slices (buffer, viewA, viewB) and explain why ```viewA```s modification affected ```buffer```.

## Constraints

* Use the ```make``` function to initialize the starting ```buffer``` slice.
* Print the slices in the order they are created/modified.

## Expected Result Explanation Focus
You should be able to explain that the modification in step 3 affects the original ```buffer``` because ```viewA``` and ```buffer``` are over the same underlying array data. The append operation in step 4 will likely create a new underlying array for ```viewB``` if the original capacity is exceeded, demonstrating when the sharing stops.