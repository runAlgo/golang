package main

// func main() {
//1. ***** Create a new File *****

// f, err := os.Create("example.txt")

// if err != nil {
// 	panic(err)
// }
// defer f.Close()

// 2. **** Write into file ****
// bytes := []byte("Hello Golang")

// err := os.WriteFile("example.txt", bytes, 0644)
// 6 -> owner can read & write
// 4 -> group can read
// 4 -> others can read
// 0644 -> narmal files
// 0755 -> executable files
// if err != nil {
// 	fmt.Println(err)
// }

//3. **** Read whole file at once, Good for config files, small text, JSON***
// 	data, err := os.ReadFile("example.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(data))
// }
