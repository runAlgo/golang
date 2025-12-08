package main

import "fmt"

// SharedBufferTest demonstrates how slices share an underlying array.
func SharedBufferTest() {
	// 1. Create initial slice (underlying array size is 5)
	buffer := make([]int, 5, 5)
	for i := 0; i < 5; i++ {
		buffer[i] = i + 1; // [1, 2, 3, 4, 5]
	}

	fmt.Println("--- Initial State ---")
	fmt.Printf("Buffer: %v (Len: %d, Cap: %d)\n", buffer, len(buffer), cap(buffer))

	// 2. Create viewA (sub-slice of buffer from index 1 to 3)
	//    viewA starts at the element '2' in the underlying array.
	viewA := buffer[1:4]
	fmt.Printf("viewA: %v (Len: %d, Cap: %d)\n", viewA, len(viewA), cap(viewA)) // [2, 3, 4], Cap is 4 (remaining elements in buffer)

	// 3. Modify the first element of viewA (which is the element '2')
	//    Since viewA shares the underlying array with 'buffer', this midification
	//    will affect 'buffer'
	viewA[0] = 99
	fmt.Println("\n-- After Modification: viewA[0] = 99 ---")
	fmt.Printf("Buffer: %v\n", buffer) // Output: [1, 99, 3, 4, 5] - MODIFIED!
	fmt.Printf("ViewA: %v\n", viewA) // Output: [99, 3, 4]


	// 4. Create viewB by APPENDING to the ORIGINAL 'buffer'.
	//    Since buffer has Len=5 and Cap=5 (fully utilized), appending '100'
	//    will cause Go to allocate a NEW, larger underlying array
	//    and copy the contents of 'buffer' into it. The sharing is broken.
	viewB := append(buffer, 100)

	// Now, if we modify viewA again, it will only affect the original 'buffer'
	// but NOT the new 'viewB'.

	viewA[1] = 77

	fmt.Println("\n-- After Append and Second Modification: view[1] = 77 ---")
	fmt.Printf("Buffer: %v\n", buffer) // Output: [1, 99, 77, 4, 5] - MODIFIED!
	fmt.Printf("ViewA: %v\n", viewA) // Output: [99, 77, 4]
	fmt.Printf("ViewB: %v (Len: %d, Cap: %d)\n", viewB, len(viewB), cap(viewB)) // Output: [1, 99, 77, 4, 5, 100] - UNMODIFIED by viewA[1]=77 because it has a new underlying array.
}

func main() {
	SharedBufferTest()
}