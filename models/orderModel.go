package models

type Type string
type Status string

const (
	Takeaway Type = "Takeaway"
	Delivery Type = "Delivery"
	DineIn   Type = "DineIn"

	Placed    Status = "Placed"
	Preparing Status = "Preparing"
	Ready     Status = "Ready"
	Canceled  Status = "Canceled"
	Delivered Status = "Delivered"
	Taken     Status = "Taken"
	Served    Status = "Served"
)

type Order struct {
	Name   string
	Type   Type
	Status Status
	Items  string
	Cost   string
}

func FetchOrders() []Order {
	return []Order{
		{"#ORD0011", DineIn, Placed, "2", "7.33"},
		{"#ORD0012", Takeaway, Taken, "1", "3.85"},
		{"#ORD0013", DineIn, Canceled, "7", "22.00"},
		{"#ORD0014", Delivery, Preparing, "8", "102.79"},
		{"#ORD0015", DineIn, Served, "5", "33.5"},
		{"#ORD0016", Delivery, Delivered, "13", "123.84"},
		{"#ORD0017", Takeaway, Ready, "2", "13.75"},
		{"#ORD0018", DineIn, Preparing, "3", "23.05"},
	}
}
