package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}
// creating constructor 
// func newOrder() *order{

// }

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
	}
	myOrder.createdAt = time.Now()
	order2 := order{
		id:     "2",
		amount: 699.00,
		status: "received",
	}
	order2.createdAt = time.Now()
	myOrder.changeStatus("pending")
	fmt.Println(myOrder, order2, myOrder.getAmount())
}
