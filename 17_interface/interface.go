package main

import "fmt"


type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct{
	getway paymenter
}

// Open close principle
func (p payment) makePayment(amount float32) {
	// razorpayPaymentGW := razorpay {}
	// razorpayPaymentGW.pay(amount)
	// stripePaymentGW.pay(amount)
	// p.getway.pay(amount);
	p.getway.refund(amount, "1234")
}

// Fake Payment Gateway implementation
type fakepayment struct {}
func (f fakepayment) pay(amount float32) {
	fmt.Println("Making payment using fake gateway for testing purpost", amount)
}

// Rozorpay Payment Gateway implementation
type razorpay struct{}
func (p razorpay) pay(amount float32) {
	fmt.Printf("Make payment using razorpay: $%.2f", amount)
}

// Paypal Payment Gateway implementation
type paypal struct{}
func (p paypal) pay(amount float32) {
	fmt.Printf("Make payment using paypal: $%.2f\n", amount)
}
func (p paypal) refund(amount float32, account string) {
	fmt.Printf("Make refund with paypal: $%.2f, account: %s\n", amount, account)
}


// Stripe Payment Gateway implementation
type stripe struct {}
func (p stripe) pay(amount float32) {
	fmt.Printf("Make payment using stripe: $%.2f\n", amount)
}



func main() {
	// stripepaymentGW := stripe{}
	// razorpaypaymentGW := razorpay{}
	paypalpaymentGW := paypal {}
	newPayment := payment {
		getway: paypalpaymentGW,
	}
	newPayment.makePayment(200)
}