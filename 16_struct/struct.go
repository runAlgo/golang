package main

import (
	"fmt"
	"time"
)

// We don't have class on Go programming thats, we have struct
type order struct {
	id string
	amount float32
	status string
	createdAt time.Time // nanosecond precision
}

// func newOrder(id string, amount float32, status string) *order {
// 	// initial setup goes here...
// 		myOrder := order{
// 		id: id,
// 		amount: amount,
// 		status: status,
// 	}
// 	return &myOrder
// }

// receiver type
// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// func (o *order) changePrice(price float32){
// 	o.amount = price
// }


// func (o order) getPrice() float32 {
// 	return o.amount
// }




func main() {

	language := struct {
		name string
		isGood bool
	}{"golang", true}

	fmt.Println(language)

	// myOrder := newOrder("1", 30.55, "received")
	// fmt.Println(myOrder.amount)
	// If you don't set any field, default value is zero value
	// int => 0, float => 0, bool => false
	// myOrder := order{
	// 	id: "1",
	// 	amount: 50.00,
	// 	// status: "received",
	// }

	// myOrder.changeStatus("Delivered")
	// myOrder.changePrice(44.99)
	// fmt.Println("Price is:",myOrder.getPrice())
	// fmt.Println(myOrder)

	// myOrder2 := order{
	// 	id: "2",
	// 	amount: 100.00,
	// 	status: "Pending",
	// 	createdAt: time.Now(),
	// }

	// myOrder.status = "paid"


	// myOrder.createdAt = time.Now()
	// fmt.Println("Order Struct:", myOrder)
	// fmt.Println("Second Order Struct:", myOrder2)
}