package main

// *** Buffered Writer (High Throughput Writes)
// Better when writing many small chunks:
// Uses a buffer to reduce system calls.
// Great for logs or batch output.

// func main() {
// 	file, err := os.Create("buffered.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer file.Close()

// 	writer := bufio.NewWriter(file)

// 	_, e := writer.WriteString("Line one\n")
// 	if e != nil {
// 		fmt.Println("Error writing:", err)
// 		return
// 	}
// 	_, er := writer.WriteString("Line two\n")
// 	if er != nil {
// 		fmt.Println("Error writing:", er)
// 		return
// 	}

// 	writer.Flush()// Ensure all buffered data is written
// 	fmt.Println("Buffered write complete")

// }
