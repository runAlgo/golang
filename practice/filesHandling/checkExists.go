package main

// // Use os.Stat with os.IsNotExist
// // Safe per-check for read/write flows.

// func main() {
// 	_, err := os.Stat("maybe.txt")
// 	if os.IsNotExist(err) {
// 		fmt.Println("File does NOT exist:", err)
// 	} else if err != nil {
// 		fmt.Println("Error checking:", err)
// 	} else {
// 		fmt.Println("File exists")
// 	}
// }
