package main

import "fmt"

// enumerated types
type OrderStatus string

const (
	// Received OrderStatus = iota
	// Confirmed
	// Prepared
	// Delivered
	
	Received  OrderStatus = "received"
	Confirmed OrderStatus = "confirmed"
	Prepared  OrderStatus = "prepared"
	Delivered OrderStatus = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	switch status {
	case Received, Confirmed, Prepared, Delivered:
		fmt.Println("Changing order status to", status)
	default:
		fmt.Println("Invalid status:", status)
	}

}
func main() {
	changeOrderStatus(Received)
}
