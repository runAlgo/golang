package main

// func main() {
// 	f, err := os.Open("example.txt")
// 	if err != nil {
// 		// log the error
// 		panic(err)
// 	}
// 	fileInfo, err := f.Stat()
// 	if err != nil {
// 		// log the error
// 		panic(err)
// 	}

// 	fmt.Println("file name:", fileInfo.Name())
// 	fmt.Println("File or folder:", fileInfo.IsDir())
// 	fmt.Println("File size:", fileInfo.Size())
// 	fmt.Println("File permission:", fileInfo.Mode())
// 	fmt.Println("File modified at:", fileInfo.ModTime())

//***** Read File ****

// f, err := os.Open("example.txt")
// if err != nil {
// 	// log error
// 	panic(err)
// }
// defer f.Close()

// fileInfo, err := f.Stat()
// if err != nil {
// 	// log error
// 	panic(err)
// }

// fileInfo.Size() gives me the size of "example.txt"
// buf := make([]byte, fileInfo.Size())

// d, err := f.Read(buf)
// if err != nil {
// 	// log the error
// 	panic(err)
// }

// fmt.Println("Size:", d)
// for i := 0; i < len(buf); i++ {
// fmt.Println(string(buf[i]))
// }

// **** Read File Data in Simple (Only use for small files) ****

// data, err := os.ReadFile("example.txt")

// if err != nil {
// 	// log err
// 	panic(err)
// }
// fmt.Println("Data:", string(data))

// ***** Read Folders ****

// dir, err := os.Open("../")
// if err != nil {
// 	panic(err)
// }
// defer dir.Close()

// fileInfo, err := dir.ReadDir(-1)
// for _, fi := range fileInfo {
// 	fmt.Println(fi.Name(), fi.IsDir())
// }

// **** Create a file *****

// f, err := os.Create("example2.txt")
// if err != nil {
// 	panic(err)
// }
// defer f.Close()

// f.WriteString("Hi Go")

//**** Simple way to Write into File *****
// f, err := os.Create("example.txt")
// if err != nil {
// 	panic(err)
// }
// defer f.Close()

// bytes := []byte("Hello Golang")

// f.Write(bytes)

// Read existing fie
// data, err := os.ReadFile("example2.txt")
// if err != nil {
// 	panic(err)
// }

// Replace text
// updated := strings.ReplaceAll(
// 	string(data),
// 	"one",
// 	"two",
// )

// Overwrite the same file
// err = os.WriteFile("example2.txt", []byte(updated), 0644)
// 0644 means: This is the file permission mode (Unix/Linux style).
// Owner: read + write
// Group: read only
// Others: read only

// if err != nil {
// 	panic(err)
// }

// }
