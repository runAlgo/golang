package main

import (
	"fmt"

	"github.com/runAlgo/podcast/auth"
	"github.com/runAlgo/podcast/user"
)

func main() {
	auth.LoginWithCredientials("codersgyan", "secret")
	session := auth.GetSession()
	fmt.Println("User Session:", session)

	user := user.User{
		Email: "kalu@gmail.com",
		Name:  "Kalu Coder",
	}
	fmt.Println("User Details:", user.Email, user.Name)

	// color.Green(user.Email)
}
