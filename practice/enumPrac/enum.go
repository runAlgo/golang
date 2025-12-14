package main

import "fmt"

// Define a custom type for the bitmask
type Permission int

// Define the permission flags using iota and bit shifting (<<)
const (
	// Bit shifting ensures each value is a power of 2 (1, 2, 4, 8, ...)
	PermRead   Permission = 1 << iota // 1 (0001)
	PermWrite                         // 2 (0010)
	PermDelete                        // 4 (0100)
	PermAdmin                         // 8 (1000)
)

// Function to check if a user has a specific permission
func HasPermission(userPerms Permission, checkPerm Permission) bool {
	// The bitwise AND operator checks if the 'checkPerm' bit is set in 'userPerms'
	return userPerms&checkPerm != 0
}

func main() {
	// A user who can read and write
	editor := PermRead | PermWrite

	// Usage example
	fmt.Printf("Editor Permissions: %d\n", editor) // Output: 3 (1 | 2)

	// Check permissions
	fmt.Println("Can Read:", HasPermission(editor, PermRead))     // true
	fmt.Println("Can Delete:", HasPermission(editor, PermDelete)) // false

	// An Admin has all permissions defined
	admin := PermRead | PermDelete | PermAdmin
	fmt.Println("Admin Can Delete:", HasPermission(admin, PermDelete)) // true
}


// Example (Super simple)
// Let’s say the user has read and write permissions:
// userPerms = 0011  (binary for 3)

// Now you want to check if write permission (0010) is set:
// checkPerm = 0010

// Bitwise AND:

// 0011   (userPerms)
// & 0010 (checkPerm)
// ------
// 0010   (result != 0)

