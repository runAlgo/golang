package main

// BadStruct: Random field order creates massive padding
type BadStruct struct {
	Flag    bool  // 1 byte + 7 padding
	Counter int64 // 8 bytes
	Active  bool  // 1 byte + 7 padding
} // Total: 24 bytes

// GoodStruct: Sorted by size (Largest => Smallest)
type GoodStruct struct {
	Counter int64 // 8 bytes
	Flag    bool  // 1 byte
	Active  bool  // 1 byte
	// + 6 padding at the end to match 8-byte alignment
} // Total: 16 bytes

// func main() {
// 	fmt.Printf("BadStruct size: %d bytes\n", unsafe.Sizeof(BadStruct{}))
// 	fmt.Printf("GoodStruct size: %d bytes\n", unsafe.Sizeof(GoodStruct{}))
// }
