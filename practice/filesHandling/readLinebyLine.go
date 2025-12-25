package main

// Best for logs or large text files without loading everything.
// ✅ Efficient buffered reading.
// ✅ Handles large files without high memory use.

// func main() {
// 	file, err := os.Open("test.txt")
// 	if err != nil {
// 		log.Fatalf("failed to open file:%v", err) // cleaner error handling
// 	}

// 	defer file.Close()

// 	// Create a scanner
// 	scanner := bufio.NewScanner(file)

// 	// Read file line by line
// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}

// 	// Check for errors during scanning
// 	if err := scanner.Err(); err != nil {
// 		log.Fatalf("Error reading file: %v", err)
// 	}
// }
