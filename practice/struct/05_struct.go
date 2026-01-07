package main

// import (
// 	"time"
// )

// // 2. The Data Mode Struct (Database/API Entity)
// // This is the most common type of struct, used to represent data
// // entities flowing in your system, often mapping to a database table
// // or a JOSN payload. It relies heavily on struct tags.

// // Confidence Builder: This pattern shows you how to model real-world
// // objects while simultaneously preparing them for database and API
// // interaction.

// type User struct {
// 	ID        int    `json:"id,omitempty" db:"user_id"`
// 	FirstName string `json:"first_name" db:"first_name"`
// 	LastName  string `json:"last_name" db:"last_name"`
// 	Email     string `json:"email" db:"email"`
// 	IsActive  bool   `json:"is_active" db:"is_active"`

// 	// Unexported (private) field for security/business logic
// 	passwordHash string

// 	// time.Time is a common type for timestamps
// 	CreatedAt time.Time `json:"created_at" db:"created_at"`
// }

// // FullName is a simple method on the User struct.

// func (u *User) FullName() string {
// 	return u.FirstName + " " + u.LastName
// }

// // func main() {
// // 	newUser := User{
// // 		passwordHash: "sfdsbfdsof",
// // 		ID:           1,
// // 		FirstName:    "Kalu",
// // 		LastName:     "Don",
// // 		Email:        "kalu@gmail.com",
// // 		IsActive:     true,
// // 		CreatedAt:    time.Now(),
// // 	}

// // 	res := newUser.FullName()
// // 	fmt.Println(res)
// // }
