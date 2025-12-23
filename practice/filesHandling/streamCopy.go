package main

// **** Stream Copy (Large File Copy) ****
// Uses io.Copy to stream data efficiently:
// Stream data (memory-friendly)
// Ideal for backups, uploads, or large media.

// func main() {
// 	src, err := os.Open("biginput.bin")
// 	if err != nil {
// 		fmt.Println("Open error:", err)
// 		return
// 	}
// 	defer src.Close()

// 	dst, err := os.Create("bigoutput.bin")
// 	if err != nil {
// 		fmt.Println("Create error:", err)
// 		return
// 	}
// 	defer dst.Close()

// 	_, e := io.Copy(dst, src)
// 	if e != nil {
// 		fmt.Println("Copy error:", e)
// 		return
// 	}

// 	fmt.Println("Copy Completed!")
// }
