package auth

import "fmt"

// If you want to use "LoginWithCredientials" package outside of its Scope
// Make sure Package Name is start with Capital letter.
func LoginWithCredientials(userName string, password string) {
	fmt.Printf("User Logged in successful:\nUsername:%s\nPassword:%s", userName, password)
}
