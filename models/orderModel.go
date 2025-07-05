package models

type Type string

const (
	Takeaway Type = "Takeaway"
	Delivery Type = "Delivery"
	DineIn   Type = "DineIn"
)

type Status string

const (
	Placed    Status = "Order Placed"
	Pending   Status = "Order Pending"
	Ready     Status = "Order Ready"
	Canceled  Status = "Order Canceled"
	Delivered Status = "Order Delivered"
	Taken     Status = "Order Taken"
	Served    Status = "Order Served"
)

type Order struct {
	Name   string  `json:"name"`
	Type   Type    `json:"type"`
	Status Status  `json:"status"`
	Items  int     `json:"items"`
	Cost   float64 `json:"cost"`
}

func GetAllOrdersModel() []Order {
	orders := FetchOrders()
	return orders
}

func FetchOrders() []Order {
	return []Order{
		{"#ORD0011", DineIn, Placed, 2, 7.33},
		{"#ORD0012", Takeaway, Taken, 1, 3.85},
		{"#ORD0013", DineIn, Canceled, 7, 22.00},
		{"#ORD0014", Delivery, Pending, 8, 102.79},
		{"#ORD0015", DineIn, Served, 5, 33.5},
		{"#ORD0015", Delivery, Delivered, 13, 123.84},
		{"#ORD0015", Takeaway, Ready, 2, 13.75},
		{"#ORD0015", DineIn, Pending, 3, 23.05},
	}
}
