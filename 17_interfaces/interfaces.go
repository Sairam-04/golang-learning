package main

import "fmt"

// Implementing Inteface for pay method such that any payment gateway can be accessed and tested
type paymenter interface{
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGW := razorpay{}
	// razorpayPaymentGW.pay(amount)

	// stripePaymentGW := stripe{}
	// stripePaymentGW.pay(amount)

	p.gateway.pay(amount)
}

type razorpay struct {}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32){
	fmt.Println("making payment using stripe", amount)
}

type fakePayment struct{}

func (f fakePayment) pay(amount float32){
	fmt.Println("making fake payment for testing", amount)
}

type payPal struct{}

func (p payPal) pay(amount float32){
	fmt.Println("making payment using paypal", amount)
}

func main() {
	razorPayGW := razorpay{}
	stripeGW := stripe{}
	fakeGW := fakePayment{}
	payPalGW := payPal{}
	newPaymentRazorPay := payment{
		gateway: razorPayGW,
	}
	newPaymentRazorPay.makePayment(5700)
	newPaymentStripe := payment{
		gateway: stripeGW,
	}
	newPaymentStripe.makePayment(1700)
	newPaymentFake := payment{
		gateway: fakeGW,
	}
	newPaymentFake.makePayment(170)
	newPaymentPayPal := payment{
		gateway: payPalGW,
	}
	newPaymentPayPal.makePayment(4700)
}