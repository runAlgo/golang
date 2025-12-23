// package main

// import (
// 	"log"
// 	"os"
// )

// func main() {
// 	f, err := os.OpenFile("example.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	_, err = f.WriteString("Appended log line\n")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

package main

// func main() {
// 	f, err := os.OpenFile("example.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	n, e := f.WriteString("\nMe Appended data")
// 	if e != nil {
// 		panic(e)
// 	}
// 	fmt.Print(n)

// }
