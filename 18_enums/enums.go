package main

import "fmt"

// enumerated types
// type OrderStatus int

// const (
// 	Received OrderStatus = iota
// 	Confirmed 
// 	Prepared
// 	Delievered
// )

type OrderStatus string

const (
	Received OrderStatus = "received"
	Confirmed = "confirmed"
	Prepared = "prepared"
	Delievered = "delievered"
)

func changeOrderStatus(status OrderStatus){
	fmt.Println("chaging order status to", status)
}

func main(){
	changeOrderStatus(Delievered)
}