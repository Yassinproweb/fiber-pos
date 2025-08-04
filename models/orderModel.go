package models

import (
	"fmt"
)

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
	Transit   Status = "Transit"
	Delivered Status = "Delivered"
	Taken     Status = "Taken"
	Served    Status = "Served"
)

type OrderItem struct {
	PdtName   string
	Quantity  int
	UnitPrice float64
}

type Order struct {
	Name        string
	Type        Type
	Status      Status
	Items       int
	Cost        float64
	CustName    string
	CustNumber  string
	Destination string
	DateTime    string
	OrderCart   []OrderItem
}

func (o *Order) UpdateItemsAndCost() error {
	totalItems := 0
	totalCost := 0.0

	products := FetchProducts()
	productMap := make(map[string]Product)
	for _, p := range products {
		productMap[p.Name] = p
	}

	for _, item := range o.OrderCart {
		if item.Quantity < 0 {
			return fmt.Errorf("invalid quantity for product %s: quantity (%d) cannot be negative", item.PdtName, item.Quantity)
		}
		pdt, exists := productMap[item.PdtName]
		if !exists {
			return fmt.Errorf("product name %s not found in product list", item.PdtName)
		}
		if item.UnitPrice <= 0 {
			item.UnitPrice = pdt.Price
		}
		if item.UnitPrice < 0 {
			return fmt.Errorf("invalid unit price for product %s: unit price (%f) cannot be negative", item.PdtName, item.UnitPrice)
		}
		totalItems += item.Quantity
		totalCost += float64(item.Quantity) * item.UnitPrice
	}

	if totalItems < 0 {
		return fmt.Errorf("total items (%d) cannot be negative", totalItems)
	}
	if totalCost < 0 {
		return fmt.Errorf("total cost (%f) cannot be negative", totalCost)
	}

	o.Items = totalItems
	o.Cost = totalCost
	return nil
}

func FetchOrders() []Order {
	return []Order{
		{
			"#ORD0011",
			DineIn,
			Placed,
			2,
			7.33,
			"Ahmad",
			"0722678837",
			"Nakasozi, Wakiso",
			"11-06-2025 09:30",
			[]OrderItem{
				{"Posho & Beans", 3, 19.77},
			},
		},
	}
}
