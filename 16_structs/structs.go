package main

import (
	"fmt"
	"time"
)

// Struct Embedding
type customer struct{
	name string
	phone string
}
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}
// creating constructor 
func newOrder(id string, amount float32, status string) *order{
	myOrder := order{
		id: id,
		amount: amount,
		status: status,
		createdAt: time.Now(),
	}
	return &myOrder
}

// receiver type
// struct does dereferencing
func (o *order) changeStatus(status string){
	o.status= status
}

func (o *order) getAmount() float32{
	return o.amount
}

func main() {
	// create instance of order
	myOrder := order{
		id:     "1",
		amount: 599.00,
		status: "received",
		// Struct embedding
		customer: customer{
			name: "Sairam",
			phone: "1234567890",
		},
	}
	myOrder.createdAt = time.Now()
	order2 := order{
		id:     "2",
		amount: 699.00,
		status: "received",
	}
	order2.createdAt = time.Now()
	myOrder.changeStatus("pending")
	myNewOrder := newOrder("3", 899.00, "delievered")
	fmt.Println(myOrder, order2, myOrder.getAmount())
	fmt.Println("myNewOrder: ", myNewOrder)
}
